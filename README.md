# friends-management
A Restful API for simple Friends Management application with GO, using go-chi and testify 

## API Documentation
https://documenter.getpostman.com/view/3974024/TzJpjg6x

#How to run

For run docker-compose, run this commands in project's root folder:

```bash
docker-compose build
docker-compose up
```

Or run local with Makefile:

```bash
make all
```

If you want to cross-compile your application to run on every OS:

```bash
make compile
```

## REST Api
```sh
1 User Registration: http://localhost:8080/api/v1/user
  Example Request
    {
	  "email": "andy@example.com"
    }
  Response Example
    {
      "success":true
    }
2 API to create a friend connection between two email addresses : http://localhost:8080/api/v1/friend

  Example Request
    {
      "friends":[
        "andy@example.com",
        "john@example.com"
        ]
    }
  Response Example
    {
      "success":true
    }
3 API to retrieve the friends list for an email address :  http://localhost:8080/api/v1/friend/friends-list
  Example Request
    {
    "email":"andy@example.com"
    }
  Response Example
    {
        "success": true,
        "friends": [
            "john@example.com"
        ],
        "count": 1
    }
4 API to retrieve the common friends list between two email addresses :  http://localhost:8080/api/v1/friend/common-friends
  Example Request
      {
        "friends":[
          "andy@example.com",
          "john@example.com"
          ]
      }
  Response Example
      {
        "success": true,
        "friends": [
            "common@example.com"
        ],
        "count": 1
      }
5 API to subscribe to updates from an email address : http://localhost:8080/api/v1/subscribe
  Example Request
    {
      "requestor" : "lisa@example.com",
      "target" : "join@example.com"
    }
  Response Example
    {
      "success": true
    }
6 API to block updates from an email address: http://localhost:8080/api/v1/block
  Example Request
    {
      "requestor" : "andy@example.com",
      "target" : "join@example.com"
    }
  Response Example
    {
      "success": true
    }
7 API to retrieve all email addresses that can receive updates from an email address :  http://localhost:8080/api/v1/friend/receive-updates
  Example Request
    {
      "sender" :"john@example.com",
      "text" : "Hello World! kate@example.com"
    }
   Response Example
    {
      "success": true,
      "recipients": [
          "lisa@example",
          "kate@example"
      ]
    }
```
## Unit Testing

From the terminal: 

### Go to handlers, service, repositories commands folder

go test -v

## Project Architecture


```bash
friends_management
├─ database
│  └─ database_connect.go
├─ handlers
│  ├─ block.go
│  ├─ block_test.go
│  ├─ errors.go
│  ├─ friend.go
│  ├─ friend_test.go
│  ├─ subscribe.go
│  ├─ subscribe_test.go
│  ├─ user.go
│  └─ user_test.go
├─ helper
│  ├─ connectDb.go
│  ├─ fixture.go
│  └─ utils.go
├─ initilization
│  └─ DBTable.go
├─ models
│  ├─ block.go
│  ├─ friend.go
│  ├─ subscribe.go
│  └─ user.go
├─ repositories
│  ├─ test_migration
│  │  ├─ block
│  │  │  └─ block.sql
│  │  ├─ friend
│  │  │  └─ friend.sql
│  │  ├─ subscribe
│  │  │  └─ subscribe.sql
│  │  ├─ user
│  │  │  └─ user.sql
│  │  └─ truncate_table.go
│  ├─ block.go
│  ├─ block_test.go
│  ├─ friend.go
│  ├─ friend_test.go
│  ├─ subscribe.go
│  ├─ subscribe_test.go
│  ├─ user.go
│  └─ user_test.go
├─ routes
│  └─ routes.go
├─ service
│  ├─ block_service.go
│  ├─ block_service_mock.go
│  ├─ block_service_test.go
│  ├─ friend_service.go
│  ├─ friend_service_mock.go
│  ├─ friend_service_test.go
│  ├─ subscribe_service.go
│  ├─ subscribe_service_mock.go
│  ├─ subscribe_service_test.go
│  ├─ user_service.go
│  ├─ user_servgice_mock.go
│  └─ user_service_test.go
├─ .env
├─ docker-compose.yml
├─ Dockerfile
├─ go.mod
│  └─ go.sum
├─ main.go
├─ Makerfile.md
└─ README.md

```

- Workflow: Request => Handlers => Services => Repositories => DB

#Note References

https://blog.golang.org/pipelines

https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html

https://www.golang-book.com/books/intro
