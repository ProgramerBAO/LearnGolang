//package main
//
//import (
//	"crypto"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/x509"
//	"encoding/pem"
//	"os"
//)
//
//func main() {
//
//	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
//	if err != nil {
//		panic(err)
//	}
//	// The public key is a part of the *rsa.PrivateKey struct
//	publicKey := privateKey.PublicKey
//	//保存私钥
//	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
//	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
//	//使用pem格式对x509输出的内容进行编码
//	//创建文件保存私钥
//	privateFile, err := os.Create("private.pem")
//	if err != nil {
//		panic(err)
//	}
//	defer privateFile.Close()
//
//	//private := "-----BEGIN PRIVATE KEY-----\nMIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEA012afkerKKeYY0bY\nytMKWSQ6J5M39Tn/mKhYojKGYZdvqaDU58FLxTIl99iTmqZHGQBYlah1Sm1erJ/p\n7N/VVQIDAQABAkBSKDYfE7iB6vMe07D81Z9WGvDH5T+BAHRLYFRjZ4Q4djExksZN\nhQ4ktHD9D8ITJpQz9+hYkhSEDMXip2GZffgZAiEA/DchskzCpjvO5pnzEadesAjh\nyg4T7HoiNHuww41ToaMCIQDWiYzneLE6nOC1RaZTRX9ucAKduDcfselDezooEqdM\npwIgY5EgLFNBRH633zFHU+DO8I+RE0Mbem98sVtjHM/eBCECICMKNElaHRhFpy30\nQkY3g6i2Ardf7yDuHfs3lTgWU9zhAiB4qKXBclDEAUpG1YIg0sPSJoyGhoUf/bmG\n6h0K89wlvw==\n-----END PRIVATE KEY-----\n"
//	//hash := "hash"
//	//res := Signed(private, hash)
//	//fmt.Println(res)
//}
//
//// Signed 签名操作
//func Signed(PrivateKey string, hashValue string) (SignH string) {
//	// func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte)
//	buf := []byte(PrivateKey)
//	block, _ := pem.Decode(buf)
//	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
//	signH, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, []byte(hashValue))
//	return string(signH[:])
//}
//
//// Signed2 Signed 签名操作
//func Signed2(PrivateKey string, hashValue string) (SignH string) {
//	// func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte)
//	buf := []byte(PrivateKey)
//	block, _ := pem.Decode(buf)
//	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
//	signH, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, []byte(hashValue))
//	return string(signH[:])
//}

package main

