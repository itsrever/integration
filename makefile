
OPENAPI_URL=https://api.byrever.com/v1/docs/openapi_integration.yaml
OPENAPI_LOCAL=tmp/openapi.yaml
JSON_SCHEMA=test/schema.json
CLIENT_PATH=client
SERVER_PATH=server
EXEC_FILE=rever-server-integration
APP_NAME=testing
# Go source files, ignore vendor directory
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./server/go/*")
##############################
# TESTING					 #
##############################

install-gotestfmt:
	go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest

unit-test: 
	go build -o ./bin/${EXEC_FILE} server/main.go 
	./bin/${EXEC_FILE} &
	sleep 2
	(go test -json -v ./test/... 2>&1 | tee /tmp/gotest.log | gotestfmt) || pkill -9  ${EXEC_FILE}

unit-test-ci: install-gotestfmt unit-test

with-docker-test-linux:
	go build -o ./bin/${EXEC_FILE} server/main.go 
	./bin/${EXEC_FILE} &
	docker run --rm -v "${PWD}/test/config.json:/rever/test/config.json" --network="host"  itsrever/testing:latest
	pkill -9  ${EXEC_FILE}

with-docker-test-linux-ci: docker-build
	go build -o ./bin/${EXEC_FILE} server/main.go 
	./bin/${EXEC_FILE} &
	docker run --rm -v "${PWD}/test/config.json:/rever/test/config.json" --network="host"  $(APP_NAME)

with-docker-test-mac:
	go build -o ./bin/${EXEC_FILE} server/main.go 
	./bin/${EXEC_FILE} &
	docker run --rm -v "${PWD}/sample/config.macos.json:/rever/test/config.json" --network="host"  itsrever/testing:latest
	pkill -9  ${EXEC_FILE}

with-docker-test-qooqer: docker-build docker-tag
	docker run --rm -v "${PWD}/sample/config.qooqer.json:/rever/test/config.json"  itsrever/testing:latest

with-docker-test-win:
	go build -o ./bin/${EXEC_FILE} server/main.go 
	./bin/${EXEC_FILE} &
	docker run --rm -v "${PWD}/sample/config.win.json:/rever/test/config.json" --network="host"  itsrever/testing:latest
	pkill -9  ${EXEC_FILE}

in-docker-test:
	go test -json -v ./test/... 2>&1 | tee /tmp/gotest.log | gotestfmt 

docker-build: 
	docker build --platform=linux/amd64 -t $(APP_NAME) .

docker-tag:
	docker tag $(APP_NAME) itsrever/testing:latest

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
	echo "Replacing `integration.` with empty string to avoid long names..."
	sed -i '' "s/integration\.//g" ${OPENAPI_LOCAL}

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
	--inline-schema-name-mappings \
		Shipping_taxes=multi_money,Shipping_amount=multi_money,\
Order_shipping_address=address,Order_billing_address=address,\
Order_total_taxes=multi_money,Order_total_amount=multi_money,\
LineItem_subtotal=multi_money,LineItem_total=multi_money,\
LineItem_total_discounts=multi_money,LineItem_total_taxes=multi_money,\
LineItem_unit_price=multi_money \
    -o /local/${SERVER_PATH}

gen-go-server: download-openapi openapi-generator-srv openapi-to-json clean-server update-libs

install-openapi-schema-to-json-schema: 
	npm install @openapi-contrib/openapi-schema-to-json-schema
	npm install ytoj 

openapi-to-json: install-openapi-schema-to-json-schema
	npx ytoj --resolve-refs --input ${OPENAPI_LOCAL} --output ${JSON_SCHEMA}

##############################
# MOCK SERVER PUBLISHING     #
##############################

docker-server:
	docker build -t itsrever/integration-server:latest -f server/Dockerfile .


install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

format: install-goimports
	@gofmt -l -w $(SRC)
	@goimports -w -e -local github.com/itsrever/integration $(SRC)
	
install-lint-ubuntu:
	echo Installing yamlint golangci-lint...
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v1.50.1
	golangci-lint --version

install-lint-macos:
	brew install golangci-lint
	
lint: format
	@golangci-lint -v --timeout=600s --skip-dirs=docs run 