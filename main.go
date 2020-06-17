package main

import (
	"encoding/hex"
	"fmt"
	"github.com/fpe-sm4/ff1"
	"github.com/fpe-sm4/ff3"
)

func main()  {
	radix :=2
	maxLen :=8
	key, err := hex.DecodeString("EF4359D8D580AA4F7F036D6F04FC6A94")
	tweak, err := hex.DecodeString("D8E7920AFA330A73")
	plaintext1 := "01001111"
	fmt.Println("**************** test for ff1 ****************")
	f1, err := ff1.NewCipher(radix, maxLen, key, tweak)
	if err != nil {
		fmt.Println("Unable to create cipher: %v", err)
	}
	ciphertext, err := f1.EncryptWithTweak(plaintext1,tweak)
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Println(ciphertext)
	tmp,err := f1.DecryptWithTweak(ciphertext,tweak)
	fmt.Println(tmp)


	fmt.Println("**************** test for ff3 ****************")
	plaintext3 := "01001111"
	f3, err := ff3.NewCipher(radix,  key, tweak)
	if err != nil {
		fmt.Println("Unable to create cipher: %v", err)
	}

	ciphertext3, err := f3.EncryptWithTweak(plaintext3,tweak)
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Println(ciphertext3)
	tmp3,err := f3.DecryptWithTweak(ciphertext3,tweak)
	fmt.Println(tmp3)






}
