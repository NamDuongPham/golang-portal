IMAGE = namduong0606/golang-portal:latest

ifeq ($(OS),Windows_NT)
# Windows: prefer bash if available, fall back to WSL; otherwise instruct user
deploy:
	powershell -NoProfile -ExecutionPolicy Bypass -Command "if (Get-Command bash -ErrorAction SilentlyContinue) { bash -c './run.sh' } elseif (Get-Command wsl -ErrorAction SilentlyContinue) { wsl bash -c './run.sh' } else { Write-Host 'ERROR: bash not found. Open Git Bash or WSL and run make there (or run bash ./run.sh).'; exit 1 }"

build:
	powershell -NoProfile -ExecutionPolicy Bypass -Command "if (Get-Command bash -ErrorAction SilentlyContinue) { bash -c 'chmod +x ./run.sh || true; ./run.sh --build-only' } elseif (Get-Command wsl -ErrorAction SilentlyContinue) { wsl bash -c 'chmod +x ./run.sh || true; ./run.sh --build-only' } else { Write-Host 'ERROR: bash not found.'; exit 1 }"

push:
	powershell -NoProfile -ExecutionPolicy Bypass -Command "if (Get-Command bash -ErrorAction SilentlyContinue) { bash -c './run.sh --push-only' } elseif (Get-Command wsl -ErrorAction SilentlyContinue) { wsl bash -c './run.sh --push-only' } else { Write-Host 'ERROR: bash not found.'; exit 1 }"

restart:
	kubectl rollout restart deployment/golang-portal

status:
	kubectl rollout status deployment/golang-portal

logs:
	kubectl logs -l app=golang-portal --tail=100 -f

else
# Unix-like shells (Git Bash, WSL, Linux, macOS)
deploy:
	chmod +x ./run.sh || true
	bash -c "./run.sh"

build:
	chmod +x ./run.sh || true
	bash -c "./run.sh --build-only"

push:
	chmod +x ./run.sh || true
	bash -c "./run.sh --push-only"

restart:
	kubectl rollout restart deployment/golang-portal

status:
	kubectl rollout status deployment/golang-portal

logs:
	kubectl logs -l app=golang-portal --tail=100 -f

endif
