# friends-management

user: postgres

db: friends-management

## REST Api
```sh
1 User Registration: http://localhost:8080/api/v1/registration
  Example Request
    {
	  "email" : "user@example"
    }
  Response Example
    {
      "success":true
    }
2 API to create a friend connection between two email addresses : http://localhost:8080/api/v1/friendConnection

  Example Request
    {
      "friends":[
        "andy@example",
        "john@example"
        ]
    }
  Response Example
    {
      "success":true
    }
3 API to retrieve the friends list for an email address :  http://localhost:8080/api/v1/retrieveFriendList
  Example Request
    {
    "email":"andy@example"
    }
  Response Example
    {
        "success": true,
        "friends": [
            "john@example"
        ],
        "count": 1
    }
4 API to retrieve the common friends list between two email addresses :  http://localhost:8080/api/v1/commonFriends
  Example Request
      {
        "friends":[
          "andy@example",
          "john@example"
          ]
      }
  Response Example
      {
        "success": true,
        "friends": [
            "common@example"
        ],
        "count": 1
      }
5 API to subscribe to updates from an email address : http://localhost:8080/api/v1/subscribe
  Example Request
    {
      "requestor" : "lisa@example",
      "target" : "join@example"
    }
  Response Example
    {
      "success": true
    }
6 API to block updates from an email address: http://localhost:8080/api/v1/blockFriend
  Example Request
    {
      "requestor" : "andy@example",
      "target" : "join@example"
    }
  Response Example
    {
      "success": true
    }
7 API to retrieve all email addresses that can receive updates from an email address :  http://localhost:8080/api/v1/receiveUpdates
  Example Request
    {
      "sender" :"john@example",
      "text" : "Hello World! kate@example"
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

### Go to service commands folder
cd service
## Run all tests
go test -v


# Note Ex:

	func ExampleHandleFunc() {
    	h1 := func(w http.ResponseWriter, _ *http.Request) {
    		io.WriteString(w, "Hello from a HandleFunc #1!\n")
    	}
    	h2 := func(w http.ResponseWriter, _ *http.Request) {
    		io.WriteString(w, "Hello from a HandleFunc #2!\n")
    	}
    
    	http.HandleFunc("/", h1)
    	http.HandleFunc("/endpoint", h2)
    
    	log.Fatal(http.ListenAndServe(":8080", nil))
    }


# references
https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97

https://codingpearls.com/go-programming/concurrency-ket-hop-voi-restful-api-trong-golang.html

https://www.golang-book.com/books/intro/9
