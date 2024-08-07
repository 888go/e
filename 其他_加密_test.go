// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*"

package e

import (
	"encoding/hex"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/test/gtest"

	"github.com/gogf/gf/v2/crypto/gdes"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"os"
	"testing"
)

var (
	content          = []byte("pibigstar")
	content_16, _    = gbase64.DecodeString("v1jqsGHId/H8onlVHR8Vaw==")
	content_24, _    = gbase64.DecodeString("0TXOaj5KMoLhNWmJ3lxY1A==")
	content_32, _    = gbase64.DecodeString("qM/Waw1kkWhrwzek24rCSA==")
	content_16_iv, _ = gbase64.DecodeString("DqQUXiHgW/XFb6Qs98+hrA==")
	content_32_iv, _ = gbase64.DecodeString("ZuLgAOii+lrD5KJoQ7yQ8Q==")
	// iv 长度必须等于blockSize，只能为16
	iv         = []byte("Hello My GoFrame")
	key_16     = []byte("1234567891234567")
	key_17     = []byte("12345678912345670")
	key_24     = []byte("123456789123456789123456")
	key_32     = []byte("12345678912345678912345678912345")
	keys       = []byte("12345678912345678912345678912346")
	key_err    = []byte("1234")
	key_32_err = []byte("1234567891234567891234567891234 ")

	// cfb模式blockSize补位长度, add by zseeker
	padding_size      = 16 - len(content)
	content_16_cfb, _ = gbase64.DecodeString("oSmget3aBDT1nJnBp8u6kA==")
)

func TestEncrypt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data, err := Aes加密(content, key_16)
		gtest.Assert(err, nil)
		gtest.Assert(data, []byte(content_16))
		data, err = Aes加密(content, key_24)
		gtest.Assert(err, nil)
		gtest.Assert(data, []byte(content_24))
		data, err = Aes加密(content, key_32)
		gtest.Assert(err, nil)
		gtest.Assert(data, []byte(content_32))
		data, err = Aes加密(content, key_16, iv)
		gtest.Assert(err, nil)
		gtest.Assert(data, []byte(content_16_iv))
		data, err = Aes加密(content, key_32, iv)
		gtest.Assert(err, nil)
		gtest.Assert(data, []byte(content_32_iv))
	})
}

func TestDecrypt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		decrypt, err := Aes解密([]byte(content_16), key_16)
		gtest.Assert(err, nil)
		gtest.Assert(decrypt, content)

		decrypt, err = Aes解密([]byte(content_24), key_24)
		gtest.Assert(err, nil)
		gtest.Assert(decrypt, content)

		decrypt, err = Aes解密([]byte(content_32), key_32)
		gtest.Assert(err, nil)
		gtest.Assert(decrypt, content)

		decrypt, err = Aes解密([]byte(content_16_iv), key_16, iv)
		gtest.Assert(err, nil)
		gtest.Assert(decrypt, content)

		decrypt, err = Aes解密([]byte(content_32_iv), key_32, iv)
		gtest.Assert(err, nil)
		gtest.Assert(decrypt, content)

		decrypt, err = Aes解密([]byte(content_32_iv), keys, iv)
		gtest.Assert(err, "invalid padding")
	})
}

func TestEncryptErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// encrypt key error
		_, err := Aes加密(content, key_err)
		gtest.AssertNE(err, nil)
	})
}

func TestDecryptErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// decrypt key error
		encrypt, err := Aes加密(content, key_16)
		_, err = Aes解密(encrypt, key_err)
		gtest.AssertNE(err, nil)

		// decrypt content too short error
		_, err = Aes解密([]byte("test"), key_16)
		gtest.AssertNE(err, nil)

		// decrypt content size error
		_, err = Aes解密(key_17, key_16)
		gtest.AssertNE(err, nil)
	})
}

