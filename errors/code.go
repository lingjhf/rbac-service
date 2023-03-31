package errors

var (
	ErrUnknown      = NewError("", "未知错误")
	ErrSucceeded    = NewError("", "成功")
	ErrBodyParser   = NewError("", "请求体解析错误")
	ErrQueryParser  = NewError("", "查询参数解析错误")
	ErrParameter    = NewError("", "参数错误")
	ErrDatabase     = NewError("", "数据库操作错误")
	ErrSignup       = NewError("", "注册失败")
	ErrSignin       = NewError("", "登录失败")
	ErrUnauthorized = NewError("", "未授权")
	ErrTokenExpired = NewError("", "令牌过期")
)
