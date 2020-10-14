package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	cmd "ktool/cmd"

	"github.com/urfave/cli/v2"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

var (
	key     string
	dataKey string
	data    string
)

func main() {

	cryptFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "dataKey",
			Aliases:     []string{"k"},
			Usage:       "data key",
			Destination: &dataKey,
		},
		&cli.StringFlag{
			Name:        "data",
			Aliases:     []string{"d"},
			Usage:       "data to encrypt/decrypt",
			Destination: &data,
		},
	}

	generateFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "key",
			Aliases:     []string{"k"},
			Usage:       "key name (alias) or ARN",
			Destination: &key,
		},
	}

	cli.HelpFlag = &cli.BoolFlag{Name: "help", Usage: "print help page", Aliases: []string{"h"}}
	cli.VersionFlag = &cli.BoolFlag{Name: "version", Usage: "print version"}

	app := &cli.App{
		Name:    "ktool",
		Usage:   "AWS KMS functions cli",
		Version: "v0.0.1",
		Commands: []*cli.Command{
			{
				Name:  "encrypt",
				Usage: "function to encrypt data.",
				Flags: cryptFlags,
				Action: func(c *cli.Context) error {
					if dataKey == "" || data == "" {
						cli.ShowSubcommandHelp(c)
						return cli.Exit("Must supply data key and data to be encrypted.", 1)
					}

					kmsSvc, err := initialize()
					if err != nil {
						fmt.Println("Error getting AWS Session: " + err.Error())
						return cli.Exit("err", 1)
					}

					decryptedDataKey, err := cmd.DecryptDataKey(kmsSvc, dataKey)
					if err != nil {
						fmt.Println("err: " + err.Error())
						return cli.Exit("err", 1)
					}

					encryptedData, err := cmd.EncryptGCM([]byte(data), decryptedDataKey)

					if err != nil {
						fmt.Println("Error encrypting data: " + err.Error())
						return cli.Exit("err", 1)
					}

					fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(encryptedData))
					return nil
				},
			},
			{
				Name:  "decrypt",
				Usage: "function to decrypt data.",
				Flags: cryptFlags,
				Action: func(c *cli.Context) error {
					if dataKey == "" || data == "" {
						cli.ShowSubcommandHelp(c)
						return cli.Exit("Must supply data key and data to be decrypted.", 1)
					}

					kmsSvc, err := initialize()
					if err != nil {
						fmt.Println("Error getting AWS Session: " + err.Error())
						return cli.Exit("err", 1)
					}

					decryptedDataKey, err := cmd.DecryptDataKey(kmsSvc, dataKey)
					if err != nil {
						fmt.Println("err: " + err.Error())
						return cli.Exit("err", 1)
					}

					cipherText, err := base64.StdEncoding.DecodeString(data)
					if err != nil {
						fmt.Println("Error decoding data: " + err.Error())
						return cli.Exit("err", 1)
					}
					decryptedData, err := cmd.DecryptGCM(cipherText, decryptedDataKey)

					if err != nil {
						fmt.Println("Error decrypting data: " + err.Error())
						return cli.Exit("err", 1)
					}

					fmt.Printf("%s\n", decryptedData)
					return nil
				},
			},
			{
				Name:  "generate",
				Usage: "function to generate a data key.",
				Flags: generateFlags,
				Action: func(c *cli.Context) error {
					if key == "" {
						cli.ShowSubcommandHelp(c)
						return cli.Exit("Must supply AWS KMS key.", 1)
					}

					kmsSvc, err := initialize()
					if err != nil {
						fmt.Println("Error getting AWS Session: " + err.Error())
						return cli.Exit("err", 1)
					}

					dataKey, err := cmd.GenerateDataKey(kmsSvc, key)
					if err != nil {
						fmt.Println("Error generating data key: " + err.Error())
						return cli.Exit("err", 1)
					}

					fmt.Println(dataKey)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initialize() (*kms.KMS, error) {
	// session using ~/.aws/credentials and ~/.aws/config
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	kmsSvc := kms.New(sess)
	return kmsSvc, nil
}
