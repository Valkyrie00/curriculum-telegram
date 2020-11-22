IMAGE="valkyrie00/curriculum-telegram:arm-v1.0.0"

.PHONY: build
build:
	@docker build --force-rm -t $(IMAGE) -f ./deployments/Dockerfile .

.PHONY: push
push:
	@docker push $(IMAGE)