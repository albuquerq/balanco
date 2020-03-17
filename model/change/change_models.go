package change

import (
	"fmt"
)

type EmployeeChange struct {
	Name         *string `json:"name"`
	ProfileImage *string `json:"profileImage"`
	Email        *string `json:"email"`
	Password     *string `json:"password"`
	Phone        *string `json:"phone"`
	IsActive     *bool   `json:"active"`
}

func Employee() *EmployeeChange {
	return &EmployeeChange{}
}

func (ec *EmployeeChange) SetName(name string) *EmployeeChange {
	if name != "" {
		ec.Name = &name
	}
	return ec
}

func (ec *EmployeeChange) SetProfileImage(profileImage string) *EmployeeChange {
	if profileImage != "" {
		ec.ProfileImage = &profileImage
	}
	return ec
}

func (ec *EmployeeChange) SetEmail(email string) *EmployeeChange {
	if email != "" {
		ec.Email = &email
	}
	return ec
}

func (ec *EmployeeChange) SetPassword(password string) *EmployeeChange {
	if password != "" {
		ec.Password = &password
	}
	return ec
}

func (ec *EmployeeChange) SetPhone(phone string) *EmployeeChange {
	if phone != "" {
		ec.Phone = &phone
	}
	return ec
}

func (ec *EmployeeChange) SetIsActive(active bool) *EmployeeChange {
	ec.IsActive = &active
	return ec
}

func (ec *EmployeeChange) Changes() map[string]interface{} {
	changes := make(map[string]interface{})

	if ec.Name != nil {
		changes["name"] = *ec.Name
	}
	if ec.ProfileImage != nil {
		changes["profileImage"] = *ec.ProfileImage
	}
	if ec.Email != nil {
		changes["email"] = *ec.Email
	}
	if ec.Password != nil {
		changes["password"] = *ec.Password
	}
	if ec.Phone != nil {
		changes["phone"] = *ec.Phone
	}
	if ec.IsActive != nil {
		changes["active"] = *ec.IsActive
	}

	if len(changes) == 0 {
		return nil
	}
	return changes
}

func (ec *EmployeeChange) ChangesByProps(props ...string) ([]interface{}, error) {
	changes := ec.Changes()
	values := make([]interface{}, 0, len(props))

	for _, prop := range props {
		if value, exists := changes[prop]; exists {
			values = append(values, value)
		} else {
			return nil, fmt.Errorf("\"%s\" not found in properties of EmployeeChange", prop)
		}
	}

	return values, nil
}

type UnitChange struct {
	Name     *string `json:"name"`
	CNES     *string `json:"CNES"`
	IsActive *bool   `json:"active"`
}

func Unit() *UnitChange {
	return &UnitChange{}
}

func (uc *UnitChange) SetName(name string) *UnitChange {
	if name != "" {
		uc.Name = &name
	}
	return uc
}

func (uc *UnitChange) SetCNES(cnes string) *UnitChange {
	if cnes != "" {
		uc.CNES = &cnes
	}
	return uc
}

func (uc *UnitChange) SetIsActive(active bool) *UnitChange {
	uc.IsActive = &active
	return uc
}

func (uc *UnitChange) Changes() map[string]interface{} {
	changes := make(map[string]interface{})
	if uc.Name != nil {
		changes["name"] = *uc.Name
	}
	if uc.CNES != nil {
		changes["cnes"] = *uc.CNES
	}
	if uc.IsActive != nil {
		changes["active"] = *uc.IsActive
	}
	return changes
}

func (uc *UnitChange) ChangesByProps(props ...string) ([]interface{}, error) {
	values := make([]interface{}, 0, len(props))
	changes := uc.Changes()

	for _, prop := range props {
		if value, exists := changes[prop]; exists {
			values = append(values, value)
		} else {
			return nil, fmt.Errorf("\"%s\" not found in properties of UnitChange", prop)
		}
	}
	return values, nil
}

type ProductChange struct {
	Code     *string `json:"code"`
	Name     *string `json:"name"`
	Unit     *string `json:"unit"`
	Type     *string `json:"type,omitempty"`
	IsActive *bool   `json:"active"`
}

func Product() *ProductChange {
	return &ProductChange{}
}

func (pc *ProductChange) SetCode(code string) *ProductChange {
	if code != "" {
		pc.Code = &code
	}
	return pc
}

func (pc *ProductChange) SetName(name string) *ProductChange {
	if name != "" {
		pc.Name = &name
	}
	return pc
}

func (pc *ProductChange) SetUnit(unit string) *ProductChange {
	if unit != "" {
		pc.Unit = &unit
	}
	return pc
}

func (pc *ProductChange) SetType(ptype string) *ProductChange {
	if ptype != "" {
		pc.Type = &ptype
	}
	return pc
}

func (pc *ProductChange) SetIsActive(active bool) *ProductChange {
	pc.IsActive = &active
	return pc
}

func (pc *ProductChange) Changes() map[string]interface{} {
	changes := make(map[string]interface{})

	if pc.Code != nil {
		changes["code"] = *pc.Code
	}
	if pc.Name != nil {
		changes["name"] = *pc.Name
	}
	if pc.Unit != nil {
		changes["unit"] = *pc.Unit
	}
	if pc.Type != nil {
		changes["type"] = *pc.Type
	}
	if pc.IsActive != nil {
		changes["active"] = *pc.IsActive
	}

	return changes
}

func (pc *ProductChange) IsZero() bool {
	changes := pc.Changes()
	if len(changes) == 0 {
		return true
	}
	return false
}

func (pc *ProductChange) ChangesByProps(props ...string) ([]interface{}, error) {
	changes := pc.Changes()

	values := make([]interface{}, 0, len(props))

	for _, prop := range props {
		if value, exits := changes[prop]; exits {
			values = append(values, value)
		} else {
			return nil, fmt.Errorf("\"%s\" not found in properties of ProductChange", prop)
		}
	}

	return values, nil
}
