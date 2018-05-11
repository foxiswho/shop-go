package auth

import (
	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/module/db"
	"fmt"
	"github.com/foxiswho/shop-go/module/model"
	"github.com/foxiswho/shop-go/models"
)

type User struct {
	models.User `xorm:"extends"`

	model.Model `gorm:"-" xorm:"-"`

	authenticated bool `form:"-" db:"-" json:"-" xorm:"-"`
}

// GetAnonymousUser should generate an anonymous user_service model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() auth.User {
	return &User{}
}

// Login will preform any actions that are required to make a user_service model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user_service's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user_service.
func (u *User) Logout() {
	// Remove from logged-in user_service's list
	// etc ...
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user_service object from a database model with
// a matching id.
func (u *User) GetById(id interface{}) error {
	//newu:=new(User)
	//newdb.DB().Engine.Id(id).Get(&u)
	_, err := db.DB().Engine.Id(id).Get(u)
	fmt.Println("GetById=>")
	fmt.Println("GetById=>")
	fmt.Println("GetById=>id", id)
	fmt.Println("GetById=>u", u)
	fmt.Println("GetById=>", err)
	//fmt.Println("GetById=>newu",newu)
	// u = newu
	//fmt.Println("GetById=>newu",u)
	if err != nil {
		return err
	}
	//if err := DB().Where("id = ?", id).First(&u).Error; err != nil {
	//	return err
	//}
	log.Debugf("GetUserById USER:", u)
	fmt.Println("GetUserById USER:", u)
	return nil
}

func (u *User) TraceGetUserById(id uint64) *User {
	if s := u.Trace(); s != nil {
		defer s.Finish()
	}

	user := new(User)
	//var count int64
	//db := DB().Where("id = ?", id)
	//if err := Cache(db).First(&user_service).Count(&count).Error; err != nil {
	//	log.Debugf("GetUserById error: %v", err)
	//	return nil
	//}
	ok, err := db.DB().Engine.Where("nickname = ?", "admin").Get(user)
	fmt.Println("TraceGetUserById err:", err)
	fmt.Println("TraceGetUserById :", ok, user)
	return user
}