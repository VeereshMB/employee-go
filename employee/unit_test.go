package employee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeCRUD(t *testing.T) {
	store := new()

	// Create
	employee, err := store.create("John Doe", "Developer", 50000)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, employee)

	// Read
	emp, err := store.readByID(employee.ID)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", emp.Name)
	assert.Equal(t, "Developer", emp.Position)
	assert.Equal(t, 50000.0, emp.Salary)

	// Update
	emp, err = store.update(emp.ID, "Jane Doe", "Manager", 60000)
	assert.NoError(t, err)

	// Check if update was successful
	updatedEmp, err := store.readByID(employee.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", updatedEmp.Name)
	assert.Equal(t, "Manager", updatedEmp.Position)
	assert.Equal(t, 60000.0, updatedEmp.Salary)

	// Delete
	err = store.deleteByID(employee.ID)
	assert.NoError(t, err)

	// Check if delete was successful
	_, err = store.readByID(employee.ID)
	assert.Error(t, err) // Employee should not exist
}

func TestListEmployees(t *testing.T) {
	store := new()

	// Add some employees
	store.create("John Doe", "Developer", 50000)
	store.create("Jane Smith", "Manager", 60000)
	store.create("Michael Johnson", "Designer", 55000)

	// List employees
	employees, err := store.list(1, 2) // Page 1, Limit 2
	assert.NoError(t, err)
	assert.Equal(t, 2, len(employees))
}

func TestGetEmployeeInvalidID(t *testing.T) {
	store := new()

	// Get non-existent employee
	_, err := store.readByID(999)
	assert.Error(t, err)
}

func TestDeleteEmployeeInvalidID(t *testing.T) {
	store := new()

	// Delete non-existent employee
	err := store.deleteByID(999)
	assert.Error(t, err)
}

func TestListEmployeesInvalidPage(t *testing.T) {
	store := new()

	// List employees with invalid page number
	_, err := store.list(-1, 10)
	assert.Error(t, err)
}

func TestListEmployeesInvalidLimit(t *testing.T) {
	store := new()

	// List employees with invalid limit
	_, err := store.list(1, 0)
	assert.Error(t, err)
}
