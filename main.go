package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"ixowhitelistdaemon/database"
	whitelist_domain "ixowhitelistdaemon/whitelist"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/getwhitelist", whitelist_domain.GetAllWhitelisteDomains)
	app.Post("/api/createwhitelistitem", whitelist_domain.CreateWhitelistedDomain)
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

func doesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	app := fiber.New()
	dbErr := database.InitDatabase()

	//Setup server RSA Keys

	filecheck := doesFileExist("private.txt")

	//Check for local key
	if !filecheck {
		fmt.Println("File Not Found")
		priv, pub := GenerateRsaKeyPair()

		// Export the keys to pem string
		priv_pem := ExportRsaPrivateKeyAsPemStr(priv)
		pub_pem, _ := ExportRsaPublicKeyAsPemStr(pub)
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
	}

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
