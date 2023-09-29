# CreateRefund

Once the item is received in the warehouse, this will create the return detail in the ecommerce platform.

## Testing scenarios

The implementation of this method **must** support the following scenarios:

| Scenario | Status | Description |
| -------- | ------ | ----------- |
| CREATERETURN001 | 200 | The requests creates a return |
| CREATERETURN002 | 401 | The requests does not contain the authentication header |
| CREATERETURN003 | 401 | The requests contains an invalid authentication header |
| CREATERETURN004 | 404 | The `customer_printed_order_id` is empty |
| CREATERETURN005 | 404 | The `customer_printed_order_id` is not empty but not found in the integration |


## Test Configuration example

```
"method":"CreateReturn",
"url_pattern": "/integration/orders/{customer_printed_order_id}/return",
"scenarios": [
   {
        "name": "CREATERETURN001",
        "customer_printed_order_id": "simple_order_1"
    }
]
```

