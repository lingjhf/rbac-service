package models

import (
	"rbac-service/tables"
	"rbac-service/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	CreateTenantForm struct {
		Name     string  `json:"name"`
		ParentId *string `json:"parentId"`
	}
	UpdateTenantForm map[string]any
	TenantQuery      struct {
		Name  *string `json:"name"`
		Page  uint    `json:"page"`
		Limit uint    `json:"limit"`
	}
	TenantItem struct {
		Id       string  `json:"id"`
		Name     string  `json:"name"`
		ParentId *string `json:"parentId"`
	}
	TenantList []TenantItem
)

func (f *CreateTenantForm) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Name, validation.Required, validation.Length(1, 32)),
		validation.Field(&f.ParentId,
			validation.When(f.ParentId != nil, validation.Length(36, 36)).Else(validation.Nil),
		),
	)
}

func (f *UpdateTenantForm) Validate() error {
	return validation.Validate(*f, validation.Map(
		validation.Key("name", validation.By(utils.IsString), validation.Length(1, 32)).Optional(),
		validation.Key("parentId",
			validation.When((*f)["parentId"] != nil, validation.By(utils.IsString), validation.Length(36, 36)).
				Else(validation.Nil),
		).Optional(),
	).AllowExtraKeys())
}

func (t *TenantList) FormTable(tenant []*tables.Tenant) *TenantList {
	for _, item := range tenant {
		*t = append(*t, TenantItem{Id: item.Id, Name: item.Name, ParentId: item.ParentId})
	}
	return t
}
