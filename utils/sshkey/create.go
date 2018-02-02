package sshkey

// MakeSSHKeyPair make a pair of public and private keys for SSH access.
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

// CreateKeyPair is used to create an ssh public / private key pair.
// key is encoded in the format for inclusion in an OpenSSH authorized_keys file.
// Private Key generated is PEM encoded.
// adapted from https://stackoverflow.com/a/34347463
func CreateKeyPair(path string, size int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return err
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(path)
	defer privateKeyFile.Close()
	if err != nil {
		return err
	}
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return err
	}

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path+".pub", ssh.MarshalAuthorizedKey(pub), 0655)
}
