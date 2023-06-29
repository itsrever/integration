# REVER integration testing framework

This framework aims to help integrators to develop an [integration with REVER](https://api.byrever.com/v1/docs/#implementing-your-own-integration). It consist in a set of testing scenarios that you can run against your implementation (in your own CI, makefile, testing framework....) to reduce the feedback loop. In that way, the output of the tests will help you to identify the integration issues and fix them before going live, and what's more important, before waiting REVER development team to verify it.

Any feedback is welcomed, please open an issue or a PR if you find any bug or you want to suggest any improvement. The REVER development team will be releasing versions when the integration contract changes, so you can pin your integration to a specific version of this framework.

## How it works

You will need `Docker` installed in your system.
The framework consists of:

    - A docker container to execute the testing scenarios against your implementation
    - A set of testing scenarios that you can have a look to investigate what the expected behavior is
    - A dummy implementation of the API that you can use to test the framework, in Go. You can use it as a reference to implement your own integration.
    - A [sample configuration file](./test/config.json) that you can use to configure the testing scenarios and adapt to your implementation

In your CI or project, execute the docker container against a running version of your API (the one that implements the integration with REVER) and will run the tests against it.

## Methods

* [FindOrderByCustomerPrintedOrderId](./docs/FindOrderByCustomerPrintedOrderId.md)
