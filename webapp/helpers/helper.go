package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"

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

func WriteResultCSV(k *knot.WebContext, data []byte) {
	k.Writer.Header().Set("Content-Type", `text/csv`)
	k.Writer.Header().Set("Content-Disposition", `attachment; filename="filename.csv"`)

	k.Write(data, http.StatusOK)
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

func Encrypt(text string) (string, error) {
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(HASHKEY))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), err
}

func Decrypt(cryptoText string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(HASHKEY))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), err
}

func TabToSpace(input string) string {
	var result []string

	for _, i := range input {
		switch {
		// all these considered as space, including tab \t
		// '\t', '\n', '\v', '\f', '\r',' ', 0x85, 0xA0
		case unicode.IsSpace(i):
			result = append(result, "   ") // replace tab with space
		case !unicode.IsSpace(i):
			result = append(result, string(i))
		}
	}

	return strings.Join(result, "")
}
