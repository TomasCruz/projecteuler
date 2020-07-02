package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 59; XOR decryption

Each character on a computer is assigned a unique code and the preferred standard is ASCII
(American Standard Code for Information Interchange). For example, uppercase A = 65, asterisk (*) = 42,
and lowercase k = 107.

A modern encryption method is to take a text file, convert the bytes to ASCII, then XOR each byte
with a given value, taken from a secret key. The advantage with the XOR function is that using
the same encryption key on the cipher text, restores the plain text;
for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.

For unbreakable encryption, the key is the same length as the plain text message, and the key is made up
of random bytes. The user would keep the encrypted message and the encryption key in different locations,
and without both "halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key.
If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the
message. The balance for this method is using a sufficiently long password key for security, but short enough
to be memorable.

Your task has been made easy, as the encryption key consists of three lower case characters.
Using p059_cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes,
and the knowledge that the plain text must contain common English words,
decrypt the message and find the sum of the ASCII values in the original text.
*/

func main() {
	var fileName string

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		fileName = "p059_cipher.txt"
	}

	projecteuler.Timed(calc, fileName)
}

func calc(args ...interface{}) (result string, err error) {
	fileName := args[0].(string)

	var textNumbers []string
	if textNumbers, err = projecteuler.FileToStrings(fileName); err != nil {
		fmt.Println(err)
		return
	}

	var builder strings.Builder
	inputStrings := strings.Split(textNumbers[0], ",")

	for _, s := range inputStrings {
		b, _ := strconv.ParseInt(s, 10, 16)
		builder.WriteRune(rune(b))
	}
	inputString := builder.String()

	var one, two, three rune
	for one = 'a'; one <= 'z'; one++ {
		for two = 'a'; two <= 'z'; two++ {
			for three = 'a'; three <= 'z'; three++ {
				key := []rune{one, two, three}

				txt, candidate := crypt(inputString, key)
				if candidate {
					res64 := sumASCII(txt)
					result = strconv.FormatInt(res64, 10)
					return
				}
			}
		}
	}

	result = "-1" // wanted combination not found
	return
}

func crypt(inputString string, key []rune) (txt string, candidate bool) {
	j := 0
	var builder strings.Builder
	for _, b := range inputString {
		builder.WriteRune(b ^ key[j])
		j++
		if j == 3 {
			j = 0
		}
	}
	txt = builder.String()

	ind1 := wordOccurencies(txt, "the")
	ind2 := wordOccurencies(txt, "of")
	ind3 := wordOccurencies(txt, "and")
	if ind1 > 1 && ind2 > 1 && ind3 > 1 {
		candidate = true
	}

	return
}

func sumASCII(txt string) (sum int64) {
	sum = 0
	for _, b := range txt {
		sum += int64(b)
	}

	return
}

func wordOccurencies(txt, word string) (occ int) {
	for {
		i := strings.Index(txt, word)
		if i == -1 {
			break
		}

		occ++
		txt = txt[i+1:]
	}

	return
}
