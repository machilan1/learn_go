default: init run
	
run:
	go run ./cmd/web 

init:
	go mod tidy

.PHONY: default run init