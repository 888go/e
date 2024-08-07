package e

import (
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/crypto/gcrc32"
	"github.com/gogf/gf/v2/crypto/gdes"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/util/gconv"
)

func Aes加密(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return gaes.EncryptCBC(plainText, key, iv...)
}
func Aes解密(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return gaes.DecryptCBC(cipherText, key, iv...)
}
func AesCFB加密(plainText []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error) {
	return gaes.EncryptCFB(plainText, key, padding, iv...)
}
func AesCFB解密(cipherText []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error) {
	return gaes.DecryptCFB(cipherText, key, unPadding, iv...)
}
func Crc32加密(v interface{}) uint32 {
	return gcrc32.Encrypt(v)
}

// EncryptECB encrypts <plainText> using ECB mode.
func DesECB加密(plainText []byte, key []byte, padding int) ([]byte, error) {
	return gdes.EncryptECB(plainText, key, padding)
}

// EncryptECB encrypts <plainText> using ECB mode.
func DesECB解密(cipherText []byte, key []byte, padding int) ([]byte, error) {
	return gdes.DecryptECB(cipherText, key, padding)
}

func Md5加密(data interface{}) (encrypt string, err error) {
	return gmd5.EncryptBytes(gconv.Bytes(data))
}

func Md5加密从文件(path string) (encrypt string, err error) {
	return gmd5.EncryptFile(path)
}

func Sha1加密(v interface{}) string {
	return gsha1.Encrypt(v)
}

func Sha1加密从文件(path string) (encrypt string, err error) {
	return gsha1.EncryptFile(path)
}
