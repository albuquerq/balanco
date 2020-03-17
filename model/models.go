package model

import "time"

// Product is the product model.
type Product struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	Type     string `json:"type,omitempty"`
	IsActive bool   `json:"active"`
}

// Unit is the unit model.
type Unit struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	CNES     string `json:"CNES"`
	IsActive bool   `json:"active"`
}

// Employee is the employee model.
type Employee struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	IsActive     bool   `json:"active"`
}

// EmployeeSimple is a simplification of the employee model.
type EmployeeSimple struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
}

// Balance is the balance model.
type Balance struct {
	ID          int64          `json:"id"`
	UnitID      int64          `json:"unitId"`
	Unit        *Unit          `json:"unit,omitempty"`
	Description string         `json:"description"`
	DateInit    time.Time      `json:"initAt"`
	DateEnd     time.Time      `json:"endAt"`
	Items       []*BalanceItem `json:"items,omitempty"`
}

// BalanceItem is the balance item model.
type BalanceItem struct {
	StockID    int64     `json:"stockId"`
	Stock      *Stock    `json:"stock"`
	CheckingID int64     `json:"checkingID"`
	Checking   *Checking `json:"checking,omitempty"`
}

// Stock is the stock model.
type Stock struct {
	BalanceID int64     `json:"balanceId"`
	ProductID int64     `json:"productId"`
	Amount    int       `json:"amount"`
	DateQuery time.Time `json:"queryAt"`
}

// Checking is the checking model.
type Checking struct {
	ID         int64           `json:"id"`
	BalanceID  int64           `json:"balanceId"`
	ProductID  int64           `json:"productId"`
	Product    *Product        `json:"product,omitempty"`
	EmployeeID int64           `json:"employeeId"`
	Employee   *EmployeeSimple `json:"employee,omitempty"`
	Amount     int             `json:"amount"`
	CheckedAt  time.Time       `json:"checkedAt"`
}
