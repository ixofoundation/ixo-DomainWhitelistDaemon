package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type WhitelistDomain struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url";unique`
	Hash string `json:"Hash"`
}

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("whitelist.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&WhitelistDomain{})

	return nil
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
	var newDomain = WhitelistDomain{Name: name, Url: url}

	db, err := gorm.Open(sqlite.Open("whitelist.db"), &gorm.Config{})
	if err != nil {
		return newDomain, err
	}
	db.Create(&WhitelistDomain{Name: name, Url: url})

	return newDomain, nil
}
