REGISTRY ?= shipudemo.azurecr.io
IMAGE_NAME := awiexample-go
IMAGE_VERSION ?= latest

DEMO_IMAGE := $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_VERSION)

# The architecture of the image
ARCH ?= amd64

.PHONY: build
build:
	docker build \
		--platform="linux/$(ARCH)" \
		--tag=$(DEMO_IMAGE) .

.PHONY: push
push:
	docker push $(DEMO_IMAGE)