func TestPKCS5UnPaddingErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// PKCS5UnPadding blockSize zero
		_, err := gaes.PKCS5UnPadding(content, 0)
		gtest.AssertNE(err, nil)

		// PKCS5UnPadding src len zero
		_, err = gaes.PKCS5UnPadding([]byte(""), 16)
		gtest.AssertNE(err, nil)

		// PKCS5UnPadding src len > blockSize
		_, err = gaes.PKCS5UnPadding(key_17, 16)
		gtest.AssertNE(err, nil)

		// PKCS5UnPadding src len > blockSize
		_, err = gaes.PKCS5UnPadding(key_32_err, 32)
		gtest.AssertNE(err, nil)
	})
}

func TestEncryptCFB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var padding int = 0
		data, err := AesCFB加密(content, key_16, &padding, iv)
		gtest.Assert(err, nil)
		gtest.Assert(padding, padding_size)
		gtest.Assert(data, []byte(content_16_cfb))
	})
}

func TestDecryptCFB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		decrypt, err := AesCFB解密([]byte(content_16_cfb), key_16, padding_size, iv)
		gtest.Assert(err, nil)
		gtest.Assert(decrypt, content)
	})
}

func TestCrc32(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "pibigstar"
		result := 693191136
		encrypt1 := Crc32加密(s)
		encrypt2 := Crc32加密([]byte(s))
		t.AssertEQ(int(encrypt1), result)
		t.AssertEQ(int(encrypt2), result)

		strmd5, _ := gmd5.Encrypt(s)
		test1 := Crc32加密(strmd5)
		test2 := Crc32加密([]byte(strmd5))
		t.AssertEQ(test2, test1)
	})
}

var (
	errKey     = []byte("1111111111111234123456789")
	errIv      = []byte("123456789")
	errPadding = 5
)

