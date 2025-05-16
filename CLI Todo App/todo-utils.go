package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Todo struct {
	CreatedAt     time.Time
	CurrentStatus string
	Content       string
	Index         int
}

var TodosList []Todo
var Reader = bufio.NewReader(os.Stdin)

func InitializeTodosList() {
	if _, err := os.Stat("./todos.json"); os.IsNotExist(err) {
		TodosList = []Todo{}
		return
	}

	dataBytes, err := os.ReadFile("./todos.json")
	if err != nil || len(dataBytes) == 0 {
		TodosList = []Todo{}
		return
	}

	err = json.Unmarshal(dataBytes, &TodosList)
	if err != nil {
		TodosList = []Todo{}
	}
}

func AddToList(content string) {
	newTodo := Todo{
		Content:       content,
		CreatedAt:     time.Now(),
		CurrentStatus: "Pending",
		Index:         len(TodosList) + 1,
	}

	TodosList = append(TodosList, newTodo)
	fmt.Println("Todo added successfully")
}

func EditTodoFromList(Index int) {
	if Index < 1 || Index > len(TodosList) {
		fmt.Println("Invalid index")
		return
	}

	fmt.Println("Enter the new content")
	newContent, _ := Reader.ReadString('\n')
	TodosList[Index-1].Content = strings.TrimSpace(newContent)
	fmt.Println("The todo was updated successfully")
}

func ListAllTodos() {
	if len(TodosList) == 0 {
		fmt.Println("Todo list is empty. Please add todos")
		return
	}

	fmt.Println("\n===== YOUR TODO LIST =====")
	fmt.Printf("%-5s | %-19s | %-10s | %s\n", "ID", "Created At", "Status", "Content")
	fmt.Println(strings.Repeat("-", 80))

	for _, todo := range TodosList {
		timeStr := todo.CreatedAt.Format("2006-01-02 15:04")
		content := strings.TrimSpace(todo.Content)
		if len(content) > 40 {
			content = content[:37] + "..."
		}
		fmt.Printf("%-5d | %-19s | %-10s | %s\n", todo.Index, timeStr, todo.CurrentStatus, content)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("Total: %d item(s)\n\n", len(TodosList))
}

func DeleteTodoByIndex(Index int) {
	if Index < 1 || Index > len(TodosList) {
		fmt.Println("Invalid index")
		return
	}

	Index--
	TodosList = slices.Delete(TodosList, Index, Index+1)

	for i := Index; i < len(TodosList); i++ {
		TodosList[i].Index = i + 1
	}

	fmt.Println("Todo deleted successfully")
}

func ToggleTodoByIndex(index int) {
	if index < 1 || index > len(TodosList) {
		fmt.Println("Invalid index")
		return
	}

	todo := &TodosList[index-1]
	todo.CurrentStatus = "Completed"
	fmt.Println("Todo marked as completed.")
}
