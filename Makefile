DEBUG_FLAG = $(if $(DEBUG),-debug)


build:
	./scripts/dist.sh

deps:
	go get -d -t ./...

install: deps
	go install
