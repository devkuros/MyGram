package main

import "mygram/routers"

func main() {
	r := routers.StartServer()
	r.Run(":8080")
}
