package demo1

import (
	"fmt"
	"os"
	"testing"
)

func before() {
	fmt.Println("before")
}

func after() {
	fmt.Println("after")
}

func Test1(t *testing.T) {
	fmt.Println("test1")
}

func Test2(t *testing.T) {
	fmt.Println("test2")
}

func TestMain(m *testing.M) {
	before()
	r := m.Run()
	after()
	os.Exit(r)
}
