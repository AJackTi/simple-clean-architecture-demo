package main

import (
	"log"
	"simple-clean-architecture-demo/modules/task/business"
	"simple-clean-architecture-demo/modules/task/repository/inmem"
	"simple-clean-architecture-demo/modules/task/transport/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// Setup full service dependencies
	apiService := rest.NewAPI(business.NewBusiness(inmem.NewInMemStorage()))

	v1 := engine.Group("v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", apiService.CreateTask())
			tasks.GET("", apiService.ListTask())
			tasks.GET("/:id", apiService.GetTask())
			tasks.PATCH("/:id", apiService.UpdateTask())
			tasks.DELETE("/:id", apiService.DeleteTask())
		}
	}

	if err := engine.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
