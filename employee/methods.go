package employee

import "sort"

func (s *store) create(name, position string, salary float64) (emp Employee, err error) {
	if name == "" || position == "" || salary == 0 {
		err = ErrMandatoryParamsMissing
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	emp = Employee{
		ID:       s.idCounter,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	s.db[s.idCounter] = emp
	s.idCounter++
	return
}

func (s *store) readByID(id int) (emp Employee, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	emp, ok := s.db[id]
	if !ok {
		err = ErrUserNotFound
		return
	}
	return
}

func (s *store) update(id int, name, position string, salary float64) (emp Employee, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	emp, ok := s.db[id]
	if !ok {
		err = ErrUserNotFound
		return
	}
	if name != "" {
		emp.Name = name
	}
	if position != "" {
		emp.Position = position
	}
	if salary != 0 {
		emp.Salary = salary
	}
	s.db[id] = emp
	return
}

func (s *store) deleteByID(id int) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.db[id]
	if !ok {
		err = ErrUserNotFound
		return
	}
	delete(s.db, id)
	return
}

func (s *store) list(page, pageSize int) (employees []Employee, err error) {
	if page < 1 {
		err = ErrInvalidPageNumber
		return
	}
	if pageSize < 1 {
		err = ErrInvalidLimit
		return
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	start := (page - 1) * pageSize
	end := page * pageSize
	employees = []Employee{}
	for _, emp := range s.db {
		employees = append(employees, emp)
	}
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].ID < employees[j].ID
	})
	if start >= len(employees) {
		err = ErrNoEmployeesFound
		return
	}
	if end > len(employees) {
		end = len(employees)
	}
	return employees[start:end], nil
}
