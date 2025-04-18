package main

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos []Todo
var nextID = 1
