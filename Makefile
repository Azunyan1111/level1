GO=$(shell which go)
domain=localhost
port=8888
url=$(domain):$(port)

run-level1:
	$(GO) run ./level1.go

setup:
	@echo do nothing

curl-sample1:
	curl -i -X POST -H "Content-Type: application/json" $(url)/api/checkout -d @sample1.json;

curl-sample2:
	curl -i -X POST -H "Content-Type: application/json" $(url)/api/checkout -d @sample2.json;
