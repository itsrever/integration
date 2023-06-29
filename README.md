# REVER integration testing framework

This framework aims to help integrators to develop an [integration with REVER](https://api.byrever.com/v1/docs/#implementing-your-own-integration). It consist in a set of testing scenarios that you can run against your implementation (in your own CI, makefile, testing framework....) to reduce the feedback loop. In that way, the output of the tests will help you to identify the integration issues and fix them before going live, and what's more important, before waiting REVER development team to verify it.

Any feedback is welcomed. Please open an issue or a PR if you find any bug or you want to suggest any improvement. The REVER development team will be releasing versions when the integration contract changes, so you can pin your integration to a specific version of this framework.

## How it works

You will need `Docker` installed in your system.
The framework consists of:

* A docker container to execute the testing scenarios against your implementation
* A set of testing scenarios that you can have a look to investigate what the expected behavior is
* A dummy implementation of the API that you can use to test the framework, in Go. You can use it as a reference to implement your own integration.
* A [sample configuration file](./test/config.json) that you can use to configure the testing scenarios and adapt to your implementation

The steps execute the suite of tests are:

* Configure your integration as in the [config.json](./test/config.json) file. You can configure it against a development server or the production environment. Just put the correct URLs in the file
* Start your development server. Any language or framework is ok. This step is not necessary if your implementation is already running.
* Run the docker container with your configuration. This will run the testing scenarios against your implementation and will output the results in the console.
* Stop your development server.

### Configuring the integration

The [config.json](./test/config.json) file contains a global section, independent of any endpoint. The mimimum configuration is setting where your API lives and how to pass the authentication token. For a development server, an example is:

``` json
    "base_url": "http://localhost:8080",
    "auth": {
        "header": "x-api-key",
        "api_key": "valid-api-key"
    },
```

Per each endpoint of the integration, below the `tests` section you will have to configure the specific tests and pass the values that matches the scenarios. For example, for the scenario `FIND04` you will need to set the value of `customer_order_printed_id` that matched the described scenario in your platform.

> Valid order with multiple `line_items`, referring products/services **without variants**. Implement this case if your e-commerce supports products but has no support for Variants...

Please refer to the [config.json](./test/config.json) file for a complete example. All of the supported tests are listed in the section [Methods included in the testing](#methods-included-in-the-testing).

``` json
{
    "method":"FindOrderByCustomerOrderPrintedId",
    "url_pattern": "/integration/orders/find?customer_order_printed_id={customer_order_printed_id}",
    "scenarios": [
        {
            "name": "FIND00",
            "customer_order_printed_id": "non-empty"
        },
        {
            "name": "FIND01",
            "customer_order_printed_id": "whatever"
        },
    ...
}
```

### Running the docker container

The docker container is available in the [Docker Hub](https://hub.docker.com/r/itsrever/testing). You can run it with the following command:

``` bash
docker run --rm -v "route_to_your_config.json/test:/rever/test/config.json" itsrever/testing:latest
```

## Methods included in the testing

* [FindOrderByCustomerPrintedOrderId](./docs/FindOrderByCustomerPrintedOrderId.md)
