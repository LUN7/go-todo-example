package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type ToDoMap struct {
	m     map[string]ToDo
	order []string
}

func MakeToDoMap() *ToDoMap {
	toDoList := new(ToDoMap)
	toDoList.m = make(map[string]ToDo)
	toDoList.order = make([]string, 0)
	return toDoList
}

func (toDoMap *ToDoMap) add(
	todo *ToDo,
) {
	toDoMap.m[todo.ID] = *todo
	toDoMap.order = append(toDoMap.order, todo.ID)
}

type IndexData struct {
	Title   string
	Content string
}

type ToDoApp struct {
	toDoList ToDoMap
}

func (toDoApp *ToDoApp) handleGetReq(
	c *gin.Context,
) {
	c.JSON(http.StatusOK, toDoApp.toDoList.m)
}

func (toDoApp *ToDoApp) handlePostReq(
	c *gin.Context,
) {
	var toDo ToDo
	if err := c.ShouldBindJSON(&toDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	toDoApp.toDoList.add(&toDo)
}

func MakeToDoApp() *ToDoApp {
	toDoList := MakeToDoMap()
	toDoApp := new(ToDoApp)
	toDoApp.toDoList = *toDoList
	return toDoApp
}

func main() {
	toDoApp := MakeToDoApp()

	server := gin.Default()
	server.GET("/", toDoApp.handleGetReq)
	server.POST("/", toDoApp.handlePostReq)
	server.Run(":8888")
}
