package internal

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	salt       = "static-salt"
	iterations = 100000
	keyLength  = 32
)

func GenerateKeyFromKeywords(keywords []string) ([]byte, error) {
	keywordString := strings.Join(keywords, "")
	keywordBytes := []byte(keywordString)
	key := pbkdf2.Key(keywordBytes, []byte(salt), iterations, keyLength, sha256.New)
	return key, nil
}

func GetUserKeywords() []string {
	var keywords []string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter 8 keywords, one at a time:")

	for i := 1; i <= 8; i++ {
		fmt.Printf("Enter keyword %d: ", i)
		keyword, _ := reader.ReadString('\n')

		keywords = append(keywords, keyword[:len(keyword)-1])
	}

	return keywords
}
