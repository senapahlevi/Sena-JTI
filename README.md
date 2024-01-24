
# JTI Task


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
-Output for see Data, Edit and Delete
http://localhost:3000/output
-Edit
example: `http://localhost:3000/output-edit/3`



