package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var todos = make([]Todo, 0)
var idCounter int
var todoLock sync.Mutex

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todos", getTodos)
	r.POST("/todos", addTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run(":8080")
}

// バリデーション
func bindTodo(c *gin.Context) (Todo, error) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return Todo{}, err
	}
	return todo, nil
}

func getTodos(c *gin.Context) {
	c.JSON(200, fetchTodos())
}

func fetchTodos() []Todo {
	todoLock.Lock()
	defer todoLock.Unlock()

	return todos
}

func addTodo(c *gin.Context) {
	todo, err := bindTodo(c)
	if err != nil {
		return
	}

	c.JSON(201, addAndFetchTodo(todo))
}

func addAndFetchTodo(todo Todo) Todo {
	todoLock.Lock()
	defer todoLock.Unlock()

	idCounter++
	todo.ID = idCounter
	todos = append(todos, todo)
	return todo
}

func updateTodo(c *gin.Context) {
	todo, err := bindTodo(c)
	if err != nil {
		return
	}

	updatedTodo, found := updateAndFetchTodo(c.Param("id"), todo)
	if found {
		c.JSON(200, updatedTodo)
	} else {
		c.JSON(404, gin.H{"error": "Todo not found"})
	}
}

func updateAndFetchTodo(id string, updatedTodo Todo) (Todo, bool) {
	todoLock.Lock()
	defer todoLock.Unlock()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Todo{}, false
	}

	for index, todo := range todos {
		if todo.ID == idInt {
			todos[index].Text = updatedTodo.Text
			return todos[index], true
		}
	}

	return Todo{}, false
}

func deleteTodo(c *gin.Context) {
	if deleteTodoByID(c.Param("id")) {
		c.JSON(200, gin.H{"status": "Todo deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Todo not found"})
	}
}

func deleteTodoByID(id string) bool {
	todoLock.Lock()
	defer todoLock.Unlock()

	fmt.Println(id)

	return false
}
