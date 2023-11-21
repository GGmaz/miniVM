package main

var firstObject = "RULURULU"
var secondObject = "DRURDRUR"
var thirdObject = "RDRURDRU"
var fourthObject = "DLDRDLDR"
var thirdObjectInv = "DLULDLUL"
var secondObjectInv = "LDLULDLU"

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
		}

	}

	commands += "X"
	println(len(commands))
	return commands
}
