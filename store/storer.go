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
	EmployeeSimpleGet(ctx context.Context, employeeID int64) (*model.EmployeeSimple, error)
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
	ProductDelete(ctx context.Context, productID int64) error
	ProductGet(ctx context.Context, productID int64) (*model.Product, error)
	ProductGetByCode(ctx context.Context, productCode string) (*model.Product, error)
	ProductList(ctx context.Context, limit, offset int64) ([]*model.Product, int64, error)
	ProductListByNameLike(ctx context.Context, limit, offset int64) ([]*model.Product, int64, error)
	ProductListInactive(ctx context.Context, limit, offset int64) ([]*model.Product, int64, error)
	ProductListByType(ctx context.Context, ptype string, limit, offset int64) ([]*model.Product, int64, error)
	ProductListByCode(ctx context.Context, pcode string, limit, offset int64) ([]*model.Product, int64, error)

	BalanceSave(ctx context.Context, b *model.Balance) error
	BalanceGet(ctx context.Context, balanceID int64) (*model.Balance, error)
	BalanceUpdate(ctx context.Context, balanceID int64, bc *change.BalanceChange) error
	BalanceList(ctx context.Context, limit, offset int64) ([]*model.Balance, int64, error)
	BalanceDelete(ctx context.Context, balanceID int64) error

	StockSave(ctx context.Context, s *model.StockItem) error
}
