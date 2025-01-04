package models

import (
	"fmt"

	"backend/internal/db"
	"backend/models/sys"
)

func Migration() {
	fmt.Println(db.DB.AutoMigrate(new(sys.Menu)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.Admins)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.RoleMenu)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.Role)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.AdminsRole)).Error)
}
