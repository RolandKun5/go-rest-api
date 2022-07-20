package types

type RawUser struct {
	ID                string `json:"id"`
	UserName          string `json:"userName"`
	PermissionLevelID string `json:"permissionLevelId"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	City              string `json:"city"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}

type User struct {
	ID              string `json:"id"`
	UserName        string `json:"userName"`
	PermissionLevel int    `json:"permissionLevel"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	City            string `json:"city"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type PermissionLevel struct {
	ID              string `json:"id"`
	PermissionLevel int    `json:"permissionLevel"`
}
