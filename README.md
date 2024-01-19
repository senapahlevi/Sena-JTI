
# JTI Task

Backend used for to create cashier invoices quickly and easily. This application is designed to help small and medium businesses manage their sales transactions.

You can try using `postman` and 

#### 
[Try Demo this]()



## Tools Stack

**Database:** PostgreSQL

**Tools:** Golang

**Framework:** Gin

**Library:** Gorm


## Design

[Figma](https://linktodocumentation)


## Run Locally

Clone the project

```bash
  git clone https://github.com/senapahlevi/Sena-JTI
```

Go to the project directory

```bash
  cd my-project
```

Start the server

```bash
go run main.go
```


## API
Using OAUTH0
```bash
  go run main.go
```
and will program will Listening and serving HTTP on `:3000`
#### Save Number 
- `POST`

`http://localhost:3000/create`

body json 
```bash
{
    "phone" :"08121082193",
    "provide" :"telkomsel"
}
```
response json:
```bash

    {
        "id": 700,
        "phone": "08121082193",
        "provider": "telkomsel",
        "is_odd": 1
    }

```
#### Update Invoice  
- `PUT`
`http://localhost:3000/update-invoice/43`
```bash
{
    "data": {
        "invoice_id": 43,
        "subject": "test subject",
        "status": "paid",
        "issued_date": "2020-01-02T00:00:00Z",
        "due_date": "2020-02-03T00:00:00Z",
        "sub_total": 6990,
        "tax": 699,
        "grand_total": 6291,
        "detail_item_json": "[{ \"item_name\":\"Toyowheels\",\"item_type\":\"tools\",\"quantity\":1,\"unit_price\":6990}]",
        "customer": "rudi tabuti",
        "address": "newcastle",
        "detail_items": null
    }
}
```

#### Detail All Invoice /GET All invoice
- `GET`
`http://localhost:3000/invoice`

```bash
{
    "data": [
        {
            "invoice_id": 43,
            "subject": "test subject",
            "status": "paid",
            "issued_date": "2020-01-02T07:00:00+07:00",
            "due_date": "2020-02-03T07:00:00+07:00",
            "sub_total": 6990,
            "tax": 699,
            "grand_total": 6291,
            "detail_item_json": "[{ \"item_name\":\"Toyowheels\",\"item_type\":\"tools\",\"quantity\":1,\"unit_price\":6990}]",
            "customer": "rudi tabuti",
            "address": "newcastle",
            "detail_items": null
        },
        {
            "invoice_id": 44,
            "subject": "test subject",
            "status": "paid",
            "issued_date": "2023-02-03T07:00:00+07:00",
            "due_date": "2020-02-03T07:00:00+07:00",
            "sub_total": 7590,
            "tax": 759,
            "grand_total": 6831,
            "detail_item_json": "[{\"invoice_id\":1, \"item_name\":\"toys\",\"item_type\":\"service\",\"quantity\":10,\"unit_price\":20 },{\"invoice_id\":1, \"item_name\":\"ice cream\",\"item_type\":\"service\",\"quantity\":1,\"unit_price\":690 }]",
            "customer": "boedi",
            "address": "Test address",
            "detail_items": null
        }
    ]
}
```
#### Delete Invoice button  
- `DELETE`
`http://localhost:3000/detail-items/1`
 
 
#### Get ID data 
`GET`

`http://localhost:3000/output-id/700 



response json:
```bash
[
    {
        "id": 700,
        "phone": "086969181293913",
        "provider": "Tri",
        "is_odd": 1
    }
]
```
#### Detail ID Invoice  

- `GET`
`http://localhost:3000/invoice/43`
```bash
{
    "data": {
        "invoice_id": 43,
        "subject": "test subject",
        "status": "paid",
        "issued_date": "2020-01-02T07:00:00+07:00",
        "due_date": "2020-02-03T07:00:00+07:00",
        "sub_total": 6990,
        "tax": 699,
        "grand_total": 6291,
        "detail_item_json": "[{ \"item_name\":\"Toyowheels\",\"item_type\":\"tools\",\"quantity\":1,\"unit_price\":6990}]",
        "customer": "rudi tabuti",
        "address": "newcastle",
        "detail_items": [
            {
                "item_id": 39,
                "InvoiceID": 43,
                "item_name": "Toyowheels",
                "item_type": "tools",
                "quantity": 1,
                "unit_price": 6990,
                "amount": 6990
            }
        ]
    }
}
```




- Add more integrations


