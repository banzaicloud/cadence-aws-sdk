bin/aws-sdk-generator:
	mkdir -p bin/
	GOBIN=$$PWD/bin go get -modfile generator.mod go.temporal.io/aws-sdk-generator

generate: bin/aws-sdk-generator
	bin/aws-sdk-generator --template-dir templates --output-dir .
