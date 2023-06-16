package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Shortlink_20230615_201530 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Shortlink_20230615_201530{}
	m.Created = "20230615_201530"

	migration.Register("Shortlink_20230615_201530", m)
}

// Run the migrations
func (m *Shortlink_20230615_201530) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE shortlink(id_shortlink serial primary key,long_url TEXT NOT NULL,alias_url TEXT UNIQUE NOT NULL,password TEXT NOT NULL,expire integer DEFAULT NULL, status integer DEFAULT NULL,total_click integer DEFAULT NULL,created_at integer DEFAULT NULL,created_by integer DEFAULT NULL,updated_at integer DEFAULT NULL,updated_by integer DEFAULT NULL)")
}

// Reverse the migrations
func (m *Shortlink_20230615_201530) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE shortlink")
}
