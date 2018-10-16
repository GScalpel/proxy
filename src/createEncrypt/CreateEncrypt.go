package createEncrypt

import (
	"math/rand"
)

const PasswordLength = 256

type Password [PasswordLength]byte

// password可以取地址创建：Password为一个array而不是slice
func CreateEncrypt(seed int64) (*Password, *Password) {
	rand.Seed(seed)
	pwd, upwd := generate()
	return pwd, upwd
}

func generate() (*Password, *Password) {
	password := new(Password)
	unPassword := new(Password)
	intArr := rand.Perm(PasswordLength)
	for k, v := range intArr {
		password[k] = byte(v)
		unPassword[v] = byte(k)
		if k == v {
			return generate()
		}
	}
	return password, unPassword
}