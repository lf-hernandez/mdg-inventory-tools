package models

import "time"

type User struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	Role        Role        `json:"role"`
	Permissions Permissions `json:"permissions"`
	Inventories []string    `json:"inventories"`
	CreatedAt   time.Time   `json:"created_at"`
	ModifiedAt  time.Time   `json:"modified_at"`
}

type Role string

const (
	Admin      Role = "Admin"
	Sales      Role = "Sales"
	Accounting Role = "Accounting"
)

type PermissionSet map[string]bool

type Permissions struct {
	ItemPermissions      PermissionSet `json:"item_permissions"`
	InventoryPermissions PermissionSet `bool:"inventory_permissions"`
	UserPermissions      PermissionSet `boo:"user_permissions"`
}

var RoleResourcePermissions = map[Role]map[string]PermissionSet{
	Admin: {
		"item": {
			"create": true,
			"read":   true,
			"update": true,
			"delete": true,
		},
		"inventory": {
			"create": true,
			"read":   true,
			"update": true,
			"delete": true,
		},
		"user": {
			"create": true,
			"read":   true,
			"update": true,
			"delete": true,
		},
	},
	Sales: {
		"items": {
			"read": true,
		},
		"inventory": {
			"read": true,
		},
	},
	Accounting: {
		"items": {
			"create": true,
			"read":   true,
			"update": true,
		},
	},
}
