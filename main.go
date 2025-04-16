/*
	why bufio?

It reads/writes data in chunks, minimizing the number of system calls, therefore efficient especially
when working with large files
allows to read data without loading the entire file into memory
To process or manipulate the file's data,

	you need to read its content into memory.

	What Does It Mean to Read content Into Memory?

	To "read content into memory" means taking data stored on a
	persistent storage device (like a hard drive or SSD) and copying it into the computer's memory (RAM) so the program can use it.
	Once the data is in memory, your program can process it quickly.

	what is os.Open?

	os.Open doesn't load the
	file's data; it creates a connection (file handle) to the file on disk

then The bufio.NewScanner reads small chunks from the file (e.g., a single line or word).
It processes one chunk at a time, keeping only the current chunk in memory.
Memory Efficiency:
After processing a chunk, the buffer is reused for the next chunk, ensuring minimal memory use.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputArg := os.Args[1]
	outputArg := os.Args[2]

	// read input file
	fmt.Println("Reading input file...")
	content, err := os.ReadFile(inputArg)
	if err != nil {
		fmt.Printf("Error reading arg %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Processing text...")
	replecedContent := fixText(string(content))

	// write to output file
	fmt.Println("Writing to output file...")
	err = os.WriteFile(outputArg, []byte(replecedContent), 0o644)
	if err != nil {
		fmt.Printf("Error writing arg %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}

func fixText(input string) string {
	input = processString(input)
	input = fixSpace(input)
	lines := strings.Split(input, "\n")
	var processedLines []string
	for _, line := range lines {
		processedLine := line
		processedLine = hexAndBinToDecimal(processedLine)
		processedLine = textModifyCase(processedLine)
		processedLine = fixPunctuations(processedLine)
		processedLine = fixSingleQuotes(processedLine)
		processedLine = fixDoubleQuotes(processedLine)
		processedLine = fixAtoAn(processedLine)
		processedLines = append(processedLines, processedLine)
	}

	return strings.Join(processedLines, "\n")
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("give me 3 args")
// 		return
// 	}

// 	inputArg := os.Args[1]
// 	outputArg := os.Args[2]

// 	// read input file
// 	content, err := os.ReadFile(inputArg)
// 	if err != nil {
// 		fmt.Printf("Error reading arg %v\n", err)
// 		os.Exit(1)
// 	}

// 	replecedcontent := fixText(string(content))

// 	// write to output file
// 	err = os.WriteFile(outputArg, []byte(replecedcontent), 0o644)
// 	if err != nil {
// 		fmt.Printf("Error writing arg %v", err)
// 		os.Exit(1)
// 	}
// }

// func fixText(input string) string {
// 	content := strings.Split(input, "\n")
// 	var processedcontent []string
// 	for _, line := range content {
// 		processedLine := line
// 		processedLine = hexAndBinToDecimal(processedLine)
// 		processedLine = textModifyCase(processedLine)
// 		processedLine = fixPunctuations(processedLine)
// 		processedLine = fixAtoAn(processedLine)
// 		processedcontent = append(processedcontent, processedLine)
// 	}

// 	return strings.Join(processedcontent, "\n")
// }
