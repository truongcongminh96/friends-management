# friends-management

user: postgres

db: friends-management

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
- 3 layers model:
    + Handlers: receive request, validate request and response
    + Services: handle business logics
    + Repositories: interact with database

- Workflow: Request => Handlers => Services => Repositories => DB
