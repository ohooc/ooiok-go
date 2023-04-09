package single_partten

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	t.Logf("instance1: %p\n", instance1)
	instance2 := GetInstance()
	t.Logf("instance1: %p\n", instance2)
}

func TestGetHttpUtils(t *testing.T) {
	instance := GetHttpUtils()
	instance.BaseUrl = "127.0.0.1"
	instance.Method = "GET"
	t.Log(instance)
}
