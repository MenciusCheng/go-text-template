package mock

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// 批量生成手机号
func Test_generateRandomString_phone(t *testing.T) {
	n := 100
	rand.Seed(time.Now().UnixNano())

	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		v := fmt.Sprintf("\"1%s\"", generateRandomString(10))
		res = append(res, v)
	}

	fmt.Println(strings.Join(res, ","))
}

// 批量生成身份证(只看长度）
func Test_generateRandomString_idCard(t *testing.T) {
	n := 100
	rand.Seed(time.Now().UnixNano())

	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		v := fmt.Sprintf("\"440%s\"", generateRandomString(15))
		res = append(res, v)
	}

	fmt.Println(strings.Join(res, ","))
}

// 批量生成账号ID(只看长度）
func Test_generateRandomString_accountId(t *testing.T) {
	n := 100
	rand.Seed(time.Now().UnixNano())

	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		v := fmt.Sprintf("1%s", generateRandomString(6))
		res = append(res, v)
	}

	fmt.Println(strings.Join(res, ","))
}
