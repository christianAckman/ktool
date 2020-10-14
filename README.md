# ktool

## Encryption tool using AWS KMS and AES-256-GCM

#### Installation
`go get github.com/christianAckman/ktool`

#### Usage
```
➜ ktool generate -k myAWSMasterKey

AQIDAHj6KIndQ20/7uMuFPi8qQkQhTIDeWMVc8w6yFUwFUVEXgHPHlG3+Qkq71yhUzOo6KpYAAAAbjBsBgkqhkiG9w0BBwagXzBdAgEAMFgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMvNP5/0nPgFylMKRWAgEQgCtfjdhBQiD9oX5tO6O7wShEU5ZHilmpWFHsxHs07Jrsk79cmBMWxrhP48z3

➜ ktool encrypt -d password123 -k AQIDAHj6KIndQ20/7uMuFPi8qQkQhTIDeWMVc8w6yFUwFUVEXgHPHlG3+Qkq71yhUzOo6KpYAAAAbjBsBgkqhkiG9w0BBwagXzBdAgEAMFgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMvNP5/0nPgFylMKRWAgEQgCtfjdhBQiD9oX5tO6O7wShEU5ZHilmpWFHsxHs07Jrsk79cmBMWxrhP48z3

aUT8R+N1IMzPtzinkXnKNxc+MqvnqIUWl5oUCqpxQNuoFlk6SeC0

➜ ktool decrypt -d aUT8R+N1IMzPtzinkXnKNxc+MqvnqIUWl5oUCqpxQNuoFlk6SeC0 -k AQIDAHj6KIndQ20/7uMuFPi8qQkQhTIDeWMVc8w6yFUwFUVEXgHPHlG3+Qkq71yhUzOo6KpYAAAAbjBsBgkqhkiG9w0BBwagXzBdAgEAMFgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMvNP5/0nPgFylMKRWAgEQgCtfjdhBQiD9oX5tO6O7wShEU5ZHilmpWFHsxHs07Jrsk79cmBMWxrhP48z3

password123
```

#### Menu
```
./ktool 
NAME:
   ktool - AWS KMS functions cli

USAGE:
   ktool [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
   encrypt   function to encrypt data.
   decrypt   function to decrypt data.
   generate  function to generate a data key.
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  print help page (default: false)
   --version   print version (default: false)
```

#### Build
```
go mod init ktool
go build .
```
