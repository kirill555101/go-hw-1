package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"./uniq"
)

func ReadText(inputFile *os.File) (text []string, err error) {
	stringReader := bufio.NewReader(inputFile)
	for {
		str, err := stringReader.ReadString('\n')
		if err != nil && len(str) == 0 {
			break
		} else if err != nil {
			return nil, errors.New("File reading error")
		}
		text = append(text, str[:len(str)-1])
	}

	return text, nil
}

func InitFlags() (*uniq.Options, error) {
	flagCPtr := flag.Bool("c", false, "For number of occurrences of lines in the input_file")
	flagDPtr := flag.Bool("d", false, "Print only those lines that were repeated in the input_file data")
	flagUPtr := flag.Bool("u", false, "Print only those lines that have not been repeated in the input_file data")

	flagFPtr := flag.Int("f", 0, "Ignore the first num_fields fields in the line")
	flagSPtr := flag.Int("s", 0, "Ignore the first num_chars characters in the string")
	flagIPtr := flag.Bool("i", false, "Case insensitive")

	flag.Parse()

	options := uniq.NewOptions(
		*flagCPtr,
		*flagDPtr,
		*flagUPtr,
		*flagFPtr,
		*flagSPtr,
		*flagIPtr,
	)

	// If not (one parameter used or all parametres unused)
	if !options.CorrectParams() {
		return options, errors.New("Invalid arguments passed")
	}

	options.ResetParams()

	return options, nil
}

func PrintUsage() {
	fmt.Println("Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
}

func main() {
	inputFile := os.Stdin
	outputFile := os.Stdout
	var err error

	options, err := InitFlags()
	if err != nil {
		PrintUsage()
		return
	}

	switch len(flag.Args()) {
	case 0:

	case 1:
		inputFile, err = os.Open(flag.Args()[len(flag.Args())-1])
		if err != nil {
			inputFile = os.Stdin
		}
		defer inputFile.Close()

	case 2:
		inputFile, err = os.Open(flag.Args()[len(flag.Args())-2])
		if err != nil {
			inputFile = os.Stdin
		}
		defer inputFile.Close()

		outputFile, err = os.Create(flag.Args()[len(flag.Args())-1])
		if err != nil {
			outputFile = os.Stdout
		}
		defer outputFile.Close()

	default:
		PrintUsage()
	}

	text, err := ReadText(inputFile)
	if err != nil || len(text) == 0 {
		fmt.Println(err)
		return
	}

	for _, str := range uniq.Uniq(options, text) {
		_, err = outputFile.WriteString(str + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
