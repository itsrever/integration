# AddNoteToOrder

Add a note to an order given its `order_id`.

## Testing scenarios

The implementation of this method **must** support the following scenarios:

| Scenario | Status | Description |
| -------- | ------ | ----------- |
| FIND00 | 401 | The requests does not contain the authentication header |
| FIND01 | 401 | The requests contains an invalid authentication header |
| FIND02 | 400 | The `order_id` is empty |
| FIND03 | 404 | The `order_id` is not empty but not found in the integration |
| ... | 200 | .... |


| Scenario | Status | Description |
| -------- | ------ | ----------- |
| FIND06 | 200 | Add a note to an order given the `order_id` and the `note` to be added.
| ... | 200 | .... |

## Configuration example

```
"method":"AddNoteToOrder",
"url_pattern": "orders/{order_id}/note",
"scenarios": [
    {  
        "name": "FIND06",
        "order_id": "simple_order_1"
    }
]