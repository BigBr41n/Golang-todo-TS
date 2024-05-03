package main 

import (
	"fmt"
	"log"
    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}


func main(){
	fmt.Println("hello world")

	app:= fiber.New()
	

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
        AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	todos := []Todo{}

	app.Get("/healthCheck" , func(c * fiber.Ctx) error {
		return c.SendString("ok")
	})


	app.Post("/api/v1/todos" , func(c * fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err!= nil {
            return c.Status(400).SendString(err.Error())
        }

		todo.ID = len(todos) + 1 


		todos = append(todos, *todo)

		return c.JSON(todos)
	})


	app.Patch("/api/v1/todos/:id/done" , func(c * fiber.Ctx) error  {
		id, err := c.ParamsInt("id")

        if err!= nil {
            return c.Status(400).SendString(err.Error())
        }

        for i, todo := range todos {
            if todo.ID == id {
                todos[i].Done = true
                return c.JSON(todos[i])
            }
        }

        return c.Status(404).SendString(fmt.Sprintf("todo with id %d not found", id))
	} )

	app.Get("/api/v1/todos", func(c * fiber.Ctx) error{
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":3000"))
}