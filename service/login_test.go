package service

import (
	"testing"

	"github.com/lcsin/go-itools/service/login"
)

func TestPasswdLogin(t *testing.T) {
	build := login.NewILoginBuilder().
		WithEmail("1847@qq.com").
		WithInputPasswd("1234").Build(login.Passwd)
	if err := build.Login(); err != nil {
		t.Fatalf("登录失败：%v", err)
	}
	t.Log("登录成功")
}

func TestPhoneLogin(t *testing.T) {
	build := login.NewILoginBuilder().
		WithPhone("187xxxxx").
		WithInputCode("123456").Build(login.PhoneCode)
	if err := build.Login(); err != nil {
		t.Fatalf("登录失败：%v", err)
	}
	t.Log("登录成功")
}
