package apperror

// 0 表示成功
// 9999 未定义通用错误

// 100-600 http系统级错误范围
// 1000-1500 mysql 错误范围 ， 具体参考 mysql.MySQLError 和 mysql 错误码
// 1999 redis 错误码，redis包未定义错误码, redis.Error = errors.New("xxx"), 因此接口返回：{code:1999, message:redis.Error}

// 3000-4000 handler 校验参数错误范围，各 handler 错误码可以重复。
// 4000-5000 service 错误范围，各 service 错误码可以重复。
// 5000-6000 model 错误范围，各 model 错误码可以重复。
var (
	SUCCESS = AppError{0, "success"}

	ErrUserNotLogin = AppError{3001, "user not login"}

	IdEmpty     = AppError{5001, "id empty"}
	PkListEmpty = AppError{5002, "id list empty"}
)
