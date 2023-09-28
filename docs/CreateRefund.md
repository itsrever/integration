# CreateRefund

Perfome a refund to the original payment method for a given order.

## Testing scenarios

The implementation of this method **must** support the following scenarios:

| Scenario | Status | Description |
| -------- | ------ | ----------- |
| CREATEREFUND00 | 401 | The requests does not contain the authentication header |
| CREATEREFUND01 | 401 | The requests contains an invalid authentication header |
| CREATEREFUND02 | 400 | The `customer_printed_order_id` is empty |
| CREATEREFUND03 | 404 | The `customer_printed_order_id` is not empty but not found in the integration |
| ... | 200 | .... |

The implementation of this method **might** support the following scenarios:

| Scenario | Status | Description |
| -------- | ------ | ----------- |
| CREATEREFUND04 | 200 | The `customer_printed_order_id` is valid and the refund is performed| 

| ... | 200 | .... |

## Test Configuration example

```
"method":"CreateRefund",
"url_pattern": "createRefund?customer_order_printed_id={customer_order_printed_id}",
"scenarios": [
    {
        "name": "CREATEREFUND04",
        "customer_printed_order_id": "order_to_be_refunded"
    },
]
```

