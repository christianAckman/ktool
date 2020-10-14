package cmd

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
)

// func describeKey(kms *kms.KMS) {}
// func getAliases(kms *kms.KMS) {}

// GenerateDataKey generates a data key using a KMS master key
func GenerateDataKey(kmsSvc *kms.KMS, keyName string) (string, error) {

	req := &kms.GenerateDataKeyWithoutPlaintextInput{
		KeyId:   aws.String("alias/" + keyName),
		KeySpec: aws.String("AES_256"),
	}

	resp, err := kmsSvc.GenerateDataKeyWithoutPlaintext(req)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(resp.CiphertextBlob), nil
}

// DecryptDataKey - Decrypts a data key using a KMS master key
func DecryptDataKey(kmsSvc *kms.KMS, dataKey string) (decryptedDataKey []byte, err error) {

	bytes, err := base64.StdEncoding.DecodeString(dataKey)

	if err != nil {
		return nil, err
	}

	req := &kms.DecryptInput{
		CiphertextBlob: []byte(bytes),
	}

	resp, err := kmsSvc.Decrypt(req)
	if err != nil {
		return nil, err
	}

	return resp.Plaintext, err
}

// DecryptKMS decrypts data using AWS KMS
func DecryptKMS(kmsSvc *kms.KMS, dataKey string, keyName string, encryptedVal string) (string, error) {

	bytes, err := base64.StdEncoding.DecodeString(encryptedVal)

	if err != nil {
		return "", err
	}

	req := &kms.DecryptInput{
		KeyId:          aws.String("alias/" + keyName),
		CiphertextBlob: []byte(bytes),
	}

	resp, err := kmsSvc.Decrypt(req)
	if err != nil {
		return "", err
	}

	return string(resp.Plaintext), err
}

// EncryptKMS encrypts data using AWS KMS
func EncryptKMS(kmsSvc *kms.KMS, dataKey string, keyName string, unencryptedVal string) (string, error) {

	req := &kms.EncryptInput{
		KeyId:     aws.String("alias/" + keyName),
		Plaintext: []byte(unencryptedVal),
	}

	resp, err := kmsSvc.Encrypt(req)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(resp.CiphertextBlob), err
}
