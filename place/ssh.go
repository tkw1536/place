package place

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	"./config"
	"golang.org/x/crypto/ssh"
)

// load an ssh key or create a new one
func loadOrCreateSSHKey(cfg *config.Config) error {
	if !isFile(cfg.SSHKeyPath) {
		cfg.Logger.Printf("SSH Key at %s does not exist, generating a new one\n", cfg.SSHKeyPath)

		publicKeyPath := cfg.SSHKeyPath + ".pub"
		if err := makeSSHKeyPair(cfg.SSHKeyPath+".pub", cfg.SSHKeyPath); err != nil {
			return err
		}
		cfg.Logger.Printf("Wrote private key to %s\n", cfg.SSHKeyPath)
		cfg.Logger.Printf("Wrote public key to %s\n", publicKeyPath)
	}

	pk, err := getPublicKeyString(cfg.SSHKeyPath)
	if err != nil {
		return err
	}
	cfg.Logger.Printf("Found SSH Public Key:\n%s\n", pk)

	return nil
}

// check if a file exists
func isFile(file string) bool {
	stat, err := os.Stat(file)
	if err != nil {
		return false
	}

	return stat.Mode().IsRegular()
}

// the size for ssh keys to generate
const sshKeySize = 2048

// MakeSSHKeyPair make a pair of public and private keys for SSH access.
// Public key is encoded in the format for inclusion in an OpenSSH authorized_keys file.
// Private Key generated is PEM encoded
// adapted from https://stackoverflow.com/a/34347463
func makeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, sshKeySize)
	if err != nil {
		return err
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(privateKeyPath)
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
	return ioutil.WriteFile(pubKeyPath, ssh.MarshalAuthorizedKey(pub), 0655)
}

// print the public key string from a file
func getPublicKeyString(file string) (string, error) {

	// load the file
	pemBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	// read the private key
	sgn, err := ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(""))
	if err != nil {
		return "", err
	}

	return string(ssh.MarshalAuthorizedKey(sgn.PublicKey())), nil
}
