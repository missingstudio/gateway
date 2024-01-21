default: build

build: build-protos build-mobius
.PHONY: build

build-mobius:
	sh -c "cd ./mobius && make all"
.PHONY: build-mobius

build-protos:
	sh -c "cd ./protos && make all"
.PHONY: build-protos