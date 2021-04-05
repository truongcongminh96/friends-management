# friends-management

user: postgres

pass: 1

db: friends-management

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
