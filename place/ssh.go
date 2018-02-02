package place

import (
	"log"

	"../utils/sshkey"
	"../utils/verify"
)

const sshKeySize = 2048

// load an ssh key or create a new one
func loadOrCreateSSHKey(path string, logger *log.Logger) error {
	if verify.File(path) != nil {
		logger.Printf("SSH Key at %s does not exist, generating a new one\n", path)
		if err := sshkey.CreateKeyPair(path, sshKeySize); err != nil {
			return err
		}
		logger.Printf("Wrote private key to %s\n", path)
		logger.Printf("Wrote public key to %s\n", path+".pub")
	}

	pk, err := sshkey.PublicString(path)
	if err != nil {
		return err
	}
	logger.Printf("Found SSH Public Key:\n%s\n", pk)

	return nil
}
