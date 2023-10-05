package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(mysql:3306)/todo_db?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todos", getTodos)
	r.POST("/todos", addTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run(":8080")
}

func bindTodo(c *gin.Context) (Todo, error) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return Todo{}, err
	}
	return todo, nil
}

func getTodos(c *gin.Context) {
	todos, err := fetchTodos()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, todos)
}

func fetchTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT todo_id, todo_text FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Text); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, rows.Err()
}

func addTodo(c *gin.Context) {
	todo, err := bindTodo(c)
	if err != nil {
		return
	}

	newTodo, err := addAndFetchTodo(todo)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(201, newTodo)
}

func addAndFetchTodo(todo Todo) (Todo, error) {
	result, err := db.Exec("INSERT INTO todos (todo_text) VALUES (?)", todo.Text)
	if err != nil {
		return Todo{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Todo{}, err
	}
	todo.ID = int(id)
	return todo, nil
}

func updateTodo(c *gin.Context) {
	todo, err := bindTodo(c)
	if err != nil {
		return
	}

	updatedTodo, found, err := updateAndFetchTodo(c.Param("id"), todo)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if found {
		c.JSON(200, updatedTodo)
	} else {
		c.JSON(404, gin.H{"error": "Todo not found"})
	}
}

func updateAndFetchTodo(id string, updatedTodo Todo) (Todo, bool, error) {
	_, err := db.Exec("UPDATE todos SET todo_text = ? WHERE todo_id = ?", updatedTodo.Text, id)
	if err != nil {
		return Todo{}, false, err
	}
	row := db.QueryRow("SELECT todo_id, todo_text FROM todos WHERE todo_id = ?", id)
	var todo Todo
	if err := row.Scan(&todo.ID, &todo.Text); err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, false, nil
		}
		return Todo{}, false, err
	}
	return todo, true, nil
}

func deleteTodo(c *gin.Context) {
	if err := deleteTodoByID(c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"status": "Todo deleted"})
}

func deleteTodoByID(id string) error {
	_, err := db.Exec("DELETE FROM todos WHERE todo_id = ?", id)
	return err
}
