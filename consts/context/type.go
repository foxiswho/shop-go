package context

const (
	//用户登录类型
	Context_Type_User  = "user"
	Context_Type_Admin = "admin"
	//会话登录类型,session_type  cookie jwt
	Session_Type_session     = "session_type"
	Session_Type_cookie = "cookie"
	Session_Type_jwt    = "jwt"
)
