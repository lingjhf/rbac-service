package models

import (
	"rbac-service/tables"
	"rbac-service/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	CreateRoleForm struct {
		Name string `json:"name"`
	}
	CreateRolePermission struct {
		RoleId       string `json:"role_id"`
		PermissionId string `json:"permission_id"`
	}
	UpdateRoleForm map[string]any
	RoleQuery      struct {
		Name  *string `json:"name"`
		Page  uint    `json:"page"`
		Limit uint    `json:"limit"`
	}
	RoleItem struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	RoleList []RoleItem
)

func (f *CreateRoleForm) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Name, validation.Required, validation.Length(1, 32)),
	)
}

func (f *CreateRolePermission) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.RoleId, validation.Required, validation.Length(36, 36)),
		validation.Field(&f.PermissionId, validation.Required, validation.Length(36, 36)),
	)
}

func (f *UpdateRoleForm) Validate() error {
	return validation.Validate(*f, validation.Map(
		validation.Key("name",
			validation.By(utils.IsString),
			validation.Length(1, 32),
		).Optional(),
	).AllowExtraKeys())
}

func (t *RoleList) FormTable(role []*tables.Role) *RoleList {
	for _, item := range role {
		*t = append(*t, RoleItem{Id: item.Id, Name: item.Name})
	}
	return t
}
