.PHONY: docker
docker:
	docker build . -t post-svc:latest
