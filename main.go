package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"FileEncryptor/filecrypt"

	"github.com/urfave/cli/v2"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	app := &cli.App{
		Name:  "File Encryptor",
		Usage: "A tiny tool that helps encrypt files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "encrypt",
				Aliases:  []string{"e"},
				Usage:    "Encrypt a file",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "decrypt",
				Aliases:  []string{"d"},
				Usage:    "Decrypt a file",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("decrypt") != "" {
				decryptHandle()
			}
			if c.String("encrypt") != "" {
				encryptHandle()
			}
			return nil
		},
	}

	// function := os.Args[1]

	// switch function {
	// case "help", "h":
	// 	printHelp()
	// case "encrypt", "e":
	// 	encryptHandle()
	// case "decrypt", "d":
	// 	decryptHandle()
	// default:
	// 	fmt.Println("Run CryptoGo encrypt to encrypt a file, and CryptoGo decrypt to decrypt a file.")
	// 	os.Exit(1)
	// }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printHelp() {
	fmt.Println("CryptoGo")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tCryptoGo encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

func encryptHandle() {

	if len(os.Args) < 3 {
		println("Missing the path to the file. For more information run CryptoGo help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\nFile successfully protected")

}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password2, _ := terminal.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match. Please try again.\n")
		return getPassword()
	}
	return password
}

func decryptHandle() {

	if len(os.Args) < 3 {
		println("Missing the path to the file. For more information run CryptoGo help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)

	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\nFile successfully decrypted.")

}

func validatePassword(password1 []byte, password2 []byte) bool {
	return bytes.Equal(password1, password2)
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

/*
subtitle main if i didn't use Cli
func main() {

	// If not enough args, return help text
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run CryptoGo encrypt to encrypt a file, and CryptoGo decrypt to decrypt a file.")
		os.Exit(1)
	}

}

func printHelp() {
	fmt.Println("CryptoGo")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tCryptoGo encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}
*/