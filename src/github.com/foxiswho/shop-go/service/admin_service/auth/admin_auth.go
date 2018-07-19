package auth

import (
	"github.com/foxiswho/shop-go/consts/login"
	"fmt"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/util"
	"github.com/foxiswho/shop-go/util/password"
	"github.com/foxiswho/shop-go/util/conv"
)

type AdminAuth struct {
	TypeLogin          int
	Username           string
	Password           string
	VerificationCode   string
	IsVerificationCode bool
}

func NewAdminAuth() *AdminAuth {
	c := new(AdminAuth)
	c.TypeLogin = login.Type_default
	c.IsVerificationCode = false
	return c
}

// 设置登录类型
func (c *AdminAuth) SetTypeLogin(type_login int) {
	c.TypeLogin = type_login
}
func (c *AdminAuth) SetUsername(username string) {
	c.Username = username
}

func (c *AdminAuth) SetPassword(password string) {
	c.Password = password
}

//验证码
func (c *AdminAuth) SetVerificationCode(verificationCode string) {
	c.VerificationCode = verificationCode
}

//是否验证验证码
func (c *AdminAuth) SetIsVerificationCode(isVerificationCode bool) {
	c.IsVerificationCode = isVerificationCode
}

func (c *AdminAuth) Process() (*models.Admin, error) {
	if len(c.Username) < 1 {
		return nil, util.NewError("用户名 不能为空")
	}
	if len(c.Password) < 1 {
		return nil, util.NewError("密码 不能为空")
	}
	if true == c.IsVerificationCode && len(c.VerificationCode) < 1 {
		return nil, util.NewError("验证码 不能为空")
	}
	admin, err := c.loginTypeProcess()
	if err != nil {
		return nil, err
	}
	return admin, nil
}

//登录处理
func (c *AdminAuth) loginTypeProcess() (*models.Admin, error) {
	if login.Type_default == c.TypeLogin {
		user := new(models.Admin)
		ok, err := db.DB().Engine.Where("username = ?", c.Username).Get(user)
		fmt.Println("GetUserByNicknamePwd :", ok, user)
		if err != nil {
			fmt.Println("GetUserByNicknamePwd error:", err)
			log.Debugf("GetUserByNicknamePwd error: %v", err)
			return nil, util.NewError("匹配错误")
		}
		if user.Password != password.SaltMake(c.Password, user.Salt) {
			log.Debugf("user_service.Password error: %v", c.Password)
			return nil, util.NewError("密码 匹配不成功")
		}
		return user, nil
	} else if login.Type_uid == c.TypeLogin {
		user := new(models.Admin)
		ok, err := db.DB().Engine.Id(conv.StrToInt(c.Username)).Get(user)
		fmt.Println("GetUserByNicknamePwd :", ok, user)
		if err != nil {
			fmt.Println("GetUserByNicknamePwd error:", err)
			log.Debugf("GetUserByNicknamePwd error: %v", err)
			return nil, util.NewError("匹配错误")
		}
		if user.Password != password.SaltMake(c.Password, user.Salt) {
			log.Debugf("user_service.Password error: %v", c.Password)
			return nil, util.NewError("密码 匹配不成功")
		}
		return user, nil
	} else if login.Type_mail == c.TypeLogin {
		user := new(models.Admin)
		ok, err := db.DB().Engine.Where("mail = ?", c.Username).Get(user)
		fmt.Println("GetUserByNicknamePwd :", ok, user)
		if err != nil {
			fmt.Println("GetUserByNicknamePwd error:", err)
			log.Debugf("GetUserByNicknamePwd error: %v", err)
			return nil, util.NewError("匹配错误")
		}
		if user.Password != password.SaltMake(c.Password, user.Salt) {
			log.Debugf("user_service.Password error: %v", c.Password)
			return nil, util.NewError("密码 匹配不成功")
		}
		return user, nil
	} else if login.Type_mobile == c.TypeLogin {
		user := new(models.Admin)
		ok, err := db.DB().Engine.Where("mail = ?", c.Username).Get(user)
		fmt.Println("GetUserByNicknamePwd :", ok, user)
		if err != nil {
			fmt.Println("GetUserByNicknamePwd error:", err)
			log.Debugf("GetUserByNicknamePwd error: %v", err)
			return nil, util.NewError("匹配错误")
		}
		if user.Password != password.SaltMake(c.Password, user.Salt) {
			log.Debugf("user_service.Password error: %v", c.Password)
			return nil, util.NewError("密码 匹配不成功")
		}
		return user, nil
	}
	user := new(models.Admin)
	ok, err := db.DB().Engine.Where("username = ?", c.Username).Get(user)
	fmt.Println("GetUserByNicknamePwd :", ok, user)
	if err != nil {
		fmt.Println("GetUserByNicknamePwd error:", err)
		log.Debugf("GetUserByNicknamePwd error: %v", err)
		return nil, util.NewError("匹配错误")
	}
	if user.Password != password.SaltMake(c.Password, user.Salt) {
		log.Debugf("user_service.Password error: %v", c.Password)
		return nil, util.NewError("密码 匹配不成功")
	}
	return user, nil
}
