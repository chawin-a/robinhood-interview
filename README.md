# Robinhood Assignment
 
## Prerequisites
- Docker
- Go
- [go-migrate](https://github.com/golang-migrate/migrate)

## Getting started
1. Clone the repository
```bash

git clone https://github.com/chawin-a/robinhood-interview.git

```
2. Run docker compose
```bash
docker compose up -d
```

3. Run migrations
```bash
migrate -database postgres://postgres:password@localhost:5432/db?sslmode=disable -path migrations up
```

4. Init mock data
```bash
go run scripts\mock\main.go
```

## Database Design
![alt text](https://github.com/chawin-a/robinhood-interview/blob/main/images/db-diagram.png?raw=true)

## Documentation

#### Get user details

<details>
 <summary><code>POST</code> <code><b>/user/:id</b></code> </summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | uuid   | N/A  |


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `Configuration created successfully`                                |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8889/
> ```

</details>

------------------------------------------------------------------------------------------
