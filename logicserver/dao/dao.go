package dao

import (
	"database/sql"
	"github.com/go-xorm/xorm"
	"im/logicserver/mysql"
	"strings"
)

type Dao struct {
	session *xorm.Session
}

func NewDao() *Dao {
	return &Dao{}
}

func ErrorIsTooManyConnections(err error) bool {
	if err == nil {
		return false
	}
	errString := err.Error()

	return strings.Contains(errString, "Error") && strings.Contains(errString, "1040")
}

func ErrorIsDuplicate(err error) bool {
	if err == nil {
		return false
	}
	errString := err.Error()
	return strings.Contains(errString, "Error") && strings.Contains(errString, "1062")
}

// Find retrieve records from table, condiBeans's non-empty fields
// are conditions. beans could be []Struct, []*Struct, map[int64]Struct
// map[int64]*Struct
func (d *Dao) Find(beans interface{}, condiBeans ...interface{}) error {
	if d.session != nil {
		return d.session.Find(beans, condiBeans...)
	}
	engine := mysql.GetXormEngine()

	return engine.Find(beans, condiBeans...)
}

// Get retrieve one record from table, bean's non-empty fields
// are conditions
func (d *Dao) Get(bean interface{}) (bool, error) {
	if d.session != nil {
		return d.session.Get(bean)
	}

	engine := mysql.GetXormEngine()

	return engine.Get(bean)

}

func (d *Dao) Where(query interface{}, args ...interface{}) *Dao {
	engine := mysql.GetXormEngine()

	d.session = engine.Where(query, args...)
	return d
}

// Insert one or more records
func (d *Dao) Insert(beans ...interface{}) (int64, error) {
	if d.session != nil {
		return d.session.Insert(beans)
	}

	engine := mysql.GetXormEngine()
	return engine.Insert(beans...)
}

func (d *Dao) Exec(sqlStr string, args ...interface{}) (sql.Result, error) {
	engine := mysql.GetXormEngine()
	return engine.Exec(sqlStr, args...)
}

// Update records, bean's non-empty fields are updated contents,
// condiBean' non-empty filds are conditions
// CAUTION:
//        1.bool will defaultly be updated content nor conditions
//         You should call UseBool if you have bool to use.
//        2.float32 & float64 may be not inexact as conditions
func (d *Dao) Update(bean interface{}, condiBeans ...interface{}) (int64, error) {
	if d.session != nil {
		return d.session.Update(bean, condiBeans...)
	}

	engine := mysql.GetXormEngine()
	return engine.Update(bean, condiBeans...)
}

// Delete records, bean's non-empty fields are conditions
func (d *Dao) Delete(bean interface{}) (int64, error) {
	if d.session != nil {
		return d.session.Delete(bean)
	}

	engine := mysql.GetXormEngine()
	return engine.Delete(bean)
}
