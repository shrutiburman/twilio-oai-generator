.PHONY: install generate test-docker

OPENAPI_GENERATOR_VERSION=5.4.0

install:
	wget -N https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar
	ln -sf openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar openapi-generator-cli.jar
	mvn clean package -DskipTests

generate: install
	bash examples/generate.sh

test-docker: generate
	bash ./prism.sh
