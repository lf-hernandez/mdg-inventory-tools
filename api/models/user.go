package models

import "time"

type Role string

const (
	Admin    Role = "Admin"
	Auditor  Role = "Auditor"
	Employee Role = "Employee"
)

type PermissionSet map[string]bool

var DefaultPermissions = PermissionSet{
	"create": true,
	"read":   true,
	"update": true,
	"delete": true,
}

type Permissions struct {
	ItemPermissions      PermissionSet `json:"item_permissions"`
	InventoryPermissions PermissionSet `bool:"inventory_permissions"`
	UserPermissions      PermissionSet `boo:"user_permissions"`
}

type InventoryAccess struct {
	InventoryId string `json:"inventory_id"`
	HasAccess   string `json:"has_access"`
}

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
