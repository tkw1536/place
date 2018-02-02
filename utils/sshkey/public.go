package sshkey

// print the public key string from a file
import (
	cryptossh "golang.org/x/crypto/ssh"
)

// PublicString turns the content inside a private key
// into a private string
func PublicString(file string) (string, error) {

	key, err := Get(file)
	if err != nil {
		return "", err
	}

	return string(cryptossh.MarshalAuthorizedKey(key.Signer.PublicKey())), nil
}
