package sections

// Franta is a testing fuction returning an int.
func Franta() int {
	return 33
}

// Task0MaxSum3 receives an array of ints and finds the largest sum of a solid sequence.
func Task0MaxSum3(recordsX []int) (maximum int) {
	k := 0
	for _, num := range recordsX {
		k = _max(0, k) + num
		maximum = _max(maximum, k)
	}
	return maximum
}

// Task1MaxSum3WithIndex receives an array of ints and finds the largest sum of a solid sequence
//  while it returns the sum and index
func Task1MaxSum3WithIndex(recordsX []int) (maximum, left, right int) {
	k := 0
	proposedLeftIndex := left
	for index, numAtIndex := range recordsX {
		if k < 0 {
			proposedLeftIndex = index
		}
		k = _max(0, k) + numAtIndex

		proposedMaximum := _max(maximum, k)
		if proposedMaximum > maximum {
			left = proposedLeftIndex
			right = index
		}
		maximum = proposedMaximum
	}
	return maximum, left, right
}

// Task2CharacterRepeated searches for longest not repeatable sequence of characters in input string
func Task2CharacterRepeated(data string) string {
	runes := []rune(data)
	var temporaryRunes []rune
	var left, right int
	for index, character := range runes {
		if _inList(character, temporaryRunes) {
			// fmt.Printf(">>>> on clear: %v at index %v\n", string(character), index)
			temporaryRunes = []rune{character}
			left = index
			right = index
		} else {
			// fmt.Printf(">>>> adding: %v at index %v\n", string(character), index)
			temporaryRunes = append(temporaryRunes, character)
			right = index
		}
	}
	// fmt.Printf(">>>> left: %v, right: %v\n", left, right)
	return string(runes[left : right+1])
}

func _max(number1, number2 int) int {
	if number1 < number2 {
		return number2
	}
	return number1
}

func _inList(stringValue rune, listing []rune) bool {
	for _, v := range listing {
		if v == stringValue {
			return true
		}
	}
	return false
}
