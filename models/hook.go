package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// hook struct.
type Hook struct {
	HookId         int       `orm:"pk;auto;unique;column(hook_id)" json:"hook_id"`
	HookType       string    `orm:"column(hook_type);size(50)" json:"hook_type"`
	RepositoryName string    `orm:"size(255);column(repo_name)" json:"repository_name"`
	BranchName     string    `orm:"size(255);column(branch_name)" json:"branch_name"`
	Shell          string    `orm:"size(1000);column(shell)" json:"shell"`
	Status         int       `orm:"type(int);column(status);default(0)" json:"status"`
	Secret         string    `orm:"size(255);column(secret)" json:"secret"`
	LastExecTime   time.Time `orm:"type(datetime);column(last_exec_time);null" json:"last_exec_time"`
	CreateTime     time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"create_time"`
	UpdateTime     time.Time `orm:"type(datetime);column(update_time);auto_now_add" json:"update_time"`
}

// TableName return Server table name.
func (m *Hook) TableName() string {
	return "hooks"
}

// TableEngine return DB engine.
func (m *Hook) TableEngine() string {
	return "INNODB"
}

// NewHook return new Hook Object.
func NewHook() *Hook {
	return &Hook{}
}

// Find
func (m *Hook) Find() error {
	if m.HookId <= 0 {
		return InvalidParameter
	}

	o := orm.NewOrm()
	if err := o.Read(m); err != nil {
		return err
	}
	return nil
}

// batch delete
func (m *Hook) DeleteMulti(id ...int) error {
	if len(id) > 0 {
		o := orm.NewOrm()
		ids := make([]int, len(id))
		params := ""

		for i := 0; i < len(id); i++ {
			ids[i] = id[i]
			params = params + ",?"
		}
		_, err := o.Raw("DELETE hooks WHERE hook_id IN ("+params[1:]+")", ids).Exec()
		if err != nil {
			return err
		}
		return nil
	}
	return InvalidParameter
}

// one delete
func (m *Hook) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(m)

	return err
}

// find by key
func (m *Hook) FindByKey(key string) error {
	o := orm.NewOrm()

	if err := o.QueryTable(m.TableName()).Filter("key", key).One(m); err != nil {
		return err
	}
	return nil
}

// add or create hook
func (m *Hook) Save() error {
	o := orm.NewOrm()
	var err error

	if m.HookId > 0 {
		_, err = o.Update(m)
	} else {
		_, err = o.Insert(m)
	}

	return err
}
