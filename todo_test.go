package main

import (
	"testing"
)

func TestTodoList_Add(t *testing.T) {
	list := NewTodoList()

	list.Add("Buy groceries")

	if len(list.todos) != 1 {
		t.Errorf("Expected 1 todo, but got %d", len(list.todos))
	}
}

func TestTodoList_Remove(t *testing.T) {
	list := NewTodoList()

	list.Add("Buy groceries")
	list.Add("Read a book")
	list.Remove(1)

	if len(list.todos) != 1 {
		t.Errorf("Expected 1 todo, but got %d", len(list.todos))
	}
}

func TestTodoList_GetByID(t *testing.T) {
	list := NewTodoList()

	list.Add("Buy groceries")
	list.Add("Read a book")

	todo, err := list.GetByID(2)
	if err != nil || todo.Title != "Read a book" {
		t.Errorf("Expected todo with title 'Read a book', but got %+v", todo)
	}
}

func TestTodoList_GetAll(t *testing.T) {
	list := NewTodoList()

	list.Add("Buy groceries")
	list.Add("Read a book")

	allTodos := list.GetAll()
	if len(allTodos) != 2 {
		t.Errorf("Expected 2 todos, but got %d", len(allTodos))
	}
}
