package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func saveToJSONFile() {
	file, err := os.Create("todos.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(TodosList, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

func main() {
	InitializeTodosList()

	addPtr := flag.String("add", "", "Add a new todo")
	listPtr := flag.Bool("list", false, "List all todos")
	deletePtr := flag.Int("delete", 0, "Delete a todo by index")
	togglePtr := flag.Int("toggle", 0, "Mark a todo as completed by index")
	editPtr := flag.Int("edit", 0, "Edit a todo by index")

	flag.Parse()

	switch {
	case *addPtr != "":
		AddToList(*addPtr)
	case *listPtr:
		ListAllTodos()
	case *deletePtr > 0:
		DeleteTodoByIndex(*deletePtr)
	case *togglePtr > 0:
		ToggleTodoByIndex(*togglePtr)
	case *editPtr > 0:
		EditTodoFromList(*editPtr)
	default:
		fmt.Println("Invalid or no flag provided. Available flags:")
		fmt.Println("-add \"content\"")
		fmt.Println("-list")
		fmt.Println("-delete [index]")
		fmt.Println("-toggle [index]")
		fmt.Println("-edit [index]")
	}

	defer saveToJSONFile()
}
