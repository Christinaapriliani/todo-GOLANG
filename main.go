package main

import (
	"fmt"
)

func main() {
	todoList := NewTodoList()

	todoList.Add("Buy groceries")
	todoList.Add("Read a book")
	todoList.Add("Go for a run")

	fmt.Println("All Todos:")
	allTodos := todoList.GetAll()
	for _, todo := range allTodos {
		fmt.Printf("[%d] %s - %s\n", todo.ID, todo.Title, getStatus(todo.Complete))
	}

	idToRemove := 2
	err := todoList.Remove(idToRemove)
	if err != nil {
		fmt.Printf("Error while removing todo with ID %d: %s\n", idToRemove, err)
	}

	fmt.Println("\nUpdated Todo List:")
	allTodos = todoList.GetAll()
	for _, todo := range allTodos {
		fmt.Printf("[%d] %s - %s\n", todo.ID, todo.Title, getStatus(todo.Complete))
	}
}

func getStatus(complete bool) string {
	if complete {
		return "Completed"
	}
	return "Incomplete"
}
