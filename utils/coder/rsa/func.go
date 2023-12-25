package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func GenerateKeyPair2File(bits int, publicFile, privateFile string) error {
	// 1、生成私钥文件的核心步骤：
	// 1) 生成密钥对
	privateKer, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return errors.New("generate key error")
	}
	// 2) 将私钥对象转换成DER编码形式
	derPrivateKer := x509.MarshalPKCS1PrivateKey(privateKer)
	// 3) 创建私钥pem文件
	file, err := os.Create(privateFile)
	if err != nil {
		return errors.New("create private file error")
	}
	// 4) 对密钥信息进行编码，写入到私钥文件中
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derPrivateKer,
	}
	err = pem.Encode(file, block)
	if err != nil {
		return errors.New("store private key failed ")
	}

	// 2、生成公钥文件的核心步骤：
	// 1) 生成公钥对象
	publicKey := &privateKer.PublicKey
	// 2) 将公钥对象序列化为DER编码格式
	derPublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	// 3) 创建公钥pem文件
	file, err = os.Create(publicFile)
	if err != nil {
		return errors.New("create public file error")
	}
	// 4) 对公钥信息进行编码，写入到公钥文件中
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublicKey,
	}
	err = pem.Encode(file, block)
	if err != nil {
		return errors.New("store public key failed ")
	}
	return nil
}

func GenerateKeyPair(bits int) (publicKeyStr, privateKeyStr string, err error) {
	// 1、生成私钥文件的核心步骤：
	// 1) 生成密钥对
	privateKer, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return publicKeyStr, privateKeyStr, err
	}
	// 2) 将私钥对象转换成DER编码形式
	derPrivateKer := x509.MarshalPKCS1PrivateKey(privateKer)

	// 3) 对密钥信息进行编码，写入到私钥文件中
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derPrivateKer,
	}
	privateKeyStr = string(pem.EncodeToMemory(block))

	// 2、生成公钥文件的核心步骤：
	// 1) 生成公钥对象
	publicKey := &privateKer.PublicKey
	// 2) 将公钥对象序列化为DER编码格式
	derPublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return publicKeyStr, privateKeyStr, err
	}
	// 3) 对公钥信息进行编码，生成字符串
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublicKey,
	}
	publicKeyStr = string(pem.EncodeToMemory(block))

	return publicKeyStr, privateKeyStr, err
}

// ParsePublicKey 通过字符串，解析出公钥对象
func ParsePublicKey(keyString string) (*rsa.PublicKey, error) {
	// 1、获取公钥字节
	publicKeyBytes := []byte(keyString)

	// 2、解码公钥字节，生成加密块对象
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	// 3、解析DER编码的公钥，生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 4、公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

// ReadParsePublicKey 读取公钥文件，解析出公钥对象
func ReadParsePublicKey(filename string) (*rsa.PublicKey, error) {
	// 1、读取公钥文件，获取公钥字节
	publicKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、解码公钥字节，生成加密块对象
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	// 3、解析DER编码的公钥，生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 4、公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

// ParsePrivaterKey 从字符串解析出私钥对象
func ParsePrivaterKey(keyString string) (*rsa.PrivateKey, error) {
	// 1、获取私钥字节
	privateKeyBytes := []byte(keyString)
	// 2、对私钥文件进行编码，生成加密块对象
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// ReadParsePrivaterKey 读取私钥文件，字符串解析出私钥对象
func ReadParsePrivaterKey(filename string) (*rsa.PrivateKey, error) {
	// 1、读取私钥文件，获取私钥字节
	privateKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、对私钥文件进行编码，生成加密块对象
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// Encrypt 加密字节数组，返回字节数组
func Encrypt(originalBytes []byte, filename string) ([]byte, error) {
	// 1、读取公钥文件，解析出公钥对象
	publicKey, err := ReadParsePublicKey(filename)
	if err != nil {
		return nil, err
	}
	// 2、加密，参数是随机数、公钥对象、需要加密的字节
	// PKCS#1 v1.5 padding
	// Reader是一个全局共享的密码安全的强大的伪随机生成器
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, originalBytes)
}

// Decrypt 解密字节数组，返回字节数组
func Decrypt(cipherBytes []byte, filename string) ([]byte, error) {
	// 1、读取私钥文件，解析出私钥对象
	privateKey, err := ReadParsePrivaterKey(filename)
	if err != nil {
		return nil, err
	}
	// 2、ras解密，参数是随机数、私钥对象、需要解密的字节
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherBytes)
}

// EncryptStringToBase64 加密字符串，返回base64处理的字符串
func EncryptStringToBase64(originalText, filename string) (string, error) {
	cipherBytes, err := Encrypt([]byte(originalText), filename)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// DecryptStringFromBase64  解密经过base64处理的加密字符串，返回加密前的明文
func DecryptStringFromBase64(cipherlText, filename string) (string, error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherlText)
	originalBytes, err := Decrypt(cipherBytes, filename)
	if err != nil {
		return "", err
	}
	return string(originalBytes), nil
}
