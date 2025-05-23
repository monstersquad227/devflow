package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"devflow/config"
	"encoding/base64"
	"errors"
	"io"
)

func EncryptAESGCM(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(config.GlobalConfig.Aes.Secret))
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机的 12 字节 nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密并附加认证标签
	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)

	// Base64 编码返回
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptAESGCM(encryptedText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(config.GlobalConfig.Aes.Secret))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 解析 nonce
	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("invalid ciphertext")
	}
	nonce, cipherText := data[:nonceSize], data[nonceSize:]

	// 解密数据
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
