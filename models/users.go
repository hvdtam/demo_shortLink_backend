package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
	"strings"

	_ "github.com/lib/pq"
)

type Users struct {
	Id         int    `orm:"column(id_user);pk;auto"`
	Username   string `orm:"column(username);unique" json:"username"`
	Password   string `orm:"column(password_hash)" json:"-"`
	Email      string `orm:"column(email)" json:"email"`
	Status     int    `orm:"column(status);null" json:"status"`
	LastOnline int    `orm:"column(last_online);null" json:"last_online"`
	CreatedAt  int    `orm:"column(created_at);null" json:"created_at"`
	UpdatedAt  int    `orm:"column(updated_at);null" json:"updated_at"`
}

func (t *Users) TableName() string {
	return "users"
}

type LoginForm struct {
	Username string `form:"username" json:"username" valid:"Required"`
	Password string `form:"password" json:"password" valid:"Required"`
}

type ChangePasswordForm struct {
	Id          int
	OldPassword string `json:"old_password" valid:"Required"`
	NewPassword string `orm:"unique;column(password)" json:"new_password" valid:"Required"`
}

func init() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	orm.RegisterModel(new(Users))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", fmt.Sprintf("postgres://%s:@%s:%s/%s?sslmode=%s",
		beego.AppConfig.DefaultString("dbuser", "vandinhtamhuynh"),
		beego.AppConfig.DefaultString("dbhost", "127.0.0.1"),
		beego.AppConfig.DefaultString("dbport", "5432"),
		beego.AppConfig.DefaultString("dbname", "shortlink"),
		beego.AppConfig.DefaultString("dbsslmode", "disable")))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.Password = string(passwordHash)
	m.Status = 10
	id, err = o.Insert(m)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, errors.New("username or phone number already exists")
		} else if strings.Contains(err.Error(), "invalid input syntax for integer") {
			return 0, errors.New("phone number must be number")
		}
	}
	return
}

func Login(m *LoginForm) (err error) {
	o := orm.NewOrm()
	user := &Users{Username: m.Username}
	err = o.Read(user, "username")
	if err == orm.ErrNoRows {
		return errors.New("username or password is incorrect")
	} else if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password)) == nil {
		return nil
	}
	return errors.New("username not found")
}