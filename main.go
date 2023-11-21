package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadVMImage(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// Function to read the next value from the file
	readNextValue := func() (string, error) {
		if scanner.Scan() {
			// Read the next line
			line := scanner.Text()

			// Convert the hexadecimal string to an integer
			value, err := strconv.ParseUint(line, 16, 32)
			if err != nil {
				return "", err
			}

			binaryRepresentation := fmt.Sprintf("%032b", value)
			return binaryRepresentation, nil

		}
		return "", fmt.Errorf("error reading from file")
	}

	// Read data-size
	dataSizeStr, err := readNextValue()
	if err != nil {
		return nil, fmt.Errorf("error reading data-size: %v", err)
	}
	dataSize, err := strconv.ParseUint(dataSizeStr, 2, 32)
	if err != nil {
		return nil, err
	}

	// Allocate data array
	data := make([]string, dataSize)
	_, err = readNextValue()
	if err != nil {
		return nil, fmt.Errorf("error reading data-size: %v", err)
	}

	//fmt.Println("dataSize:", dataSize)

	// Read data into the array
	for i := 0; i < int(dataSize); i++ {
		// Read each data element
		data[i], err = readNextValue()
		if err != nil {
			return data, nil
		}
	}

	return data, nil
}

func main() {
	filename := "resource/task2.bin" //9
	data, err := loadVMImage(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print loaded data for verification
	//fmt.Println("Data:", data)

	// Execute VM instructions
	executeVM(data)
}

// Global variables
var index = 0

// var commands = "RULDX"
// var commands = "PRUP RULDX"
// var commands = "PRRUP CBEAX"
// var commands = "PRUP J R K PRP JJJJJ R KKKKK PRP JJJJJJJJJ R X"
// var commands = "PRUP M U N PUP MMMMM U NNNNN PUP MMMMMMMMM U X"
//var commands = "PRUP KKKKKKKKKK KKKKKKKKKK NNNNNNNNNN A X"

var commands = generateCommands()

func executeVM(data []string) {
	// Initialize registers
	ip := 0
	sp := len(data)

	// Function to get the next value from the data array
	g := func() string {
		value := data[sp]
		sp++
		return value
	}

	// Function to push a value onto the data array
	f := func(value string) {
		sp--
		data[sp] = value
	}

	// Loop forever
	for {
		// Get the current instruction
		currentInstruction := data[ip]
		ip++

		// Decode the instruction
		binop := currentInstruction[0]
		operation, _ := strconv.ParseInt(currentInstruction[1:8], 2, 32)
		optionalData := currentInstruction[8:]

		// Perform action based on operation
		if binop == '0' {
			// Binop = 0 instructions
			switch operation {
			case 0:
				// "pop"
				sp++
			case 1:
				// "push <const>"
				f(optionalData)
			case 2:
				// "push ip"
				f(fmt.Sprintf("%032b", ip))
			case 3:
				// "push sp"
				f(fmt.Sprintf("%032b", sp))
			case 4:
				// "load"
				addr := g()
				addrInt, err := strconv.ParseInt(addr, 2, 32)
				if err != nil {
					// handle the error
					panic(err)
				}

				f(data[int(addrInt)])
			case 5:
				// "store"
				stData := g()
				addr := g()
				addrInt, err := strconv.ParseInt(addr, 2, 32)
				if err != nil {
					panic(err)
				}

				data[addrInt] = stData
			case 6:
				// "jmp"
				cond := g()
				addr := g()
				if binaryToDecimal(cond) != 0 {
					ip = binaryToDecimal(addr)
				}
			case 7:
				// "not"
				if binaryToDecimal(g()) == 0 {
					f("00000000000000000000000000000001")
				} else {
					f("00000000000000000000000000000000")
				}
			case 8:
				// "putc"
				output := binaryToDecimal(g()) & 0xff
				fmt.Printf("%c", output)
			case 9:
				// "getc"
				input, err := getBinaryInput()
				if err != nil {
					panic(err)
				}
				f(fmt.Sprintf("%032b", binaryToDecimal(input)&0xff))
			case 10:
				// halt
				return
			}
		} else {
			// Binop = 1 instructions
			b := g()
			a := g()

			// Perform the operation based on the value of 'operation'
			switch operation {
			case 0:
				// add
				result := binaryToDecimal(a) + binaryToDecimal(b)
				f(fmt.Sprintf("%032b", result))
			case 1:
				// sub
				result := binaryToDecimal(a) - binaryToDecimal(b)
				f(fmt.Sprintf("%032b", result))
			case 2:
				// mul
				result := binaryToDecimal(a) * binaryToDecimal(b)
				f(fmt.Sprintf("%032b", result))
			case 3:
				// div
				result := binaryToDecimal(a) / binaryToDecimal(b)
				f(fmt.Sprintf("%032b", result))
			case 4:
				// and
				result := binaryToDecimal(a) & binaryToDecimal(b)
				f(fmt.Sprintf("%032b", result))
			case 5:
				// or
				result := binaryToDecimal(a) | binaryToDecimal(b)
				f(fmt.Sprintf("%032b", result))
			case 6:
				// xor
				result := binaryToDecimal(a) ^ binaryToDecimal(b)
				f(fmt.Sprintf("%032b", result))
			case 7:
				// eq
				if binaryToDecimal(a) == binaryToDecimal(b) {
					f("00000000000000000000000000000001")
				} else {
					f("00000000000000000000000000000000")
				}
			case 8:
				// lt
				if binaryToDecimal(a) < binaryToDecimal(b) {
					f("00000000000000000000000000000001")
				} else {
					f("00000000000000000000000000000000")
				}
			}
		}
	}
}

// Function to convert binary string to decimal
func binaryToDecimal(binaryStr string) int {
	decimal, _ := strconv.ParseInt(binaryStr, 2, 32)
	return int(decimal)
}

// Function to get a single byte from stdin and return its binary representation
func getBinaryInput() (string, error) {
	char := commands[index]
	index++

	// Cast the byte to 32 bits and return its binary representation
	return fmt.Sprintf("%032b", char), nil
}

//func getBinaryInput() (string, error) {
//	// Read exactly one byte from stdin
//	var input [1]byte
//	_, err := os.Stdin.Read(input[:])
//	if err != nil {
//		return "", err
//	}
//
//	char := commands[index]
//	index++
//
//	// Cast the byte to 32 bits and return its binary representation
//	return fmt.Sprintf("%032b", char), nil
//}
