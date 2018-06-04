package tools

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// GenPemPrivateKey GenPemPrivateKey
func GenPemPrivateKey(bits int) (string, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	var _pem string
	buf := bytes.NewBufferString(_pem)
	err = pem.Encode(buf, block)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
