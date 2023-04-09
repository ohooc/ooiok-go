package single_partten

import (
	"fmt"
	"sync"
)

type Singleton interface {
	Do()
}

type singleton struct {
}

func (s *singleton) Do() {
	fmt.Println("do some thing")
}

var (
	once     sync.Once
	instance *singleton
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

type httpUtils struct {
	BaseUrl string
	Method  string
}

var httpInstance *httpUtils

func GetHttpUtils() *httpUtils {
	once.Do(func() {
		httpInstance = &httpUtils{}
	})
	return httpInstance
}
