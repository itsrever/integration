# REVER integration testing framework

This framework aims to help integrators to develop an [integration with REVER](https://api.byrever.com/v1/docs/#implementing-your-own-integration). It consists in a set of testing scenarios that you can run against your implementation (in your own CI, makefile, testing framework....) to reduce the feedback loop. This way, the output of the tests will help you identify and fix integration issues before going live, without waiting for the REVER engineering team to verify it.

Any feedback is welcome. Please open an issue or a PR if you find any bug or you want to suggest any improvement. The REVER team will be releasing versions when the integration contract changes, so you can pin your integration to a specific version of this framework.

## How it works

You will need `Docker` installed in your system.
The framework consists of:

* A docker container to execute the testing scenarios against your implementation
* A set of testing scenarios that you can have a look to investigate what the expected behavior is
* A dummy implementation of the API that you can use to test the framework, in Go. You can use it as a reference to implement your own integration.
* A [sample configuration file](./test/config.json) that you can use to configure the testing scenarios and adapt to your implementation

The steps to execute the suite of tests are:

* Configure your integration as in the [config.json](./test/config.json) file. You can configure it against a development server or the production environment. Just put the correct URLs in the file
* Start your development server. Any language or framework is ok. This step is not necessary if your implementation is already running.
* Run the docker container with your configuration. This will run the testing scenarios against your implementation and will output the results in the console.
* Stop your development server.

If all of the tests go well, you will see something like this:

``` bash
📦 REVER integration testing framework
  ...
  ✅ Test_FindOrderByCustomerOrderPrintedId (10ms)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND00 (0s)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND01 (0s)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND02 (0s)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND03 (0s)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND04 (0s)
  ...
```

If a test fails, you will see the reason why it fails.

``` text
  ❌ Test_FindOrderByCustomerOrderPrintedId/FIND04 (0s)
     Error:  Not equal: 
                expected: 404
                actual  : 200
     Test:   Test_FindOrderByCustomerOrderPrintedId/FIND04
```

### Configuring the integration

The [config.json](./test/config.json) file contains a global section, independent of any endpoint. The minimum configuration is setting where your API lives and what authentication needs to be used. For the development server and using an API Key is:

``` json
    "base_url": "http://localhost:8080",
    "auth": {
        "header": "x-api-key",
        "api_key": "valid-api-key"
    },
```

For using OAuth2 with the client credential flow (two-legged), the configuration is:

``` json
    "base_url": "http://localhost:8080",
    "oauth2": {
        "client_id": "client_id",
        "client_secret": "client_secret",
        "token_url": "https://your-oauth2-server/token"
    },
```

For each endpoint of the integration, under the tests section, you will need to configure the specific tests and pass the values that match the scenarios. For example, for the scenario `FIND04` you will need to set the value of `customer_order_printed_id` that matches the described scenario in your platform.

Some test scenarios are optional. If you don't implement them, then do not add a configuration for them.

For example, with this JSON, the optional test `FIND05` won't be executed:

``` json
{
    "method":"FindOrderByCustomerOrderPrintedId",
    "url_pattern": "/integration/orders/find?customer_printed_order_id={customer_printed_order_id}",
    "scenarios": [
        {
            "name": "FIND04",
            "customer_printed_order_id": "your-order-id"
        },
    ]
}
```

The result will indicate that the test was skipped:

``` bash
📦 github.com/itsrever/integration/test
  ✅ TestApplyVars (0s)
  ✅ TestComposeURL (0s)
  ✅ Test_FindOrderByCustomerOrderPrintedId (130ms)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND00 (30ms)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND01 (10ms)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND02 (10ms)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND03 (10ms)
  ✅ Test_FindOrderByCustomerOrderPrintedId/FIND04 (50ms)
  🚧 Test_FindOrderByCustomerOrderPrintedId/FIND05 (0s)
    config.go:87: Skipping: Scenario FIND05 not present in config
```

> **FIND04**: Valid order with multiple `line_items`, referring products/services **without variants**. Implement this case if your e-commerce supports products but has no support for Variants...

Please refer to the [config.json](./test/config.json) file for a complete example. All of the supported tests are listed in the section [Methods included in the testing](#methods-included-in-the-testing).

``` json
{
    "method":"FindOrderByCustomerOrderPrintedId",
    "url_pattern": "/integration/orders/find?customer_printed_order_id={customer_printed_order_id}",
    "scenarios": [
        {
            "name": "FIND04",
            "customer_printed_order_id": "your-order-id"
        },
    ]
}
```

> **FIND05**: Valid order with multiple `line_items`, referring products/services **with variants**. Implement this case if your e-commerce supports products and variants...

Please refer to the [config.json](./test/config.json) file for a complete example. All of the supported tests are listed in the section [Methods included in the testing](#methods-included-in-the-testing).

``` json
{
    "method":"FindOrderByCustomerOrderPrintedId",
    "url_pattern": "/integration/orders/find?customer_printed_order_id={customer_printed_order_id}",
    "scenarios": [
        {
            "name": "FIND05",
            "customer_printed_order_id": "your-order-id"
        },
    ]
}
```

> **FIND06**: Add a note to an order given the `order_id` and the `note` to be added.

``` json
{
    "method":"AddNoteToOrder",
    "url_pattern": "/integration/orders/{order_id}/note",
    "scenarios": [
        {
            "name": "FIND06",
            "order_id": "simple_order_1"
        }
    ]
}
```

### Running the docker container

The docker container is available in the [Docker Hub](https://hub.docker.com/r/itsrever/testing).
You can run it with the following command:

```bash
docker run --rm -v "/path/to/your/config.json:/rever/test/config.json" \
    --network="host" \
    itsrever/testing:latest
```

Please note the `/path/to/your/config.json` must be an absolute path.

### Running the docker container against localhost (local dev server)

If you're running your API locally (`localhost`), the docker container needs to access the `host` network. This is because the container needs to access your API running in the host machine. You can do it with the `--network="host"` option. In linux that should be it. If you're running MacOS or Windows, you will need modify the URL of your api in the `config.json` file like in the [macos config file sample](./test/config.mac.json) or [Windows config file sample](./test/config.win.json).

Then, you can specify this file to docker:

````bash
docker run --rm -v "${PWD}/sample/config.macos.json:/rever/test/config.json" \
    --network="host" \
    itsrever/testing:latest
````

#### Running a specific test

The run command can be overridden to only run a specific test. Here is an example. Replace TestMyFunction with whatever test you need. For example Test_FindOrderByCustomerOrderPrintedId/FIND03

```bash

docker run --rm -v "${PWD}/sample/config.macos.json:/rever/test/config.json" \
    --network="host" \
    itsrever/testing:latest go test -run TestMyFunction ./test/... -json -v 2>&1 | tee /tmp/gotest.log | gotestfmt 
```

#### Developing Locally

This is a Go project, another way to run it is by opening it in your favorite editor. To make this simpler we've included a devcontainer definition.

For more info on how to setup the devcontainer and VSCode please look at https://code.visualstudio.com/docs/devcontainers/create-dev-container#_set-up-a-folder-to-run-in-a-container. Only VSCode with the devcontainer extension along with Docker is needed to setup the editor.

Once you install and launch the project you can run and debug specific tests whenever needed.

## Deploying the dummy implementation

The commands can be tested using a mock server::

``` bash
curl --header "x-rever-api-key:valid-api-key" "https://server-tsem47dtaa-ey.a.run.app/integration/orders/find?customer_printed_order_id=simple_order_1" 
```

## Methods included in the testing

* [FindOrderByCustomerPrintedOrderId](./docs/FindOrderByCustomerPrintedOrderId.md)
* [AddNoteToOrder](./docs/AddNoteToOrder.md)
* [CreateRefund](./docs/CreateRefund.md)
* [CreateReturn](./docs/CreateReturn.md)
