package user

import (
	"time"
	"fmt"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/service/user/auth"
)

type UserService struct {
	id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username    string    `json:"username" xorm:"comment('用户名') index CHAR(30)"`
	Password    string    `json:"password" xorm:"comment('密码') CHAR(32)"`
	Mail        string    `json:"mail" xorm:"comment('邮箱') VARCHAR(80)"`
	Salt        string    `json:"salt" xorm:"comment('干扰码') VARCHAR(10)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Ip          string    `json:"ip" xorm:"comment('添加IP') CHAR(15)"`
	JobNo       string    `json:"job_no" xorm:"comment('工号') VARCHAR(15)"`
	NickName    string    `json:"nick_name" xorm:"comment('昵称') VARCHAR(50)"`
	TrueName    string    `json:"true_name" xorm:"comment('真实姓名') VARCHAR(50)"`
	Qq          string    `json:"qq" xorm:"comment('qq') VARCHAR(50)"`
	Phone       string    `json:"phone" xorm:"comment('电话') VARCHAR(50)"`
	Mobile      string    `json:"mobile" xorm:"comment('手机') VARCHAR(20)"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 comment('删除0否1是') index TINYINT(1)"`
}

func NewUserService() *UserService {
	return new(UserService)
}

func GetUserByNicknamePwd(nickname string, pwd string) *auth.User {
	user := new(auth.User)
	//if err := DB().Where("nickname = ? AND password = ?", nickname, pwd).First(&user).Error; err != nil {
	//	log.Debugf("GetUserByNicknamePwd error: %v", err)
	//	return nil
	//}
	ok, err := db.DB().Engine.Where("nickname = ?", nickname).Get(user)
	fmt.Println("GetUserByNicknamePwd :", ok, user)
	if err != nil {
		fmt.Println("GetUserByNicknamePwd error:", err)
		log.Debugf("GetUserByNicknamePwd error: %v", err)
		return nil
	}
	if user.Password != pwd {
		fmt.Println("user.Password != pwd", user.Password, pwd)
		log.Debugf("user.Password error: %v", pwd)
		return nil
	}
	fmt.Println("GetUserByNicknamePwd xxxxxxx:", user)
	return user
}

func AddUserWithNicknamePwd(nickname string, pwd string) *auth.User {
	user := auth.User{Nickname: nickname, Password: pwd, Birthday: time.Now()}
	if _,err:=db.DB().Engine.Insert(&user); err != nil {
		return nil
	}
	return &user
}

func GetUserById(id uint64) *auth.User {
	user := new(auth.User)
	//var count int64
	//db := DB().Where("id = ?", id)
	//if err := Cache(db).First(&user).Count(&count).Error; err != nil {
	//	log.Debugf("GetUserById error: %v", err)
	//	return nil
	//}

	if _, err := db.DB().Engine.Id(id).Get(user); err != nil {
		log.Debugf("GetUserById error: %v", err)
		return nil
	}
	log.Debugf("GetUserById USER:", user)
	fmt.Println("GetUserById USER:", user)
	return user
}