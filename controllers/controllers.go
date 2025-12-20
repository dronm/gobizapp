// Package controllers
package controllers

import (
	"sync"
)

var (
	UserRoleID   string
	userRoleIDMx sync.Mutex
)

func UserSetRoleID(roleID string) {
	userRoleIDMx.Lock()
	UserRoleID = roleID
	userRoleIDMx.Unlock()
}
