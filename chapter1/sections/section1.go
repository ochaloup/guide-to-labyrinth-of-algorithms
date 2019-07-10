package sections

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

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
/*
 * Upravte algoritmus MaxSoučet3, aby oznámil nejen maximální součet, ale také
 * polohu příslušného úseku.
 */
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
/*
 * Na vstupu je text složený z písmen české abecedy a mezer. Vymyslete algoritmus,
 * který najde nejdelší úsek textu, v němž se žádné písmeno neopakuje.
 */
func Task2CharacterRepeated(data string) string {
	runes := []rune(data)
	var temporaryRunes []rune
	var left, right int
	for index, character := range runes {
		if _inList(character, temporaryRunes) && string(character) != " " {
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

// Task3ShortestPartWithAllLetters finds the shortest sequence of characters from the input string
//  which contains all letters of alphabet
/*
 * Najděte v českém textu nejkratší úsek, který obsahuje všechna písmena abecedy.
 * Malá a velká písmena nerozlišujte.
 */
func Task3ShortestPartWithAllLetters(data string) (string, int, int, error) {
	// TODO: FIX for letter 'ch'

	// having prepared map with alphabet characters initialized with index -1
	// alphabet := "aábcčdďeéěfghiíjklmnňoópqrřsštťuúůvwxyýzž"
	// for _, character := range strings.Split(alphabet, "") {
	//   firstSeenIndex[character] = -1
	//   lastSeenIndex[character] = -1
	// }

	firstSeenIndex := make(map[string]int)
	lastSeenIndex := make(map[string]int)
	// going through the input string and index all letters
	for index, character := range strings.Split(data, "") {
		val, ok := firstSeenIndex[character]
		if !ok || val == -1 {
			firstSeenIndex[character] = index
		}
		lastSeenIndex[character] = index
	}
	fmt.Println("first seen index:", firstSeenIndex)
	fmt.Println("last seen index:", lastSeenIndex)

	// sorting by value where
	// the first "last seen index" declares left side index
	// the last "first seen index" declares right side index
	sortedByFirstSeen := make([]string, 0, len(firstSeenIndex))
	sortedByLastSeen := make([]string, 0, len(lastSeenIndex))
	for key := range firstSeenIndex {
		sortedByFirstSeen = append(sortedByFirstSeen, key)
		sortedByLastSeen = append(sortedByLastSeen, key)
	}
	sort.Slice(sortedByFirstSeen,
		func(i, j int) bool {
			return firstSeenIndex[sortedByFirstSeen[i]] > firstSeenIndex[sortedByFirstSeen[j]]
		})
	sort.Slice(sortedByLastSeen,
		func(i, j int) bool {
			return lastSeenIndex[sortedByLastSeen[i]] > lastSeenIndex[sortedByLastSeen[j]]
		})
	fmt.Printf("first seen sorted %v\n", sortedByFirstSeen)
	fmt.Printf("last seen sorted %v\n", sortedByLastSeen)

	if len(sortedByFirstSeen) == 0 || len(sortedByLastSeen) == 0 {
		return "", -1, -1, errors.New("input data was wrong")
	}
	leftIndex := lastSeenIndex[sortedByLastSeen[len(sortedByLastSeen)-1]]
	rightIndex := firstSeenIndex[sortedByFirstSeen[0]]
	fmt.Printf("if>> lefti: %v[%v], righti: %v[%v]\n", sortedByFirstSeen[0], leftIndex, sortedByLastSeen[0], rightIndex)
	/*if leftIndex > rightIndex {
		tmp := leftIndex
		leftIndex = rightIndex
		rightIndex = tmp
		fmt.Printf("if>> lefti: %v[%v], righti: %v[%v]\n", sortedByFirstSeen[0], leftIndex, sortedByLastSeen[0], rightIndex)
	} else {
		fmt.Printf("else>> lefti: %v[%v], righti: %v[%v]\n", sortedByLastSeen[0], leftIndex, sortedByFirstSeen[0], rightIndex)
	}*/
	return string([]rune(data)[leftIndex : rightIndex+1]), leftIndex, rightIndex, nil
}

// Task3ShortestPartWithAllLetters2 finds the shortest sequence of characters from the input string
//  which contains all letters of alphabet
/*
 * Najděte v českém textu nejkratší úsek, který obsahuje všechna písmena abecedy.
 * Malá a velká písmena nerozlišujte.
 */
func Task3ShortestPartWithAllLetters2(data string) (string, int, int, error) {
	// TODO: FIX for letter 'ch'
	// TODO: need to remove characters from the other side (from end to the start)
	//       and take the better result

	// having prepared map with alphabet characters initialized with index -1
	// alphabet := "aábcčdďeéěfghiíjklmnňoópqrřsštťuúůvwxyýzž"
	// for _, character := range strings.Split(alphabet, "") {
	//   firstSeenIndex[character] = -1
	//   lastSeenIndex[character] = -1
	// }

	dataSlice := strings.Split(data, "")
	charactersCounting := make(map[string]int)
	// going through the input string and index all letters
	for _, character := range dataSlice {
		val, ok := charactersCounting[character]
		if !ok || val == -1 {
			charactersCounting[character] = 1
		} else {
			charactersCounting[character]++
		}
	}
	// from left
	var indexLeft, indexRight int
	for i := 0; i < len(dataSlice); i++ {
		indexLeft = i
		if charactersCounting[dataSlice[i]] == 1 {
			break
		} else {
			charactersCounting[dataSlice[i]]--
		}
	}
	for i := len(dataSlice) - 1; i >= 0; i-- {
		indexRight = i
		if charactersCounting[dataSlice[i]] == 1 {
			break
		} else {
			charactersCounting[dataSlice[i]]--
		}
	}
	return strings.Join(dataSlice[indexLeft:indexRight+1], ""), indexLeft, indexRight, nil
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
