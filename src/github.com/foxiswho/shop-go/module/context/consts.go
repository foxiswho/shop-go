package context

const (
	//用户登录类型
	Type_User  = "user"
	Type_Admin = "admin"
	//会话登录类型,session  cookie jwt
	session_session = "session"
	session_cookie  = "cookie"
	session_jwt     = "jwt"
)
