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

</details>
<details>
 
#### Get interview details

 <summary><code>POST</code> <code><b>/interview/:id</b></code> </summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | uuid   | N/A  |


##### Responses

```javascript
{
    "interview": {
        "Id": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
        "UserId": "86839593-8d05-4798-8d86-e93f411edc08",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ligula purus, pulvinar vel nisi in, feugiat gravida lorem. Phasellus elit nunc, posuere ac ante sit amet, bibendum iaculis mi.",
        "Status": "To Do",
        "IsArchived": false,
        "CreatedAt": "2023-08-15T03:08:02.970709Z",
        "UpdatedAt": "2023-08-15T03:08:02.970709Z"
    }
}
```

</details>
<details>
 
 #### Archive interview
 
 <summary><code>POST</code> <code><b>/interview/:id/archive</b></code> </summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | uuid   | N/A  |


##### Responses

```javascript
{
    "interview": {
        "Id": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
        "UserId": "86839593-8d05-4798-8d86-e93f411edc08",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ligula purus, pulvinar vel nisi in, feugiat gravida lorem. Phasellus elit nunc, posuere ac ante sit amet, bibendum iaculis mi.",
        "Status": "To Do",
        "IsArchived": true,
        "CreatedAt": "2023-08-15T03:08:02.970709Z",
        "UpdatedAt": "2023-08-15T03:17:08.492502Z"
    }
}
```

</details>
<details>
 
 #### Update status interview
 
 <summary><code>POST</code> <code><b>/interview/:id/update_status</b></code> </summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | uuid   | N/A  |


##### Body

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | status      |  required | ENUM("To Do", "In Progress", "Done")   | N/A  |

##### Responses

```javascript
{
    "interview": {
        "Id": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
        "UserId": "86839593-8d05-4798-8d86-e93f411edc08",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ligula purus, pulvinar vel nisi in, feugiat gravida lorem. Phasellus elit nunc, posuere ac ante sit amet, bibendum iaculis mi.",
        "Status": "In Progress",
        "IsArchived": true,
        "CreatedAt": "2023-08-15T03:08:02.970709Z",
        "UpdatedAt": "2023-08-15T03:18:15.789981Z"
    }
}
```

</details>
<details>
 
 #### List interview's comments
 
 <summary><code>POST</code> <code><b>/interview/:id/comment</b></code> </summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | uuid   | N/A  |


##### Responses

```javascript
{
    "comments": [
        {
            "Id": "92035213-14fa-4182-9d5c-2ae05c87dafa",
            "UserId": "7474c95f-5a85-433a-bc39-6bfe66c4ddf4",
            "InterviewId": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
            "Comment": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
            "CreatedAt": "2023-08-15T03:08:02.981083Z"
        },
        {
            "Id": "209530ab-59bb-4d8a-a682-77a250eb6e2b",
            "UserId": "a228d32b-0fd0-47d8-8dbc-4d9fc9a9b9c2",
            "InterviewId": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
            "Comment": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
            "CreatedAt": "2023-08-15T03:08:02.979885Z"
        },
        {
            "Id": "0eb265ba-0197-4a0f-a12f-47cddc6e3e9b",
            "UserId": "7474c95f-5a85-433a-bc39-6bfe66c4ddf4",
            "InterviewId": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
            "Comment": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
            "CreatedAt": "2023-08-15T03:08:02.978672Z"
        },
        {
            "Id": "74d0b58a-0bd2-40ed-b171-27b4f05ecf23",
            "UserId": "1074ece6-2c8c-4fc4-9df6-35bcd2d11f38",
            "InterviewId": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
            "Comment": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
            "CreatedAt": "2023-08-15T03:08:02.975957Z"
        }
    ]
}
```

</details>

<details>
 
 #### Create interview's comments

 <summary><code>POST</code> <code><b>/interview/:id/comment/create</b></code> </summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | uuid   | N/A  |

##### Body

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | comment      |  required | string   | N/A  |



##### Responses

```javascript
{
    "comment": {
        "Id": "fa380cb2-f49a-48d6-ba5f-4c3e0308f4de",
        "UserId": "1074ece6-2c8c-4fc4-9df6-35bcd2d11f38",
        "InterviewId": "2b0f6430-48c5-484c-a0b0-010d386bddbb",
        "Comment": "hello",
        "CreatedAt": "2023-08-15T03:22:27.255629Z"
    }
}
```

</details>

------------------------------------------------------------------------------------------
