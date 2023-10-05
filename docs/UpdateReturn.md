# CreateRefund

this method will be used to udpate the previous return detail in the ecommerce platform.

## Testing scenarios

The implementation of this method **must** support the following scenarios:

| Scenario | Status | Description |
| -------- | ------ | ----------- |
| UPDATERETURN001 | 200 | The requests updates a return |
| UPDATERETURN002 | 401 | The requests does not contain the authentication header |
| UPDATERETURN003 | 401 | The requests contains an invalid authentication header |
| UPDATERETURN004 | 404 | The `customer_printed_order_id` is empty |
| UPDATERETURN005 | 404 | The `customer_printed_order_id` is not empty but not found in the integration |
| UPDATERETURN006 | 404 | The `return_id` is not found in the integration |


## Test Configuration example

```
"method":"UpdateeReturn",
"url_pattern": "/integration/orders/{customer_printed_order_id}/{return_id}/return",
"scenarios": [
   {
        "name": "UPDATERETURN001",
        "customer_printed_order_id": "simple_order_1"
        "return_id": "simple_return_1",
    }
]
```

