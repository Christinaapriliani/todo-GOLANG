package main

import "errors"

type Todo struct {
	ID       int
	Title    string
	Complete bool
}

type TodoList struct {
	todos     map[int]*Todo
	currentID int
}

func NewTodoList() *TodoList {
	return &TodoList{
		todos:     make(map[int]*Todo),
		currentID: 1,
	}
}

func (list *TodoList) Add(title string) {
	list.todos[list.currentID] = &Todo{ID: list.currentID, Title: title}
	list.currentID++
}

func (list *TodoList) Remove(id int) error {
	if _, exists := list.todos[id]; !exists {
		return errors.New("todo not found")
	}
	delete(list.todos, id)
	return nil
}

func (list *TodoList) GetByID(id int) (*Todo, error) {
	if todo, exists := list.todos[id]; exists {
		return todo, nil
	}
	return nil, errors.New("todo not found")
}

func (list *TodoList) GetAll() []*Todo {
	result := make([]*Todo, 0, len(list.todos))
	for _, todo := range list.todos {
		result = append(result, todo)
	}
	return result
}
