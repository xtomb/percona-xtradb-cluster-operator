package encryption

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
)

const (
	version    = "Keyring file version:2.0"
	keyType    = "AES"
	eof        = "EOF"
	userID     = ""
	keyLen     = 32
	paddingLen = 5
)

func NewKeyring() ([]byte, error) {
	key, err := key()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new key: %v", err)
	}

	sha := sha256.Sum256(key)

	keyring := bytes.NewBuffer(make([]byte, 0))
	_, err = keyring.WriteString(version)
	if err != nil {
		return nil, fmt.Errorf("failed to write version: %v", err)
	}

	_, err = keyring.Write(key)
	if err != nil {
		return nil, fmt.Errorf("failed to write key: %v", err)
	}

	_, err = keyring.WriteString(eof)
	if err != nil {
		return nil, fmt.Errorf("failed to write eof: %v", err)
	}

	_, err = keyring.Write(sha)
	if err != nil {
		return nil, fmt.Errorf("failed to write SHA sum : %v", err)
	}

	return keyring.Bytes(), nil
}

func key() ([]byte, error) {
	keyID := keyID()
	buf := make([]byte, 0)
	key := bytes.NewBuffer(buf)

	err := binary.Write(key, binary.LittleEndian, int64(podSize(len(keyID))))
	if err != nil {
		return nil, fmt.Errorf("failed to write key pod size: %v", err)
	}

	err = binary.Write(key, binary.LittleEndian, int64(len(keyID)))
	if err != nil {
		return nil, fmt.Errorf("failed to write length of key id: %v", err)
	}

	err = binary.Write(key, binary.LittleEndian, int64(len(keyType)))
	if err != nil {
		return nil, fmt.Errorf("failed to write length of key type: %v", err)
	}

	err = binary.Write(key, binary.LittleEndian, int64(len(userID)))
	if err != nil {
		return nil, fmt.Errorf("failed to write length of user id: %v", err)
	}

	err = binary.Write(key, binary.LittleEndian, int64(keyLen))
	if err != nil {
		return nil, fmt.Errorf("failed to write length of AES key: %v", err)
	}

	_, err = key.WriteString(keyID)
	if err != nil {
		return nil, fmt.Errorf("failed to write key id: %v", err)
	}

	_, err = key.WriteString(keyType)
	if err != nil {
		return nil, fmt.Errorf("failed to write key type: %v", err)
	}

	_, err = key.WriteString(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to write user id: %v", err)
	}

	aes, err := aesKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate AES key: %v", err)
	}

	_, err = key.Write(aes)
	if err != nil {
		return nil, fmt.Errorf("failed to write AES key: %v", err)
	}

	_, err = key.Write(make([]byte, paddingLen))
	if err != nil {
		return nil, fmt.Errorf("failed to write padding: %v", err)
	}

	return key.Bytes(), nil
}

func aesKey() ([]byte, error) {
	aes := make([]byte, keyLen)
	_, err := rand.Read(aes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random sequence of bytes: %v", err)
	}

	obfuscator := []byte("*305=Ljt0*!@$Hnm(*-9-w;:")
	i := 0
	l := 0
	for i < len(aes) {
		aes[i] ^= obfuscator[l]
		i++
		l = (l + 1) % len(obfuscator)
	}

	return aes, nil
}

func podSize(keyIDLen int) int {
	size := 4*8 + keyIDLen + len(keyType) + len(userID) + 8 + keyLen
	padding := (8 - (size % 8)) % 8
	return size + padding
}

func keyID() string {
	return fmt.Sprintf("INNODBKey-%s-1", uuid.New())
}
