package main

import (
	"awesomeProject/test/internal/server"
	"fmt"
	"log"
)

func main() {
	ginServer := server.NewGinServer("localhost", 3000)
	ginServer.SetRoute()
	err := ginServer.StartServer()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("server started ad port %d", 3000))
	select {}
}
