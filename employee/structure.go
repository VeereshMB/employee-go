package employee

import "sync"

type store struct {
	db        map[int]Employee
	mu        *sync.RWMutex
	idCounter int
}

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

var Store *store
