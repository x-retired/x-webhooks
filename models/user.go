package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/xiexianbin/webhooks/utils"
)

// User struct.
type User struct {
	UserId        int       `orm:"pk;auto;unique;column(user_id)"`
	UserName      string    `orm:"size(64);column(user_name)"`
	Password      string    `orm:"size(255);column(password)"`
	Email         string    `orm:"size(255);column(email);null;default(null)"`
	Telephone     string    `orm:"size(20);column(telephone);null;default(null)"`
	Role          int       `orm:"column(role);type(int);default(1)"`   // user role：0 admin/1 user
	Status        int       `orm:"column(status);type(int);default(0)"` // user status: 0 disable /1 enable
	CreateTime    time.Time `orm:"type(datetime);column(create_time);auto_now_add"`
	UpdateTime    time.Time `orm:"type(datetime);column(update_time);auto_now_add"`
	LastLoginTime time.Time `orm:"type(datetime);column(last_login_time);null"`
}

// TableName return User table name.
func (m *User) TableName() string {
	return "users"
}

// TableEngine return DB engine.
func (m *User) TableEngine() string {
	return "INNODB"
}

// NewUser return new User Object.
func NewUser() *User {
	return new(User)
}

// Find find user
func (u *User) Find() error {
	o := orm.NewOrm()

	err := o.Read(u)

	if err == orm.ErrNoRows {
		return UserNotExist
	}

	return nil
}

// Login user login.
func (u *User) Login(account string, password string) (*User, error) {
	o := orm.NewOrm()
	User := &User{}
	err := o.QueryTable(u.TableName()).Filter("account", account).Filter("status", 0).One(User)

	if err != nil {
		return User, UserNotExist
	}

	ok := encrypt.Md5Verify(User.Password, password)

	if ok {
		return User, nil
	}

	return User, UserPasswordError
}

// Create a new user.
func (User *User) Create() error {
	o := orm.NewOrm()
	User.Password = encrypt.Md5String(User.Password)

	_, err := o.Insert(User)
	if err != nil {
		return err
	}

	return nil
}

// Update User information.
func (u *User) Update(cols ...string) error {
	o := orm.NewOrm()

	if _, err := o.Update(u, cols...); err != nil {
		return err
	}
	return nil
}

// Delete an User.
func (u *User) Delete() error {
	o := orm.NewOrm()

	if _, err := o.Delete(u); err != nil {
		return err
	}
	return nil
}
