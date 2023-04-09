package simple_factory

// 专门定义一个类来创建其它类的实例

const (
	FrontUser = iota
	AdminUser
)

type UserType int

func CreateUser(t UserType) interface{} {
	switch t {
	case FrontUser:
		return NewUser()
	case AdminUser:
		return NewAdmin()
	default:
		return NewUser()
	}
}

func GenerateUser(t UserType) UserGenFunc {
	switch t {
	case FrontUser:
		return GenUser()
	case AdminUser:
		return GenAdmin()
	default:
		return GenUser()
	}
}
