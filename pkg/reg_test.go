package pkg

import (
	"testing"
)

func TestOnlyNumberReg(t *testing.T) {
	println(OnlyNumber.MatchString(""))            // false
	println(OnlyNumber.MatchString("123123123"))   // true
	println(OnlyNumber.MatchString("sdf2131sf"))   // false
	println(OnlyNumber.MatchString("123sdl123fs")) // false
	println(OnlyNumber.MatchString("-123123123"))  // false
}

func TestContainChineseReg(t *testing.T) {
	println(ContainChinese.MatchString("你好，你叫什么"))       // true
	println(ContainChinese.MatchString("hello, world!")) // false
	println(ContainChinese.MatchString("hello,你好!"))     // true
	println(ContainChinese.MatchString("你好，world!"))     // true
}

func TestEmail(t *testing.T) {
	println(Email.MatchString("1847@qq.com")) // true
	println(Email.MatchString("1847xqq.com")) // false
	println(Email.MatchString("1847@qq.xxx")) // true
}

func TestPhone(t *testing.T) {
	println(Phone.MatchString("18758387692")) // true
	println(Phone.MatchString("11111111111")) // false
	println(Phone.MatchString("12345678901")) // false
}
