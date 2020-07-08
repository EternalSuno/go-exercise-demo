package src

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func encrypt() {
	//encode
	hello := "你好, 世界! hello world"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)

	//decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}
	fmt.Println(string(enbyte))

	//高级加解密
	//Go 语言的 crypto 里面支持对称加密的高级加解密包有：
	//crypto/aes 包：AES (Advanced Encryption Standard)，又称 Rijndael 加密法，是美国联邦政府采用的一种区块加密标准。
	//crypto/des 包：DES (Data Encryption Standard)，是一种对称加密标准，是目前使用最广泛的密钥系统，
	//特别是在保护金融数据的安全中。曾是美国联邦政府的加密标准，但现已被 AE S 所替代。

}
