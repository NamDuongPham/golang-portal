IMAGE = namduong0606/golang-portal:latest

deploy:
	sh run.sh

build:
	docker build -t $(IMAGE) .

push:
	docker push $(IMAGE)

restart:
	kubectl rollout restart deployment/golang-portal

status:
	kubectl rollout status deployment/golang-portal

logs:
	kubectl logs -l app=golang-portal --tail=100 -f
