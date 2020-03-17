package store

import (
	"context"

	"github.com/albuquerq/balanco/model"
	"github.com/albuquerq/balanco/model/change"
)

// Storer interface to store model data.
type Storer interface {
	EmployeeSave(ctx context.Context, employee *model.Employee) error
	EmployeeGet(ctx context.Context, employeeID int64) (*model.Employee, error)
	EmployeeList(ctx context.Context, limit, offset int64) ([]*model.Employee, int64, error)
	EmployeeUpdate(ctx context.Context, employeeID int64, employeeChange change.EmployeeChange) error
	EmployeeDeactivate(ctx context.Context, employeeID int64) error
	EmployeeDelete(ctx context.Context, employeeID int64) error
}
