package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'foo' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING_ARRAY codeList
 *  2. STRING_ARRAY shoppingCart
 */

func hasOrder(shoppingCart []string, index int, codeMap []string) bool {
	for i := 0; i < len(codeMap); i++ {
		if index < len(shoppingCart) && codeMap[i] == "anything" || shoppingCart[index] == codeMap[i] {
			index++
		} else {
			return false
		}
	}

	return true
}

func foo(codeList []string, shoppingCart []string) int32 {
	var cartIndex, codeIndex int
	codesMap := make(map[int][]string)

	for k, v := range codeList {
		codesMap[k] = strings.Split(v, " ")
	}

	for {
		if cartIndex < len(shoppingCart) && codeIndex < len(codesMap) {
			current := shoppingCart[cartIndex]

			if codesMap[codeIndex][0] == "anything" || codesMap[codeIndex][0] == current && hasOrder(shoppingCart, cartIndex, codesMap[codeIndex]) {
				cartIndex += len(codesMap[codeIndex])
				codeIndex++
			} else {
				cartIndex++
			}
		} else {
			break
		}
	}

	if codeIndex == len(codesMap) {
		return 1
	}

	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	codeListCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var codeList []string

	for i := 0; i < int(codeListCount); i++ {
		codeListItem := readLine(reader)
		codeList = append(codeList, codeListItem)
	}

	shoppingCartCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var shoppingCart []string

	for i := 0; i < int(shoppingCartCount); i++ {
		shoppingCartItem := readLine(reader)
		shoppingCart = append(shoppingCart, shoppingCartItem)
	}

	result := foo(codeList, shoppingCart)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
