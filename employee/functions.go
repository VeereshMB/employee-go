package employee

import "sync"

func Init() {
	Store = new()
}

func new() (s *store) {
	s = &store{
		mu:        &sync.RWMutex{},
		db:        map[int]Employee{},
		idCounter: 1,
	}
	return
}
