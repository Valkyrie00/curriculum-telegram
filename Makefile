IMAGE="valkyrie00/curriculum-telegram:arm"

.PHONY: build
build:
	@docker build --force-rm -t $(IMAGE) -f ./deployments/Dockerfile .

.PHONY: push
push:
	@docker push $(IMAGE)