
OPENAPI_URL=https://api.byrever.com/v1/docs/openapi_integration.yaml
OPENAPI_LOCAL=tmp/openapi.yaml

clear-client:
	rm -rf client

clean-client:
	echo Cleaning up...
	rm -rf client/.openapi-generator
	rm -rf client/api
	rm -rf client/docs
	rm -rf client/test
	rm -rf client/.gitignore
	rm -rf client/.openapi-generator-ignore
	rm -rf client/.travis.yml
	rm -rf client/go.mod
	rm -rf client/go.sum
	rm -rf client/git_push.sh
	rm -rf client/README.md

download-openapi:
	mkdir -p tmp
	curl -o ${OPENAPI_LOCAL} ${OPENAPI_URL}

update-libs:
	go get golang.org/x/oauth2
	go mod tidy
	go mod vendor

openapi-generator-cli:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.1 generate \
	 -i /local/${OPENAPI_LOCAL} \
	--additional-properties=packageName=client \
	--ignore-file-override=/local/.openapi-generator-ignore \
    -g go \
    -o /local/client/

gen-go-client: download-openapi clear-client openapi-generator-cli clean-client update-libs
	
	