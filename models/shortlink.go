package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type Shortlink struct {
	Id         int    `orm:"column(id_shortlink);pk"`
	LongUrl    string `orm:"column(long_url)"`
	AliasUrl   string `orm:"column(alias_url)"`
	Password   string `orm:"column(password)"`
	Expire     int    `orm:"column(expire);null"`
	TotalClick int    `orm:"column(total_click);null"`
	CreatedAt  int    `orm:"column(created_at);null"`
	CreatedBy  int    `orm:"column(created_by);null"`
	UpdatedAt  int    `orm:"column(updated_at);null"`
	UpdatedBy  int    `orm:"column(updated_by);null"`
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
	id, err = o.Insert(m)
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
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
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
