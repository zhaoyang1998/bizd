package main

import (
	"bizd/router"
	"fmt"
)

func main() {
	router := router.SetupRouter()
	if err := router.Run(":8888"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
