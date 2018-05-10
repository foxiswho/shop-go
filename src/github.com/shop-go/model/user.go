package model

import (
	"time"

	"github.com/hb-go/echo-web/model/orm"
	"github.com/hb-go/echo-web/module/auth"
	"github.com/hb-go/echo-web/module/log"
)

func (u *User) TraceGetUserById(id uint64) *User {
	if s := u.Trace(); s != nil {
		defer s.Finish()
	}

	user := User{}
	var count int64
	db := DB().Where("id = ?", id)
	if err := Cache(db).First(&user).Count(&count).Error; err != nil {
		log.Debugf("GetUserById error: %v", err)
		return nil
	}

	return &user
}

func (u *User) GetUserById(id uint64) *User {
	user := User{}
	var count int64
	db := DB().Where("id = ?", id)
	if err := Cache(db).First(&user).Count(&count).Error; err != nil {
		log.Debugf("GetUserById error: %v", err)
		return nil
	}

	return &user
}

func (u *User) GetUserByNicknamePwd(nickname string, pwd string) *User {
	user := User{}
	if err := DB().Where("nickname = ? AND password = ?", nickname, pwd).First(&user).Error; err != nil {
		log.Debugf("GetUserByNicknamePwd error: %v", err)
		return nil
	}
	return &user
}

func (u *User) AddUserWithNicknamePwd(nickname string, pwd string) *User {
	user := User{Nickname: nickname, Password: pwd, Birthday: time.Now()}

	if err := DB().Create(&user).Error; err != nil {
		return nil
	}
	return &user
}

type User struct {
	orm.Model `gorm:"-"`

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

func (u User) TableName() string {
	return "user"
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

// GetById will populate a user object from a database model with
// a matching id.
func (u *User) GetById(id interface{}) error {
	if err := DB().Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	return nil
}