import (
	"crypto"

	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`  
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9
kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5
LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQAB
AoGAbSNg7wHomORm0dWDzvEpwTqjl8nh2tZyksyf1I+PC6BEH8613k04UfPYFUg1
0F2rUaOfr7s6q+BwxaqPtz+NPUotMjeVrEmmYM4rrYkrnd0lRiAxmkQUBlLrCBiF
u+bluDkHXF7+TUfJm4AZAvbtR2wO5DUAOZ244FfJueYyZHECQQD+V5/WrgKkBlYy
XhioQBXff7TLCrmMlUziJcQ295kIn8n1GaKzunJkhreoMbiRe0hpIIgPYb9E57tT
/mP/MoYtAkEA4Ti6XiOXgxzV5gcB+fhJyb8PJCVkgP2wg0OQp2DKPp+5xsmRuUXv
720oExv92jv6X65x631VGjDmfJNb99wq5QJBAMSHUKrBqqizfMdOjh7z5fLc6wY5
M0a91rqoFAWlLErNrXAGbwIRf3LN5fvA76z6ZelViczY6sKDjOxKFVqL38ECQG0S
pxdOT2M9BM45GJjxyPJ+qBuOTGU391Mq1pRpCKlZe4QtPHioyTGAAMd4Z/FX2MKb
3in48c0UX5t3VjPsmY0CQQCc1jmEoB83JmTHYByvDpc8kzsD8+GmiPVrausrjj4p
y2DQpGmUic2zqCxl6qXMpBGtFEhrUbKhOiVOJbRNGvWW
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`  
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfw1/P15GQzGGYvNwVmXIGGxea
8Pb2wJcF7ZW7tmFdLSjOItn9kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQol
DOkEzNP0B8XKm+Lxy4giwwR5LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TU
GQD6QzKY1Y8xS+FoQQIDAQAB
-----END PUBLIC KEY-----    
`)

var private = "-----BEGIN RSA PRIVATE KEY-----\nMIICXQIBAAKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9\nkvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5\nLJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQAB\nAoGAbSNg7wHomORm0dWDzvEpwTqjl8nh2tZyksyf1I+PC6BEH8613k04UfPYFUg1\n0F2rUaOfr7s6q+BwxaqPtz+NPUotMjeVrEmmYM4rrYkrnd0lRiAxmkQUBlLrCBiF\nu+bluDkHXF7+TUfJm4AZAvbtR2wO5DUAOZ244FfJueYyZHECQQD+V5/WrgKkBlYy\nXhioQBXff7TLCrmMlUziJcQ295kIn8n1GaKzunJkhreoMbiRe0hpIIgPYb9E57tT\n/mP/MoYtAkEA4Ti6XiOXgxzV5gcB+fhJyb8PJCVkgP2wg0OQp2DKPp+5xsmRuUXv\n720oExv92jv6X65x631VGjDmfJNb99wq5QJBAMSHUKrBqqizfMdOjh7z5fLc6wY5\nM0a91rqoFAWlLErNrXAGbwIRf3LN5fvA76z6ZelViczY6sKDjOxKFVqL38ECQG0S\npxdOT2M9BM45GJjxyPJ+qBuOTGU391Mq1pRpCKlZe4QtPHioyTGAAMd4Z/FX2MKb\n3in48c0UX5t3VjPsmY0CQQCc1jmEoB83JmTHYByvDpc8kzsD8+GmiPVrausrjj4p\ny2DQpGmUic2zqCxl6qXMpBGtFEhrUbKhOiVOJbRNGvWW\n-----END RSA PRIVATE KEY-----"

//var private = "-----BEGIN PRIVATE KEY-----\nMIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEA012afkerKKeYY0bY\nytMKWSQ6J5M39Tn/mKhYojKGYZdvqaDU58FLxTIl99iTmqZHGQBYlah1Sm1erJ/p\n7N/VVQIDAQABAkBSKDYfE7iB6vMe07D81Z9WGvDH5T+BAHRLYFRjZ4Q4djExksZN\nhQ4ktHD9D8ITJpQz9+hYkhSEDMXip2GZffgZAiEA/DchskzCpjvO5pnzEadesAjh\nyg4T7HoiNHuww41ToaMCIQDWiYzneLE6nOC1RaZTRX9ucAKduDcfselDezooEqdM\npwIgY5EgLFNBRH633zFHU+DO8I+RE0Mbem98sVtjHM/eBCECICMKNElaHRhFpy30\nQkY3g6i2Ardf7yDuHfs3lTgWU9zhAiB4qKXBclDEAUpG1YIg0sPSJoyGhoUf/bmG\n6h0K89wlvw==\n-----END PRIVATE KEY-----\n"
var hash = "hash"

// 生成哈希值 私钥签名在这里完成 然后返回哈希值和签名后的证书
func rsaSignature(plainText []byte, privateKeyFile string) ([]byte, error) {

	block, _ := pem.Decode([]byte(private))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	// 计算原始内容的散列值
	h := sha256.New()
	h.Write(plainText)
	hValue := h.Sum(nil)

	// 通过rsa.SignPKCS1v15使用私钥对原始内容散列值进行签名
	digestSign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hValue)
	return digestSign, err
}

// 首先对比哈希值 再验证内容是否一样（这个顺序可以改）
func rsaVerifySign(plainText []byte, publicKeyFile string, signed []byte) bool {

	block, _ := pem.Decode(publicKey)
	publicKeyInt, err := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := publicKeyInt.(*rsa.PublicKey)

	// 计算原始内容的散列值
	h := sha256.New()
	h.Write(plainText)
	hValue := h.Sum(nil)

	// 确认签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hValue, signed)

	return err == nil
}

func main() {
	content := []byte("Hello digest sign")
	sign, err := rsaSignature(content, "private.pem")
	if err != nil {
		return
	}
	fmt.Println("signature:", string(sign))
	fmt.Println("verify result:", rsaVerifySign(content, "public.pem", sign))
}
