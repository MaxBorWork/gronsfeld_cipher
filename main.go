package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

type Letter struct {
	index int
	shift int
	letter string
}

var alphabet = []string{"а","б","в","г","д","е","ё","ж","з","и","к","л","м","н","о","п","р","с","т","у","ф","х","ц","ч",
						"ш","щ","ъ","ы","ь","э","ю","я"}

func main()  {
	text := reaInputString("Enter text: ")
	password := reaInputString("Enter password: ")
	passArr, err := decodePass(password)
	if err != nil {
		panic(err)
	}
	encodedString := encodeString(text, passArr)
	decodeString(encodedString, passArr)
}

func reaInputString(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	fmt.Println(text)
	return text
}

func decodePass(password string) ([]int, error) {
	passArr := make([]int, len(password))
	stringArr := strings.Split(password, "")
	for i, v := range stringArr {
		var err error
		passArr[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("Password must contain only numbers")
		}
	}
	return passArr, nil
}

func encodeString(text string, password []int) string {
	var encodedString string
	pass_index := 0;
	textArr := strings.Split(text, "")
	for _, val := range textArr {
		isFound := false
		for alphabetIndex, alphabetValue := range alphabet {
			if alphabetValue == val {
				shiftIndex := alphabetIndex + password[pass_index]
				encodedString = encodedString + alphabet[shiftIndex]
				isFound = true
			}
		}
		if !isFound {
			encodedString = encodedString + val
		}
		pass_index++

		if (pass_index) == len(password) {
			pass_index = 0;
		}
	}
	fmt.Println(encodedString)
	return encodedString
}

func decodeString(encodedText string, password []int) {
	var decodedString string
	pass_index := 0;
	textArr := strings.Split(encodedText, "")
	for _, val := range textArr {
		isFound := false
		for alphabetIndex, alphabetValue := range alphabet {
			if alphabetValue == val {
				shiftIndex := alphabetIndex - password[pass_index]
				decodedString = decodedString + alphabet[shiftIndex]
				isFound = true
			}
		}
		if !isFound {
			decodedString = decodedString + val
		}
		pass_index++

		if (pass_index) == len(password) {
			pass_index = 0;
		}
	}
	fmt.Println(decodedString)
}
