# Robinhood Assignment
 
## Prerequisites
- Docker
- Go 1.21.0
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
migrate -path migrations -database "postgres://postgres:password@localhost:5432/db?sslmode=disable" up
```

4. Init mock data
```bash
go run scripts/mock/main.go
```

## Database Design
![alt text](https://github.com/chawin-a/robinhood-interview/blob/main/images/db-diagram.png?raw=true)

## Documentation
<details>
 
 #### Get user details
 
 <summary><code>POST</code> <code><b>/user/:id</b></code> </summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | uuid   | N/A  |


##### Responses

```javascript
{
    "user": {
        "Id": "033d6d81-23ce-484b-97f7-d06b7ae0aba6",
        "Username": "user5",
        "Name": "user5",
        "Email": "user5@robinhood.co.th",
        "CreatedAt": "2023-08-15T03:02:11.326823Z"
    }
}
```

</details>
<details>
 
 #### List interviews details
 
 <summary><code>POST</code> <code><b>/interview</b></code> </summary>

##### Body

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | limit      |  required | int   | N/A  |
> | latest_timestamp      |   | timestamptz   | N/A  |


##### Responses

```javascript
{
    "interviews": [
        {
            "Id": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
            "UserId": "86839593-8d05-4798-8d86-e93f411edc08",
            "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ligula purus, pulvinar vel nisi in, feugiat gravida lorem. Phasellus elit nunc, posuere ac ante sit amet, bibendum iaculis mi.",
            "Status": "To Do",
            "IsArchived": false,
            "CreatedAt": "2023-08-15T03:08:02.970709Z",
            "UpdatedAt": "2023-08-15T03:08:02.970709Z"
        },
        {
            "Id": "ebbeca29-dfc9-49a0-b392-c1668be9520e",
            "UserId": "7474c95f-5a85-433a-bc39-6bfe66c4ddf4",
            "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ligula purus, pulvinar vel nisi in, feugiat gravida lorem. Phasellus elit nunc, posuere ac ante sit amet, bibendum iaculis mi.",
            "Status": "To Do",
            "IsArchived": false,
            "CreatedAt": "2023-08-15T03:08:02.982283Z",
            "UpdatedAt": "2023-08-15T03:08:02.982283Z"
        }
    ],
    "latest_timestamp": "2023-08-15T03:08:02.982283Z"
}
```

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8889/
> ```

</details>
<details>
 
#### Get interview details

 <summary><code>POST</code> <code><b>/interview/:id</b></code> </summary>

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
<details>
 
 #### Archive interview
 
 <summary><code>POST</code> <code><b>/interview/:id/archive</b></code> </summary>

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
<details>
 
 #### Update status interview
 
 <summary><code>POST</code> <code><b>/interview/:id/update_status</b></code> </summary>

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
<details>
 
 #### List interview's comments
 
 <summary><code>POST</code> <code><b>/interview/:id/comment</b></code> </summary>

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

<details>
 
 #### Create interview's comments

 <summary><code>POST</code> <code><b>/interview/:id/comment/create</b></code> </summary>

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
