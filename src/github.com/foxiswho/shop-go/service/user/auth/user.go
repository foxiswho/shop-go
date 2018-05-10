package auth

import (
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/module/auth"
	"time"
	oldorm "github.com/foxiswho/shop-go/model/orm"
	oldmodel "github.com/foxiswho/shop-go/model"
)

type User struct {
	oldorm.Model `gorm:"-"`

	Id        uint64    `json:"id,omitempty"`
	Nickname  string    `form:"nickname" json:"nickname,omitempty"`
	Password  string    `form:"password" json:"-"`
	Gender    int64     `json:"gender,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty"`

	authenticated bool `form:"-" db:"-" json:"-"`
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() auth.User {
	return &User{}
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

func (u *User) GetById(id interface{}) error {
	//user := &User{}
	//_, err := db.DB().Engine.Id(id).Get(&user)
	//fmt.Println("user",user)
	//u = user
	//if err == nil {
	//	return nil
	//}
	//return err
	if err := oldmodel.DB().Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	return nil
}
func (u *User) GetUserById(id uint64) *User {
	//user := &User{}
	//user.Id = id;
	//_, err := db.DB().Engine.Get(&user)
	//fmt.Println("user",user)
	//if err == nil {
	//	return user
	//}
	user := User{}
	var count int64
	db := oldmodel.DB().Where("id = ?", id)
	if err := oldmodel.Cache(db).First(&user).Count(&count).Error; err != nil {
		//log.Debugf("GetUserById error: %v", err)
		return nil
	}
	return &user
}

func (u *User) TraceGetUserById(id uint64) *User {
	if s := u.Trace(); s != nil {
		defer s.Finish()
	}

	user := User{}
	user.Id = id;
	_, err := db.DB().Engine.Get(user)
	if err == nil {
		return &user
	}
	return nil
}
