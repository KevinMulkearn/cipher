package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cFlag := flag.String("c", "", "Type of encryption ciper to use")
	fFlag := flag.String("f", "", "Input file to encrypt")
	mFlag := flag.String("m", "", "Message string to encrypt")
	oFlag := flag.String("o", "", "Output file with encrypted data")
	flag.Parse()

	// Check if we have a cipher type selected and something to encrypt
	if *cFlag == "" || (*fFlag == "" && *mFlag == "") {
		printUsage()
		os.Exit(0)
	}

	if *fFlag != "" {
		handleFileInput(*fFlag, *oFlag)
	} else if *mFlag != "" {
		handleMessageInput(*mFlag)
	}

	// TODO - Add ability to change cipher type

}
func handleFileInput(filename string, outputfile string) {
	// Read from file
	fileIn, err := os.Open(filename)
	defer fileIn.Close()
	if err != nil {
		fmt.Println("Error opening file:", filename)
		fmt.Println(err)
		os.Exit(1)
	}

	var lines []string

	scanner := bufio.NewScanner(fileIn)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Write to file
	if outputfile == "" {
		outputfile = "encrypted_" + filename
	}

	fileOut, err := os.Create(outputfile)
	defer fileOut.Close()
	if err != nil {
		fmt.Println("Error creating file:", outputfile)
		fmt.Println(err)
		os.Exit(1)
	}

	writer := bufio.NewWriter(fileOut)
	for _, line := range lines {
		fmt.Fprintln(writer, encrypt(line))
	}
	writer.Flush()
}

func handleMessageInput(message string) {
	fmt.Println("Input:", message)
	fmt.Println("Output:", encrypt(message))
}

func whichCipher(cipher string, char rune) rune {
	if cipher == "caeser" {
		return caeser(char, 3)
	}
	return -1
}

// encrypt encrypts a string of text
func encrypt(text string) string {
	output := []rune(text)
	for i, character := range text {
		output[i] = caeser(character, 3)
	}
	return string(output)
}

// caeser returns the new character for a caeser shift
func caeser(char rune, shift int) rune {
	if char < 'a' || char > 'z' {
		return char
	}

	new_char := char + rune(shift)
	if new_char > 'z' {
		return new_char - 26
	}
	return new_char
}

func printUsage() {
	exe_name := filepath.Base(os.Args[0])
	fmt.Println(exe_name, "[OPTION]... [FILE|MESSAGE]...")
	fmt.Println()
	fmt.Println("-c,		cipher encryption type")
	fmt.Println("-f,		input file to encrypt")
	fmt.Println("-m,		input message to encrypt")
	fmt.Println("-o,		output file with encrypted data")
}
