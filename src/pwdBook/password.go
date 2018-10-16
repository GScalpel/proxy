package main

import (
	"createEncrypt"
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type PWDBook map[string][256]byte

func main()  {
	num := *flag.Int("num",1,"need password book's number")
	flag.Parse()
	seed := time.Now().Unix()
	savePWD := make(PWDBook, num)
	for i := num; i > 0; i-- {
		encrypt, unencrypt := createEncrypt.CreateEncrypt(seed + int64(i))
		keyL := "local" + strconv.Itoa(i)
		keyS := "server" + strconv.Itoa(i)
		savePWD[keyL], savePWD[keyS] = *encrypt, *unencrypt
	}
	file, err := os.OpenFile("pwdBook.gob", os.O_CREATE|os.O_WRONLY,0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	enc.Encode(savePWD)
	fmt.Println(savePWD)
}
