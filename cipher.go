package interfaces

// Cipherable is an interface which allows to encypt and decrypt data
type Cipherable interface {
	Encryptable
	Decryptable
}

// Encryptable is an interface which is used to encrypt data
type Encryptable interface {
	Encrypt(text []byte, key []byte) (value []byte, err error)
}

// Decryptable is an interface which is used to decrypt data
type Decryptable interface {
	Decrypt(text []byte, key []byte) (value []byte, err error)
}
