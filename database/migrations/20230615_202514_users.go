package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Users_20230615_202514 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20230615_202514{}
	m.Created = "20230615_202514"

	migration.Register("Users_20230615_202514", m)
}

// Run the migrations
func (m *Users_20230615_202514) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(id_user serial primary key,username TEXT UNIQUE NOT NULL,password_hash TEXT NOT NULL,email TEXT NOT NULL,status integer DEFAULT NULL,last_online integer DEFAULT NULL,created_at integer DEFAULT NULL,updated_at integer DEFAULT NULL)")
}

// Reverse the migrations
func (m *Users_20230615_202514) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
