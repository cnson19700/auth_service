package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetText() string {
	reader := bufio.NewReader(os.Stdin)
	stri, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(stri, "\n")
}

func GetFloatNumber() float64 {
	str := GetText()
	num, err := strconv.ParseFloat(str, 64)

	if err != nil {
		log.Fatal(err)
	}
	return num
}

func GetIntNumber() int {
	str := GetText()
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
