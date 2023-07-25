# AddNoteToOrder

Add a note to an order given its `customer_printed_order_id`. 
When recovering back the order, the note should be present.

## Testing scenarios

The implementation of this method **must** support the following scenarios:

| Scenario | Status | Requires Config | Description | 
| -------- | ------ | --------------- | ----------- |
| ADDNOTE00 | 401 | no | The requests does not contain the authentication header |
| ADDNOTE01 | 401 | no | The requests contains an invalid authentication header |
| ADDNOTE02 | 400 | no | The `customer_printed_order_id` is empty |
| ADDNOTE03 | 404 | no | The `customer_printed_order_id` is not empty but not found in the integration |
| ADDNOTE04 (*) | 200 | yes | The `customer_printed_order_id` is valid and the note is appended |
| ... | 200 | .... |

The `Requires Config` column indicates if the configuration file must contain an entry for this scenario.

(*) In the ADDNOTE04 scenario, the testing platform will send a different note for each test run. It's possible
that the same order, with the ID indicated in the configuration file, gets accumulated with multiple random notes.

## Test Configuration example

```
"method":"AddNoteToOrder",
"url_pattern": "orders/{customer_printed_order_id}/note",
"scenarios": [
    {  
        "name": "ADDNOTE04",
        "customer_printed_order_id": "order_to_be_appended"
    }
]