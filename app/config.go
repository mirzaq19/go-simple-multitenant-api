package app

import "os"

type TenantDB struct {
	Name string
	DSN  string
}

var TenantsDB = map[string]TenantDB{
	"tenant1": {
		Name: "tenant1",
		DSN:  os.Getenv("TENANT1_DSN"),
	},
	"tenant2": {
		Name: "tenant2",
		DSN:  os.Getenv("TENANT2_DSN"),
	}}
