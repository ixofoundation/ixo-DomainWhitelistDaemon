package database

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type WhitelistDomain struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url";unique`
	Hash string `json:"hash"`
}

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("whitelist.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&WhitelistDomain{})

	return nil
}

func Sign(msg, key []byte) <-chan string {
	r := make(chan string)

	go func() {
		mac := hmac.New(sha256.New, key)
		mac.Write(msg)
		var hexvalue = hex.EncodeToString(mac.Sum(nil))
		r <- hexvalue
	}()

	return r
}

func Verify(msg, key []byte, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil)), nil
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

	var newDomain = WhitelistDomain{Name: name, Url: url, Hash: hash}

	db, err := gorm.Open(sqlite.Open("whitelist.db"), &gorm.Config{})
	if err != nil {
		return newDomain, err
	}
	db.Create(&WhitelistDomain{Name: name, Hash: hash, Url: url})

	return newDomain, nil
}
