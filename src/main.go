package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", server.AddMiddleware(PostRequest, CheckAuth(), Logging()))
	server.Handle("POST", "/user", server.AddMiddleware(UserPostRequest, CheckAuth(), Logging()))
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
