IMAGE="valkyrie00/curriculum-telegram"

.PHONY: build
build:
	@docker build --force-rm -t $(IMAGE) -f ./deployments/Dockerfile .

.PHONY: push
push:
	@docker push $(IMAGE)