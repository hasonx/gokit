package cryptogram

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// AESEncrypt 加密
func AESEncrypt(plaintext, key string) (string, error) {
	plainData := []byte(plaintext)
	aesKey := []byte(key)
	block, err := aes.NewCipher(aesKey) // 生成加密用的block
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	padding := blockSize - len(plainData)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plainData = append(plainData, padText...)
	// 采用KEY MD5做初始向量
	md5Val := []byte(fmt.Sprintf("%x", md5.Sum(aesKey)))
	iv := md5Val[0:block.BlockSize()]
	// CBC
	mode := cipher.NewCBCEncrypter(block, iv)
	cipherData := make([]byte, len(plainData))
	mode.CryptBlocks(cipherData, plainData)
	return hex.EncodeToString(cipherData), nil
}

// AESDecrypt 解密
func AESDecrypt(ciphertext, key string) (string, error) {
	aesKey := []byte(key)
	cipherData, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}
	if len(cipherData) < aes.BlockSize ||
		len(cipherData)%block.BlockSize() != 0 {
		return "", fmt.Errorf("密文格式不合法")
	}
	// 采用KEY MD5做初始向量
	md5Val := []byte(fmt.Sprintf("%x", md5.Sum(aesKey)))
	iv := md5Val[0:block.BlockSize()]
	mode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	mode.CryptBlocks(plainData, cipherData)
	if len(plainData) == 0 {
		return "", fmt.Errorf("解密出错")
	}
	// 反补
	length := len(plainData)
	unpadding := int(plainData[length-1])
	if unpadding > length {
		return "", fmt.Errorf("AES解密反补位时发现数组操作越界,可能是加解密key不一致导致,plainData length:%d,unpadding:%d", length, unpadding)
	}
	plainData = plainData[:(length - unpadding)]
	return string(plainData), nil
}
