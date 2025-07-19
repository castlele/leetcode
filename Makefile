.PHONY: test fmt

test:
	go test ./go/... -coverpkg=./...

fmt:
	go fmt ./go/...
