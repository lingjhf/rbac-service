package models

import (
	"rbac-service/tables"
	"rbac-service/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	CreatePermissionForm struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		Description string `json:"description"`
	}
	UpdatePermissionForm map[string]any
	PermissionQuery      struct {
		Name  *string `json:"name"`
		Code  *string `json:"code"`
		Page  uint    `json:"page"`
		Limit uint    `json:"limit"`
	}
	PermissionItem struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}
	PermissionList []PermissionItem
)

func (f *CreatePermissionForm) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(&f.Code, validation.Required, validation.Length(1, 64)),
		validation.Field(&f.Description, validation.Length(1, 128)),
	)
}

func (f *UpdatePermissionForm) Validate() error {
	return validation.Validate(*f, validation.Map(
		validation.Key("name",
			validation.By(utils.IsString),
			validation.Length(1, 64),
		).Optional(),
		validation.Key("code",
			validation.By(utils.IsString),
			validation.Length(1, 64),
		).Optional(),
		validation.Key("description",
			validation.By(utils.IsString),
			validation.Length(1, 128),
		).Optional(),
	).AllowExtraKeys())
}

func (t *PermissionList) FormTable(permission []*tables.Permission) *PermissionList {
	for _, item := range permission {
		*t = append(*t, PermissionItem{
			Id:   item.Id,
			Name: item.Name,
			Code: item.Code,
		})
	}
	return t
}
