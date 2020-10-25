package dto

import (
	"Food/domain"
	"time"
)

// UserDTO godoc
type UserDTO struct {
	ID       uint
	UserID   string
	Name     string
	Password string
	Roles    []domain.UserRole
	Address  string

	// features
	Age        uint8
	Gender     string
	Occupation string
	Long       float32
	Lat        float32
	ZipCode    uint16

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (dto *UserDTO) GetRolesInterface() []interface{} {
	arr := make([]interface{}, len(dto.Roles))

	for i, role := range dto.Roles {
		arr[i] = role
	}

	return arr
}

func (dto *UserDTO) GetRolesStr() []string {
	arr := make([]string, len(dto.Roles))

	for i, role := range dto.Roles {
		arr[i] = role.String()
	}

	return arr
}
