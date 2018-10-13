package createEncrypt

import (
	"math/rand"
	"time"
)

const PasswordLength = 256

type Password [PasswordLength]byte

func init()  {
	rand.Seed(time.Now().Unix())
}
// password可以取地址创建：Password为一个array而不是slice
func CreateEncrypt() *Password {
	intArr := rand.Perm(256)
	rand.Seed(6)
	password := &Password{}
	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			return CreateEncrypt()
		}
	}
	return password
}
