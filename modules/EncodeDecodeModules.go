package modules

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func EncryptQRCODE(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func DecryptQRCODE(data []byte, passphrase string) ([]byte, bool) {
	status := true
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Error : ", err)
		status = false
	}
	return plaintext, status
}

func toSha256() {
	hash := SHA256("136f30ad5ae3dfeedbc26d0f2eae4f76d3863706cdca3548f6d05c85cf63f0eac33e1e340f83b36250fb2228b84474dc375dcb00bc591a2b2489b0e39136c4640a6c1a64250114d733462ebb882302b03d50a119a0b6d24d23269a0f949d69f7ba68e9de428f7e28e2430de0818bf7df991ed565ba4b415081fcc6813de22a9e516cfb473e1ed71f24bf9d46124716847dd503947311bad75b486b7b823f519f8731e190f457bc6d68b24f4f0c7e6033e90470fd41151ba75f5b3b4315765e39b49d9ede4f79a7b32462c51ac824308ef6ba3641b16574f3c3e02eb00288ae576e28d99f5334f02ccea6684e7bacea8936aa76b0d84f5b3a3f04e22db2406ec2adb5042a5189ab338a3d59245d3bcbf13b80124596f2543f77f63a9329b9ebc37b3dc39300c76fd284efa5210342603306923bf85b00d2740545f8e793a49db9227a55cfec68fb9ec60f655a0424ead958c6afdd25e6b28d9c8af68579c800b7e7b65c1c73b0dc410bdd2913f8a0ac0189912603495789ed7585303bd5876bad8d836b67bd23df4fd29cce1680c825dad287c70811ba82ffa4d0fc12498b99ca029b95e0d5cd70401ca3ddf61849e3f78ab5d1d3d239b121d0f7be3f8e7257a887173678513ba12fc959a5638e71d759")
	fmt.Println(hash)
}

// Source : https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/