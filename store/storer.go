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
	EmployeeUpdate(ctx context.Context, employeeID int64, ec *change.EmployeeChange) error
	EmployeeDeactivate(ctx context.Context, employeeID int64) error
	EmployeeDelete(ctx context.Context, employeeID int64) error

	UnitSave(ctx context.Context, unit *model.Unit) error
	UnitGet(ctx context.Context, unitID int64) (*model.Unit, error)
	UnitList(ctx context.Context, limit, offset int64) ([]*model.Unit, int64, error)
	UnitUpdate(ctx context.Context, unitID int64, uc *change.UnitChange) error

	ProductSave(ctx context.Context, p *model.Product) error
	ProductUpdate(ctx context.Context, productID int64, pc *change.ProductChange) error
	ProductList(ctx context.Context, limit, offset int64) ([]*model.Product, int64, error)
	ProductListByNameLike(ctx context.Context, limit, offset int64) ([]*model.Product, int64, error)
	ProductListInactive(ctx context.Context, limit, offset int64) ([]*model.Product, int64, error)
	ProductListByType(ctx context.Context, ptype string, limit, offset int64) ([]*model.Product, int64, error)
}
