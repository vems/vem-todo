package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	todoclient "github.com/vems/todo-service-go/client"
	"github.com/vems/todo-service-go/todo"
	"google.golang.org/grpc"
)

type client struct {
	Todo todo.Service
}

func mustDial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	return conn
}

func main() {
	port := os.Getenv("PORT")
	// default for port
	if port == "" {
		port = "8080"
	}

	todoAddr := os.Getenv("TODO_ADDR")
	// default for port
	if todoAddr == "" {
		todoAddr = "todo:8080"
	}

	c := client{
		Todo: todoclient.New(mustDial(todoAddr), log.NewNopLogger()),
	}

	r := gin.Default()

	r.GET("/todos", listTodos(c))
	r.POST("/todos", createTodo(c))
	r.DELETE("/todos", deleteAllTodos(c))
	r.GET("/todos/:id", getTodo(c))
	r.PUT("/todos/:id", updateTodo(c))
	r.DELETE("/todos/:id", deleteTodo(c))

	r.Run(":" + port)
}
