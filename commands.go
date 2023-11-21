package main

var firstObject = "RULURULU"
var secondObject = "DRURDRUR"
var thirdObject = "RDRURDRU"
var fourthObject = "DLDRDLDR"
var thirdObjectInv = "DLULDLUL"
var secondObjectInv = "LDLULDLU"
var sixthObjectInv = "ULDLULDL"
var fifthObjectInv = "LULDLULD"
var fifthObject = "URDRURDR"
var sixthObject = "RURDRURD"

func generateCommands() string {
	var commands string
	for i := 0; i <= 5; i++ {
		if i == 0 {
			commands += firstObject

			commands += "P"
			for j := 0; j < 10; j++ {
				commands += "R"
			}
			commands += "UUP"

			commands += secondObject

			commands += "PRRRRDDP"
			commands += secondObject

			commands += "PRRRRUUP"
			commands += thirdObject

			commands += "P"
			for j := 0; j < 10; j++ {
				commands += "R"
			}
			commands += "DDP"
			commands += fourthObject
		} else if i == 1 {
			commands += "P"
			for j := 0; j < 12; j++ {
				commands += "U"
			}
			commands += "LLLLLP"

			commands += thirdObjectInv

			commands += "P"
			for j := 0; j < 22; j++ {
				commands += "L"
			}
			commands += "P"

			commands += secondObjectInv
		} else if i == 2 {
			commands += "P"
			for j := 0; j < 8; j++ {
				commands += "U"
			}
			commands += "LLP"

			commands += secondObject

			commands += "P"
			for j := 0; j < 26; j++ {
				commands += "R"
			}
			commands += "P"

			commands += thirdObject
		} else if i == 3 {
			commands += "P"
			for j := 0; j < 8; j++ {
				commands += "U"
			}
			commands += "LLP"

			commands += sixthObjectInv

			commands += "P"
			for j := 0; j < 22; j++ {
				commands += "L"
			}
			commands += "P"

			commands += fifthObjectInv
		} else if i == 4 {
			commands += "P"
			for j := 0; j < 8; j++ {
				commands += "U"
			}
			commands += "LLLLLP"

			commands += firstObject

			commands += "P"
			for j := 0; j < 10; j++ {
				commands += "R"
			}
			commands += "DDDDDDP"

			commands += fifthObject

			commands += "PRRRRUUP"
			commands += fifthObject

			commands += "PRRRRDDP"
			commands += sixthObject

			commands += "P"
			for j := 0; j < 10; j++ {
				commands += "R"
			}
			commands += "UUUUUUP"
			commands += fourthObject
		}
	}

	// center figures
	commands += "P"
	for j := 0; j < 7; j++ {
		commands += "D"
	}
	for j := 0; j < 21; j++ {
		commands += "L"
	}
	commands += "PJJJJJMMMMM"

	// first figure
	commands += secondObject
	commands += "KKKKKNNNNN"

	// second figure
	commands += "PDRP"
	commands += secondObjectInv

	// third figure
	for j := 0; j < 10; j++ {
		commands += "KN"
	}
	commands += "PDLP"

	commands += secondObject

	// fourth figure
	commands += "PRP"
	for j := 0; j < 10; j++ {
		commands += "KN"
	}
	commands += "PDP"

	commands += secondObjectInv

	// fifth figure
	commands += "JJJJJJJJJJPLDDPKKKKKKKKKK"
	for j := 0; j < 8; j++ {
		commands += "KN"
	}

	commands += secondObject

	commands += "X"
	return commands
}
