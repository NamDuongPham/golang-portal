IMAGE = namduong0606/golang-portal:latest

.PHONY: deployapp build push deploy restart status logs help

deployapp: ## Build, push and deploy
	bash ./run.sh --tag $(TAG)

build: ## Build image only
	bash ./run.sh --tag $(TAG) --build-only

push: ## Push image only
	bash ./run.sh --tag $(TAG) --push-only

deploy: ## Deploy only (kubectl set-image/rollout)
	bash ./run.sh --tag $(TAG) --deploy-only

restart: ## Rollout restart deployment
	kubectl rollout restart deployment/golang-portal -n $(NAMESPACE)

status: ## Check rollout status
	kubectl rollout status deployment/golang-portal -n $(NAMESPACE)

logs: ## Tail pod logs
	kubectl logs -l app=golang-portal --tail=100 -f -n $(NAMESPACE)

help: ## Show this help
	@echo "Usage: make <target> [TAG=tag] [NAMESPACE=namespace]"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS=":.*?## "} {printf "  %-15s %s\n", $$1, $$2}'
	@echo ""
	@echo "Examples:"
	@echo "  make deployapp TAG=v1.0.0"
	@echo "  make build TAG=ci-123"
	@echo "  make push TAG=ci-123"
	@echo "  make deploy TAG=ci-123"

