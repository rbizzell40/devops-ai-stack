# Makefile for devops-ai-stack

CLUSTER_NAME= ?= devops-ai-stack
KIND_CONFIG=kind/kind-cluster.yaml

.PHONY: help kind-up kind-delete bootstrap-argocd install-secrets deploy-mcp demo

help:
	@echo "Usage:"
	@echo "  make kind-up             # Create KIND cluster"
	@echo "  make kind-delete         # Delete KIND cluster"
	@echo "  make bootstrap-argocd    # Install ArgoCD and App of Apps"
	@echo "  make install-secrets     # Deploy ESO and inject secrets"
	@echo "  make deploy-mcp          # Build and deploy MCP agent"
	@echo "  make demo                # Run the demo script"

kind-up:
	kind create cluster --name $(CLUSTER_NAME) --config $(KIND_CONFIG)

kind-delete:
	kind delete cluster --name $(CLUSTER_NAME)

bootstrap-argocd:
	helm repo add argo https://argoproj.github.io/argo-helm
	helm repo update
	helm upgrade --install argocd argo/argo-cd -n argocd --create-namespace
	kubectl apply -f argo-apps/apps-of-apps.yaml

install-secrets:
	helm repo add external-secrets https://charts.external-secrets.io
	helm repo update
	helm upgrade --install external-secrets external-secrets/external-secrets -n external-secrets --create-namespace
	kubectl apply -f secrets/

deploy-mcp:
	cd mcp-agent/scripts && go build -o ../../tools/mcp-cli/mcp
	kubectl apply -f mcp-agent/workflows/

demo:
	chmod +x ./demo.sh
	./demo.sh
