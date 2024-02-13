package modules

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	Config "canopyCore/Configuration"
	"strings"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CopyReversedSubStringToString(origString string, startPosFromSource int, endPosFromSource int, destString string, destPosition int) (bool, string) {
	isSuccess := false
	theResult := ""
	origRune := []rune(origString)
	destRune := []rune(destString)

	//fmt.Printf("origString: %s, startPosFromSource: %d, endPosFromSource: %d, origRune: %d\n", origString, startPosFromSource, endPosFromSource, len(origRune))
	//fmt.Printf("destString: %s, destPosition: %d\n", destString, destPosition)

	if startPosFromSource < 0 && endPosFromSource > (len(origRune)) && startPosFromSource <= endPosFromSource {
		fmt.Println("POSs noting is valid bro.")
		isSuccess = false
	} else {
		cutOrigString := string(origRune[startPosFromSource:endPosFromSource])
		reversedCutOrigString := ReverseString(cutOrigString)
		//fmt.Println("cutOrigString: " + cutOrigString + ", reversedCutOrigString: " + reversedCutOrigString)

		depanDestRune := string(destRune[0:destPosition])
		belakangDestRune := string(destRune[destPosition:len(destString)])
		theResult = depanDestRune + reversedCutOrigString + belakangDestRune

		isSuccess = true

		//fmt.Printf("depanRune: %s, belakangRune: %s, completeDestRune: %s\n", depanDestRune, belakangDestRune, theResult)
	}

	return isSuccess, theResult
}

func DeleteSubStringByIndex(theString string, index int, length int) string {
	theRune := []rune(theString)

	theFinalString := string(theRune[0:index]) + string(theRune[index+length:])

	return theFinalString
}

func SimpleStringEncrypt(theString string) (bool, string) {
	byteString := []byte(theString + Config.ConstDefaultSaltEncryption)

	origB64enc := b64.StdEncoding.EncodeToString(byteString)
	destB64enc := origB64enc

	// move reversed string from 2 - 5 to pos 10
	isOK, result := CopyReversedSubStringToString(origB64enc, 2, 5, destB64enc, 10)

	// Replace = to _ in result
	result = strings.Replace(result, "=", "_dRP", 1)

	return isOK, result
}

func SimpleStringEncryptWithSalt(theString string, itsSalt string) (bool, string) {
	byteString := []byte(theString + itsSalt)

	origB64enc := b64.StdEncoding.EncodeToString(byteString)
	destB64enc := origB64enc

	// move reversed string from 2 - 5 to pos 10
	isOK, result := CopyReversedSubStringToString(origB64enc, 2, 5, destB64enc, 10)

	// Replace = to _ in result
	result = strings.Replace(result, "=", "_dRP", 1)

	return isOK, result
}

func SimpleStringDecrypt(theEncString string) string {
	// Remove pos 10, 4 chars
	cleanEncString := DeleteSubStringByIndex(theEncString, 10, 3)
	cleanEncString = strings.Replace(cleanEncString, "_dRP", "=", 1)
	//fmt.Println("cleanES: " + cleanEncString)

	data, err := b64.StdEncoding.DecodeString(cleanEncString)
	if err != nil {
		return ""
	}

	rawData := string(data)
	cleanData := strings.Replace(rawData, Config.ConstDefaultSaltEncryption, "", 1)

	return cleanData
}

func SimpleStringDecryptWithSalt(theEncString string, itsSalt string) string {
	// Remove pos 10, 4 chars
	cleanEncString := DeleteSubStringByIndex(theEncString, 10, 3)
	cleanEncString = strings.Replace(cleanEncString, "_dRP", "=", 1)
	//fmt.Println("cleanES: " + cleanEncString)

	data, err := b64.StdEncoding.DecodeString(cleanEncString)
	if err != nil {
		return ""
	}

	rawData := string(data)
	cleanData := strings.Replace(rawData, itsSalt, "", 1)

	return cleanData
}

func SimpleStringDecryptX(theEncString string) (bool, string) {
	// Remove pos 10, 4 chars
	cleanEncString := DeleteSubStringByIndex(theEncString, 10, 3)
	cleanEncString = strings.Replace(cleanEncString, "_dRP", "=", 1)
	//fmt.Println("cleanES: " + cleanEncString)

	data, err := b64.StdEncoding.DecodeString(cleanEncString)
	if err != nil {
		return false, ""
	}

	rawData := string(data)
	cleanData := strings.Replace(rawData, Config.ConstDefaultSaltEncryption, "", 1)

	return true, cleanData
}

func DoBCryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func EncryptString(theData string) (bool, string) {
	byteString := []byte(theData)
	byteSalt := []byte(Config.ConstDefaultSaltEncryption)

	c, err := aes.NewCipher(byteSalt)
	// if there are any errors, handle them
	if err != nil {
		return false, ""
	}

	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		return false, ""
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return false, ""
	}

	// Wrap with simple encryption
	isOK, simpleEncrypted := SimpleStringEncrypt(b64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, byteString, nil)))
	if isOK == false {
		return false, ""
	} else {
		return true, simpleEncrypted
	}
}

func DecryptStringPassword(theEncryptedStringX string) (bool, string) {
	//Simple decrypt
	simpleDecryptedStr := SimpleStringDecrypt(theEncryptedStringX)

	byteEncString, errD := b64.StdEncoding.DecodeString(simpleDecryptedStr)

	if errD != nil {
		return false, ""
	}

	c, err := aes.NewCipher([]byte(Config.ConstDefaultSaltEncryption))
	if err != nil {
		return false, ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return false, ""
	}

	nonceSize := gcm.NonceSize()
	if len(byteEncString) < nonceSize {
		return false, ""
	}

	nonce, ciphertext := byteEncString[:nonceSize], byteEncString[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, ""
	}

	return true, string(plaintext)
}

func EncryptStringPasswordWithBCrypt(thePlainPassword string) (bool, string) {
	isSuccess := false
	encBPassword := ""

	// BCrypt the plain password
	bcryptedPassword, err := DoBCryptPassword(thePlainPassword)

	if err != nil {
		isSuccess = false
		encBPassword = ""
	} else {
		// Double security - encrypt bcryptedPassword with string encryption
		isEncOK, encBCryptPassword := EncryptString(bcryptedPassword)

		if isEncOK {
			isSuccess = true
			encBPassword = encBCryptPassword
		} else {
			isSuccess = false
			encBPassword = ""
		}
	}

	return isSuccess, encBPassword
}

func CheckBCryptHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateEncBCryptedPassword(theToCheckPassword string, theEncBCryptedPassword string) bool {
	isValid := false

	// Decrypt the encrypted-bcrypted password
	isDecOK, bCryptedPass := DecryptStringPassword(theEncBCryptedPassword)

	if isDecOK {
		// check if theToCheckPassword is matching with bcryptedPass
		isValid = CheckBCryptHashPassword(theToCheckPassword, bCryptedPass)
	} else {
		isValid = false
	}

	return isValid
}
