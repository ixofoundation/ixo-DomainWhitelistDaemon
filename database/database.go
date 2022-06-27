package database

import (
	"crypto"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type WhitelistDomain struct {
	gorm.Model
	Name      string `json:"name"`
	Url       string `json:"url";unique`
	Signature string `json:"hash"`
}

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("whitelist.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&WhitelistDomain{})

	return nil
}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

func Sign(msg, secret []byte) <-chan string {
	r := make(chan string)

	go func() {
		priv, pub := GenerateRsaKeyPair()

		// Export the keys to pem string
		priv_pem := ExportRsaPrivateKeyAsPemStr(priv)
		pub_pem, _ := ExportRsaPublicKeyAsPemStr(pub)

		// Import the keys from pem string
		// priv_parsed, _ := ParseRsaPrivateKeyFromPemStr(priv_pem)
		// pub_parsed, _ := ParseRsaPublicKeyFromPemStr(pub_pem)

		// Before signing, we need to hash our message
		// The hash is what we actually sign
		msgHash := sha256.New()
		_, err := msgHash.Write(msg)
		if err != nil {
			panic(err)
		}
		msgHashSum := msgHash.Sum(nil)

		// In order to generate the signature, we provide a random number generator,
		// our private key, the hashing algorithm that we used, and the hash sum
		// of our message

		signature, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, msgHashSum, nil)
		if err != nil {
			panic(err)

		}
		fmt.Println(string(priv_pem))
		fmt.Println(string(pub_pem))
		priv_data := []byte(priv_pem)
		pub_data := []byte(pub_pem)
		ioerr := ioutil.WriteFile("private.txt", priv_data, 777)

		if ioerr != nil {
			log.Fatal(err)
		}

		ioerr2 := ioutil.WriteFile("public.txt", pub_data, 777)

		if ioerr2 != nil {
			log.Fatal(err)
		}

		fmt.Println(hex.EncodeToString(signature))
		r <- hex.EncodeToString(signature)

	}()

	return r
}

func GetAllWhitelisteDomains() ([]WhitelistDomain, error) {
	var domains []WhitelistDomain

	db, err := gorm.Open(sqlite.Open("whitelist.db"), &gorm.Config{})
	if err != nil {
		return domains, err
	}

	db.Find(&domains)

	return domains, nil
}

func CreateWhitelistedDomain(name string, url string) (WhitelistDomain, error) {

	msg := []byte(url)
	key := []byte(os.Getenv("SECRETKEY"))

	hash := <-Sign(msg, key)
	fmt.Println("HASH:", hash)

	var newDomain = WhitelistDomain{Name: name, Url: url, Signature: hash}

	db, err := gorm.Open(sqlite.Open("whitelist.db"), &gorm.Config{})
	if err != nil {
		return newDomain, err
	}
	db.Create(&WhitelistDomain{Name: name, Signature: hash, Url: url})

	return newDomain, nil
}
