package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"
)

func WriteResultError(k *knot.WebContext, res *toolkit.Result, errorText string) {
	toolkit.Println(errorText)
	res.SetErrorTxt(errorText)
	k.WriteJSON(res, http.StatusUnauthorized)
}

func WriteResultErrorOK(k *knot.WebContext, res *toolkit.Result, errorText string) {
	toolkit.Println(errorText)
	res.SetErrorTxt(errorText)
	k.WriteJSON(res, http.StatusOK)
}

func WriteResultOK(k *knot.WebContext, res *toolkit.Result, data interface{}) {
	res.SetData(data)
	k.WriteJSON(res, http.StatusOK)
}

func ObjectKeys(data interface{}) []string {
	var datax map[string]interface{}

	dataJson, _ := json.Marshal(data)
	json.Unmarshal(dataJson, &datax)

	keys := make([]string, 0)
	for key := range datax {
		keys = append(keys, key)
	}

	return keys
}

const HASHKEY = "datacatalogue123"

func createHash() string {
	hasher := md5.New()
	hasher.Write([]byte(HASHKEY))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(HASHKEY))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypt(cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(HASHKEY))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
