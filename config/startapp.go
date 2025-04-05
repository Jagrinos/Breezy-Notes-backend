package config

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

func SetupKeys() (err error) {
	CONNSTR, err = setupConnstr()
	if err != nil {
		return err
	}

	PRIVATEKEY, err = loadPrivateKeyFromFile()
	if err != nil {
		return err
	}

	PUBLICKEY, err = loadPublicKeyFromFile()
	if err != nil {
		return err
	}

	return nil
}

func setupConnstr() (string, error) {
	return fmt.Sprintf(
			"postgres://%s:%s@uasbreezydb:5432/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		),
		nil
}

func loadPrivateKeyFromFile() (*ecdsa.PrivateKey, error) {
	data, err := os.ReadFile(filepath.Join(".", "config", "private.pem"))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("can't decode PEM")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadPublicKeyFromFile() (*ecdsa.PublicKey, error) {
	data, err := os.ReadFile(filepath.Join(".", "config", "public.pem"))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("can't decode PEM")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not ECDSA public key")
	}

	return ecdsaPubKey, nil
}
