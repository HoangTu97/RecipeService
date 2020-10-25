package domain

import (
	"Food/helpers/constants"
	"database/sql/driver"
	"fmt"
)

type UserRole int

const (
	ROLE_USER UserRole = iota
	ROLE_ADMIN
)

var _roleNames = []string{
	constants.ROLE_USER,
	constants.ROLE_ADMIN,
}

func (role UserRole) String() string {
	return _roleNames[role]
}

func (role UserRole) MarshalText() []byte {
	return []byte(role.String())
}

func (role *UserRole) UnmarshalText(text []byte) error {
	str := string(text)
	_role, err := ParseUserRole(str)
	if err != nil {
		return err
	}
	*role = _role
	return nil
}

func (role *UserRole) Scan(value interface{}) error {
	_role, err := ParseUserRole(value.(string))
	if err != nil {
		return err
	}
	*role = _role
	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	return role.String(), nil
}

func ParseUserRole(value string) (UserRole, error) {
	for i, v := range _roleNames {
		if v == value {
			return UserRole(i), nil
		}
	}
	return ROLE_USER, fmt.Errorf("%s is not a valid UserRole", value)
}
