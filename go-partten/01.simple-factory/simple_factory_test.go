package simple_factory

import "testing"

func TestCreateUser(t *testing.T) {
	user := CreateUser(FrontUser).(*User)
	t.Log(user)
}

func TestCreateApi(t *testing.T) {
	api := NewAPI("zh")
	say := api.Say("ooiok")
	t.Log(say)
}

func TestGenerateUser(t *testing.T) {
	admin := GenerateUser(AdminUser)(1, "ooiok").(*Admin)
	t.Log(admin)
}
