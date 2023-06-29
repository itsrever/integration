
OPENAPI_URL=https://api.byrever.com/v1/docs/openapi_integration.yaml
OPENAPI_LOCAL=tmp/openapi.yaml
CLIENT_PATH=client
SERVER_PATH=server
EXEC_FILE=rever-server-integration

##############################
# TESTING					 #
##############################

unit-test:
	go build -o ./bin/${EXEC_FILE} server/main.go 
	./bin/${EXEC_FILE} &
	sleep 2
	go test -json -v ./test/... 2>&1 | tee /tmp/gotest.log | gotestfmt || pkill -9  ${EXEC_FILE}

##################
# API GENERATION #
##################

clean-server:
	echo Cleaning up...
	rm -rf ${SERVER_PATH}/.openapi-generator
	rm -rf ${SERVER_PATH}/.openapi-generator-ignore
	rm -rf ${SERVER_PATH}/api
	rm -rf ${SERVER_PATH}/go.mod
	rm -rf ${SERVER_PATH}/go.sum
	rm -rf ${SERVER_PATH}/README.md
	rm -rf ${SERVER_PATH}/Dockerfile

download-openapi:
	mkdir -p tmp
	curl -o ${OPENAPI_LOCAL} ${OPENAPI_URL}

update-libs:
	go get golang.org/x/oauth2
	go get github.com/gorilla/mux
	go mod tidy
	go mod vendor

openapi-generator-srv:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.1 generate \
	 -i /local/${OPENAPI_LOCAL} \
	--additional-properties=packageName=server \
    -g go-server \
	--git-repo-id=integration/server --git-user-id=itsrever \
	--ignore-file-override=/local/.openapi-generator-ignore \
    -o /local/${SERVER_PATH}

gen-go-server: download-openapi openapi-generator-srv clean-server update-libs
	