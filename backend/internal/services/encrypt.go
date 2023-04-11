package services

import (
	"comp4913-backend/pkg/reEncrypt/curve"
	"comp4913-backend/pkg/reEncrypt/recrypt"
	"crypto/ecdsa"
	"fmt"
	"gorm.io/gorm"
	"math/big"
)

type EncryptService struct {
	db *gorm.DB
}

func NewEncryptService(db *gorm.DB) *EncryptService {
	return &EncryptService{db}
}

func (s EncryptService) GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	pri, pub, _ := curve.GenerateKeys()
	return pri, pub
}

func (s EncryptService) EncryptFile(inputFile, outPutFile string, pubKey *ecdsa.PublicKey) recrypt.Capsule {
	cap, err := recrypt.EncryptFile(inputFile, outPutFile, pubKey)
	if err != nil {
		fmt.Println(err)
	}
	return *cap
}

func (s EncryptService) ReEncryptFile(rk *big.Int, capsule *recrypt.Capsule) recrypt.Capsule {
	cap, err := recrypt.ReEncryption(rk, capsule)
	if err != nil {
		fmt.Println("ReEncryption Error:", err)
	}
	return *cap
}

func (s EncryptService) DecryptFile(inputFile, outPutFile string, bPriKey *ecdsa.PrivateKey, capsule *recrypt.Capsule, pubX *ecdsa.PublicKey) error {
	err := recrypt.DecryptFile(inputFile, outPutFile, bPriKey, capsule, pubX)
	if err != nil {
		fmt.Println("Decrypt Error:", err)
	}
	return err
}
