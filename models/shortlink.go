package models

import (
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"shortlink/helper"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Shortlink struct {
	Id           int    `orm:"column(id_shortlink);pk;auto" json:"id"`
	LongUrl      string `orm:"column(long_url)" json:"longUrl"`
	AliasUrl     string `orm:"column(alias_url);unique" json:"aliasUrl"`
	FullAliasUrl string `orm:"column(full_alias_url);unique" json:"fullAliasUrl"`
	Password     string `orm:"column(password)" json:"password"`
	Status       int    `orm:"column(status)" json:"status"`
	Expire       int    `orm:"column(expire);null" json:"expire"`
	TotalClick   int    `orm:"column(total_click);null" json:"totalClick"`
	CreatedAt    int    `orm:"column(created_at);null" json:"createdAt"`
	CreatedBy    int    `orm:"column(created_by);null" json:"createdBy"`
	CreatedName  string `orm:"-" json:"createdName"`
	UpdatedAt    int    `orm:"column(updated_at);null" json:"updatedAt"`
	UpdatedBy    int    `orm:"column(updated_by);null" json:"-"`
}

func (t *Shortlink) TableName() string {
	return "shortlink"
}

func init() {
	orm.RegisterModel(new(Shortlink))
}

// AddShortlink insert a new Shortlink into database and returns
// last inserted Id on success.
func AddShortlink(m *Shortlink) (id int64, err error) {
	o := orm.NewOrm()
	if m.AliasUrl == "" {
		var aliasUrl = helper.GenerateRandom(10)
		m.AliasUrl = aliasUrl
		m.FullAliasUrl = beego.AppConfig.DefaultString("urlShorten", "http://localhost:3000/s/") + aliasUrl
	} else {
		m.FullAliasUrl = beego.AppConfig.DefaultString("urlShorten", "http://localhost:3000/s/") + m.AliasUrl
	}
	m.CreatedAt = int(time.Now().Unix())
	if m.Expire != 0 {
		m.Expire = m.CreatedAt + m.Expire
	}
	if m.Password != "" {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
		m.Password = string(passwordHash)
	}
	m.Status = 10
	id, err = o.Insert(m)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, errors.New("Alias Link already exists")
		}
	}
	return
}

// GetShortlinkById retrieves Shortlink by Id. Returns error if
// Id doesn't exist
func GetShortlinkById(id int) (v *Shortlink, err error) {
	o := orm.NewOrm()
	v = &Shortlink{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetShortlinkById retrieves Shortlink by Id. Returns error if
// Id doesn't exist
func GetShortlinkByAlias(alias string) (v *Shortlink, err error) {
	o := orm.NewOrm()
	v = &Shortlink{AliasUrl: alias}
	if err = o.Read(v, "alias_url"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShortlink retrieves all Shortlink matches certain condition. Returns empty list if
// no records exist
func GetAllShortlink(query map[string]string, fields []string, sortby []string, order []string,

	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Shortlink))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Shortlink
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateShortlink updates Shortlink by Id and returns error if
// the record to be updated doesn't exist
func UpdateShortlinkById(m *Shortlink) (err error) {
	o := orm.NewOrm()
	v := Shortlink{Id: m.Id}
	if m.AliasUrl == "" {
		var aliasUrl = helper.GenerateRandom(10)
		m.AliasUrl = aliasUrl
		m.FullAliasUrl = beego.AppConfig.DefaultString("urlShorten", "http://localhost:3000/s/shortlink/") + aliasUrl
	}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateShortlink updates Shortlink by Id and returns error if
// the record to be updated doesn't exist
func UpdateShortlinkByAlias(m *Shortlink) (err error) {
	o := orm.NewOrm()
	idShortlink := 0
	passwordJson := m.Password

	if err = o.Read(m, "alias_url"); err == nil {
		idShortlink = m.Id
	}
	v := Shortlink{Id: idShortlink, AliasUrl: m.AliasUrl}
	if m.AliasUrl == "" {
		var aliasUrl = helper.GenerateRandom(10)
		m.AliasUrl = aliasUrl
		m.FullAliasUrl = beego.AppConfig.DefaultString("urlShorten", "http://localhost:3000/s/shortlink/") + aliasUrl
	}
	if passwordJson != "" {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(passwordJson), bcrypt.DefaultCost)
		m.Password = string(passwordHash)
	}
	// ascertain id exists in the database
	if err = o.Read(&v, "alias_url"); err == nil {
		var num int64
		if num, err = o.Update(m, "alias_url", "password"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteShortlink deletes Shortlink by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShortlink(id int) (err error) {
	o := orm.NewOrm()
	v := Shortlink{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Shortlink{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
