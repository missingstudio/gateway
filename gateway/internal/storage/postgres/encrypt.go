package postgres

import "github.com/gtank/cryptopasta"

type Crypto struct {
	key *[32]byte
}

func NewCrypto(key *[32]byte) *Crypto {
	return &Crypto{
		key: key,
	}
}

// Encrypt encrypts a string
func (e *Crypto) Encrypt(plaintext []byte) ([]byte, error) {
	return cryptopasta.Encrypt(plaintext, e.key)
}

// Decrypt decrypts a string
func (e *Crypto) Decrypt(ciphertext []byte) ([]byte, error) {
	return cryptopasta.Decrypt(ciphertext, e.key)
}

type CryptoNoop struct{}

func (e *CryptoNoop) Encrypt(plaintext []byte) ([]byte, error) {
	return plaintext, nil
}

func (e *CryptoNoop) Decrypt(ciphertext []byte) ([]byte, error) {
	return ciphertext, nil
}
