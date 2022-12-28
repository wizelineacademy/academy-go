package main

import (
	"fmt"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	router.NewRouter(r)

	fmt.Println("running app") // TODO - replace with logger
	err := r.Run()             // TODO - port from env vars
	if err != nil {
		fmt.Println("error running app") // TODO - replace with logger
	}
}