func TestDesECB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.NOPADDING
		result := "858b176da8b12503"
		// encrypt test
		cipherText, err := DesECB加密(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := DesECB解密(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")

		// encrypt err test. when throw exception,the err is not equal nil and the string is nil
		errEncrypt, err := DesECB加密(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		errEncrypt, err = DesECB加密(text, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// err decrypt test.
		errDecrypt, err := DesECB解密(cipherText, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		errDecrypt, err = DesECB解密(cipherText, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		errPadding := 5
		result := "858b176da8b12503ad6a88b4fa37833d"
		cipherText, err := DesECB加密(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := DesECB解密(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")

		// err test
		errEncrypt, err := DesECB加密(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		errDecrypt, err := DesECB解密(cipherText, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})
}

func Test3DesECB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("1111111111111234")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		result := "a23ee24b98c26263a23ee24b98c26263"
		// encrypt test
		cipherText, err := gdes.EncryptECBTriple(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECBTriple(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "1234567812345678")
		// err test
		errEncrypt, err := DesECB加密(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		key := []byte("111111111111123412345678")
		text := []byte("123456789")
		padding := gdes.PKCS5PADDING
		errPadding := 5
		result := "37989b1effc07a6d00ff89a7d052e79f"
		// encrypt test
		cipherText, err := gdes.EncryptECBTriple(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECBTriple(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "123456789")
		// err test, when key is err, but text and padding is right
		errEncrypt, err := gdes.EncryptECBTriple(text, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// when padding is err,but key and text is right
		errEncrypt, err = gdes.EncryptECBTriple(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// decrypt err test,when key is err
		errEncrypt, err = gdes.DecryptECBTriple(text, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
	})
}

func TestDesCBC(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		iv := []byte("12345678")
		result := "40826a5800608c87585ca7c9efabee47"
		// encrypt test
		cipherText, err := gdes.EncryptCBC(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBC(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test.
		errEncrypt, err := gdes.EncryptCBC(text, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.EncryptCBC(text, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = gdes.EncryptCBC(text, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// decrypt err test. the key is err
		errDecrypt, err := gdes.DecryptCBC(cipherText, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.DecryptCBC(cipherText, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = gdes.DecryptCBC(cipherText, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := gdes.EncryptCBC(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBC(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")
		// err test
		errEncrypt, err := gdes.EncryptCBC(text, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
	})
}

func Test3DesCBC(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := []byte("1111111112345678")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		iv := []byte("12345678")
		result := "bfde1394e265d5f738d5cab170c77c88"
		// encrypt test
		cipherText, err := gdes.EncryptCBCTriple(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBCTriple(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test
		errEncrypt, err := gdes.EncryptCBCTriple(text, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.EncryptCBCTriple(text, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = gdes.EncryptCBCTriple(text, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// decrypt err test
		errDecrypt, err := gdes.DecryptCBCTriple(cipherText, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.DecryptCBCTriple(cipherText, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = gdes.DecryptCBCTriple(cipherText, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		key := []byte("111111111234567812345678")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := gdes.EncryptCBCTriple(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBCTriple(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")
	})

}

var (
	s = "pibigstar"
	// online generated MD5 value
	result = "d175a1ff66aedde64344785f7f7a3df8"
)

type user2 struct {
	name     string
	password string
	age      int
}

func TestEncryptmd5(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		encryptString, _ := Md5加密(s)
		gtest.Assert(encryptString, result)

		result := "1427562bb29f88a1161590b76398ab72"
		encrypt, _ := Md5加密(123456)
		gtest.AssertEQ(encrypt, result)
	})

	gtest.C(t, func(t *gtest.T) {
		user2 := &user2{
			name:     "派大星",
			password: "123456",
			age:      23,
		}
		result := "70917ebce8bd2f78c736cda63870fb39"
		encrypt, _ := Md5加密(user2)
		gtest.AssertEQ(encrypt, result)
	})
}

func TestEncryptString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		encryptString, _ := gmd5.EncryptString(s)
		gtest.Assert(encryptString, result)
	})
}

func TestEncryptFilemd5(t *testing.T) {
	path := "test.text"
	errorPath := "err.txt"
	result := "e6e6e1cd41895beebff16d5452dfce12"
	gtest.C(t, func(t *gtest.T) {
		file, err := os.Create(path)
		defer os.Remove(path)
		defer file.Close()
		gtest.Assert(err, nil)
		_, _ = file.Write([]byte("Hello Go Frame"))
		encryptFile, _ := Md5加密从文件(path)
		gtest.AssertEQ(encryptFile, result)
		// when the file is not exist,encrypt will return empty string
		errEncrypt, _ := Md5加密从文件(errorPath)
		gtest.AssertEQ(errEncrypt, "")
	})

}

type user struct {
	name     string
	password string
	age      int
}

func TestEncryptsha1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		user := &user{
			name:     "派大星",
			password: "123456",
			age:      23,
		}
		result := "97386736e3ee4adee5ca595c78c12129f6032cad"
		encrypt := Sha1加密(user)
		gtest.AssertEQ(encrypt, result)
	})
	gtest.C(t, func(t *gtest.T) {
		result := "5b4c1c2a08ca85ddd031ef8627414f4cb2620b41"
		s := Sha1加密("pibigstar")
		gtest.AssertEQ(s, result)
	})
}

func TestEncryptFile(t *testing.T) {
	path := "test.text"
	errPath := "err.text"
	gtest.C(t, func(t *gtest.T) {
		result := "8b05d3ba24b8d2374b8f5149d9f3fbada14ea984"
		file, err := os.Create(path)
		defer os.Remove(path)
		defer file.Close()
		gtest.Assert(err, nil)
		_, _ = file.Write([]byte("Hello Go Frame"))
		encryptFile, _ := Sha1加密从文件(path)
		gtest.AssertEQ(encryptFile, result)
		// when the file is not exist,encrypt will return empty string
		errEncrypt, _ := Sha1加密从文件(errPath)
		gtest.AssertEQ(errEncrypt, "")
	})
}
