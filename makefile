# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)


run:
	go run api/cmd/service/sales/main.go | go run api/cmd/tooling/logfmt/main.go

# ==============================================================================
# Define dependencies
GOLANG := golang:1.22.4
ALPINE := alpine:3.20
KIND   := kindest/node:v1.25.16

KIND_CLUSTER    := nicholas-starter-cluster

# ==============================================================================
# Install dependencies
dev-brew:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list pgcli || brew install pgcli
	brew list watch || brew install watch

# ==============================================================================
# Running from within k8s/kind
dev-up:
	kind create cluster \
		--image $(KIND) \
		--name	$(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pod -o wide --watch --all-namespaces


.PHONY: run dev-brew dev-up dev-down dev-status-all