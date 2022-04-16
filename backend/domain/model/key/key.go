package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
)

var (
	privateKey *rsa.PrivateKey
	PublicKey  []byte
)

func init() {
	// 自前でpemキーを用意する場合
	// k, err := readKeyFromFile()
	// if err != nil {
	// 	log.Fatalf("failed to read pem key file: %v", err)
	// }
	// key, err := exportPEMStrToPrivKey(k)

	// firebaseが発行する秘密鍵から公開鍵を作る場合
	priv := os.Getenv("PRIVATE_KEY")

	key, err := exportPEMStrToPrivKey([]byte(priv))

	privateKey = key

	if err != nil {
		log.Fatalf("failed to export pem key: %v", err)
	}

	b, err := x509.MarshalPKIXPublicKey(&key.PublicKey)

	if err != nil {
		log.Fatalf("failed to marshal public key: %v", err)
	}

	PublicKey = pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: b,
		},
	)

}

func Decrypt(encrypted []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encrypted, nil)
}

// func readKeyFromFile() ([]byte, error) {
// 	return ioutil.ReadFile("key.pem")
// }

func exportPEMStrToPrivKey(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)

	// 自前でpemキーを用意する場合
	// key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	pk, ok := key.(*rsa.PrivateKey)

	if !ok {
		return nil, errors.New("private key is invalid")
	}

	return pk, nil
}
