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
