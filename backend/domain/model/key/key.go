package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/errors"
)

var key *Key

type Key struct {
	private *rsa.PrivateKey
	public  []byte
}

func init() {
	priv := os.Getenv("PRIVATE_KEY")

	block, _ := pem.Decode([]byte(priv))

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	b, err := x509.MarshalPKIXPublicKey(&private.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	public := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: b,
		},
	)

	key = &Key{
		private: private,
		public:  public,
	}
}

func Decrypt(encrypted []byte) ([]byte, error) {
	decrepted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key.private, encrypted, nil)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	return decrepted, nil
}

func GetPublic() []byte {
	return key.public
}
