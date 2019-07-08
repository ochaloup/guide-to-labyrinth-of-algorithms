package sections

// Franta is a testing fuction returning an int.
func Franta() int {
	return 33
}

// MaxSum3 receives an array of ints and finds the largest sum of a solid sequence.
func MaxSum3(recordsX []int) (maximum int) {
	k := 0
	for _, num := range recordsX {
		k = _max(0, k) + num
		maximum = _max(maximum, k)
	}
	return maximum
}

// MaxSum3WithIndex receives an array of ints and finds the largest sum of a solid sequence
//  while it returns the sum and index
func MaxSum3WithIndex(recordsX []int) (maximum, left, right int) {
	k := 0
	for index, numAtIndex := range recordsX {
		if k < 0 {
			left = index
		}
		k = _max(0, k) + numAtIndex

		proposedMaximum := _max(maximum, k)
		if proposedMaximum > maximum {
			right = index
		}
		maximum = proposedMaximum
	}
	return maximum, left, right
}

func _max(number1, number2 int) int {
	if number1 < number2 {
		return number2
	}
	return number1
}
