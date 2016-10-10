package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"

	"github.com/vems/todo-service-go/model"
	"golang.org/x/net/context"
)

type NewTodoForm struct {
	Title string `json:"title" binding:"required,lte=255"`
}

type UpdateTodoForm struct {
	Title     string `json:"title" binding:"required,lte=255"`
	Completed bool   `json:"completed"`
}

func listTodos(c client) gin.HandlerFunc {
	return func(ginc *gin.Context) {
		ctx := context.Background()

		res, err := c.Todo.All(ctx)
		if err != nil {
			fmt.Println(err)
			ginc.JSON(500, gin.H{"status": err.Error()})
			return
		}

		ginc.JSON(200, res)
	}
}

func deleteAllTodos(c client) gin.HandlerFunc {
	return func(ginc *gin.Context) {
		ctx := context.Background()

		err := c.Todo.DeleteAll(ctx)
		if err != nil {
			fmt.Println(err)
			ginc.JSON(500, gin.H{"status": err.Error()})
			return
		}

		ginc.JSON(200, gin.H{"status": "ok"})
	}
}

func createTodo(c client) gin.HandlerFunc {
	return func(ginc *gin.Context) {
		t := NewTodoForm{}

		ginc.BindWith(&t, binding.JSON)

		ctx := context.Background()

		newTodo := model.Todo{
			Title: t.Title,
		}

		res, err := c.Todo.Create(ctx, newTodo)
		if err != nil {
			fmt.Println(err)
			ginc.JSON(422, gin.H{"status": err.Error()})
			return
		}

		ginc.JSON(200, res)
	}
}

func getTodo(c client) gin.HandlerFunc {
	return func(ginc *gin.Context) {
		ctx := context.Background()

		id := ginc.Param("id")

		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ginc.JSON(500, err)
			return
		}

		res, err := c.Todo.Find(ctx, intid)
		if err != nil {
			fmt.Println(err)
			ginc.JSON(404, gin.H{"status": "not found"})
			return
		}

		ginc.JSON(200, res)
	}
}

func updateTodo(c client) gin.HandlerFunc {
	return func(ginc *gin.Context) {
		t := UpdateTodoForm{}

		ginc.BindWith(&t, binding.JSON)

		id := ginc.Param("id")

		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ginc.JSON(500, err)
			return
		}

		ctx := context.Background()

		updateTodo := model.Todo{
			Title:     t.Title,
			Completed: t.Completed,
		}

		res, err := c.Todo.Update(ctx, intid, updateTodo)
		if err != nil {
			fmt.Println(err)
			ginc.JSON(404, gin.H{"status": "not found"})
			return
		}

		ginc.JSON(200, res)
	}
}

func deleteTodo(c client) gin.HandlerFunc {
	return func(ginc *gin.Context) {
		id := ginc.Param("id")

		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ginc.JSON(500, err)
			return
		}

		ctx := context.Background()

		res, err := c.Todo.Delete(ctx, intid)
		if err != nil {
			fmt.Println(err)
			ginc.JSON(404, gin.H{"status": "not found"})
			return
		}

		ginc.JSON(200, res)
	}
}
