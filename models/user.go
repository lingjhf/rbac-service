package models

import (
	"rbac-service/tables"
	"rbac-service/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	CreateUserForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
		TenantId string `json:"tenant_id"`
	}
	CreateUserRoleForm struct {
		UserId string `json:"user_id"`
		RoleId string `json:"role_id"`
	}
	CreateUserTenantForm struct {
		UserId   string `json:"user_id"`
		TenantId string `json:"tenant_id"`
	}
	UpdateUserForm map[string]any
	UserQuery      struct {
		Username *string `json:"username"`
		Page     uint    `json:"page"`
		Limit    uint    `json:"limit"`
	}
	UserItem struct {
		Id       string `json:"id"`
		Username string `json:"username"`
	}
	UserList []UserItem
)

func (f *CreateUserForm) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Username, validation.Required, validation.Length(1, 32)),
		validation.Field(&f.Password, validation.Required, validation.Length(6, 32)),
		validation.Field(&f.TenantId, validation.Required, validation.Length(1, 36)),
	)
}

func (f *CreateUserRoleForm) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.UserId, validation.Required, validation.Length(36, 36)),
		validation.Field(&f.RoleId, validation.Required, validation.Length(36, 36)),
	)
}

func (f *CreateUserTenantForm) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.UserId, validation.Required, validation.Length(36, 36)),
		validation.Field(&f.TenantId, validation.Required, validation.Length(36, 36)),
	)
}

func (f *UpdateUserForm) Validate() error {
	return validation.Validate(*f, validation.Map(
		validation.Key("username",
			validation.By(utils.IsString),
			validation.Length(1, 32),
		).Optional(),
		validation.Key("password",
			validation.By(utils.IsString),
			validation.Length(6, 32),
		).Optional(),
	).AllowExtraKeys())
}

func (u *UserList) FormTable(user []*tables.User) *UserList {
	for _, item := range user {
		*u = append(*u, UserItem{Id: item.Id, Username: item.Username})
	}
	return u
}
