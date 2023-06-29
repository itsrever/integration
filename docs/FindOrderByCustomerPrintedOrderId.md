# FindOrderByCustomerPrintedOrderId

Finds a single order by its `customer_printed_order_id`.
In case of multiple matches, the most recent one is returned.

## Testing scenarios

The implementation of this method **must** support the following scenarios:

| Scenario | Status | Description |
| -------- | ------ | ----------- |
| FIND00 | 401 | The requests does not contain the authentication header |
| FIND01 | 401 | The requests contains an invalid authentication header |
| FIND02 | 400 | The `customer_printed_order_id` is empty |
| FIND03 | 404 | The `customer_printed_order_id` is not empty but not found in the integration |
| ... | 200 | .... |

The implementation of this method **might** support the following scenarios:

| Scenario | Status | Description |
| -------- | ------ | ----------- |
| FIND04 | 200 | Valid order with multiple `line_items`, referring products/services **without variants**. Implement this case if your e-commerce supports products but has no support for Variants. Product variants are a requirement for supporting exchange orders as compensation method. The order must have a positive amount in EUR, with taxes and shipping costs. Regarding the payment method, must be paid with a non-cash, non-cash on delivery, non-BNPL payment method. It should have a discount applied. It must be associated with a valid customer. It must be fulfilled and paid |
| FIND05 | 200 | Valid order with multiple `line_items` referring products/services **with variants**. Implement this case if your e-commerce supports products and variants. Please note that supporting variants is required to allow exchange orders as compensation method. The order must have a positive amount in EUR, with taxes and shipping costs. Regarding the payment method, must be paid with a non-cash, non-cash on delivery, non-BNPL payment method. It should have a discount applied. It must be associated with a valid customer. It must be fulfilled and paid |
| ... | 200 | .... |

## Configuration example

```
"method":"FindOrderByCustomerOrderPrintedId",
"url_pattern": "FindOrderByCustomerOrderPrintedId?customer_order_printed_id={customer_order_printed_id}",
"scenarios": [
    {
        "name": "FIND00",
        "customer_order_printed_id": "non-empty"
    },
    {
        "name": "FIND01",
        "customer_order_printed_id": "non-empty"
    },
    {
        "name": "FIND02",
        "customer_order_printed_id": ""
    },
    {
        "name": "FIND03",
        "customer_order_printed_id": "non-existing"
    },
    {
        "name": "FIND04",
        "customer_order_printed_id": "simple_order_1"
    }
]
```

