// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

// What is the total of all the name scores in the file?

// Just felt like writting my own sorting algorithm for the words read from file
// Also just felt like writting it using a linked list for each letter
// which, I know, is not the best way to sort things
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/urfave/cli"
)

const (
	asciiStartUppercase int32 = 64
	asciiEndUppercase   int32 = 91
)

func namesScores() {
	app := cli.NewApp()
	app.Name = "Names Scores"
	app.Usage = "Names Scores app will return the total of all the name scores in the file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Usage: "file path of names.txt",
		},
	}
	app.Action = func(c *cli.Context) error {
		err := validateContext(c)
		if err != nil {
			return err
		}
		_, err = readAndSort(c.String("path"))
		if err != nil {
			return err
		}
		return nil
	}
	app.Run(os.Args)
}

func validateContext(c *cli.Context) error {
	if c.String("path") == "" {
		return fmt.Errorf("path is empty and should be a valid unix path to the names.txt file")
	}
	return nil
}

//readAndSort reads the file and add the names in an alphabetical order to
//a data structure
func readAndSort(path string) (map[string]LinkedList, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening file at path %v. Got %v", path, err)
	}
	sortedNames, _ := readNames(f)
	// printNames(sortedNames)
	// verifMySort(sortedNames, sortedWithGo)
	// fmt.Println(computeTruth(sortedWithGo))

	fmt.Println(computeScore(sortedNames))

	return nil, nil
}

//computeScore computes the score of the file
func computeScore(sortedNames map[string]*LinkedList) int32 {
	var score int32
	var count int32 = 1
	for i := asciiStartUppercase + 1; i < asciiEndUppercase; i++ {
		ptr := sortedNames[string(i)].root
		for {
			if ptr == nil {
				break
			}
			score = score + nameScore(count, ptr.name)
			ptr = ptr.next
			count++
		}
	}
	return score
}

//Compute the score of a name
func nameScore(count int32, name string) int32 {
	var score int32
	for _, c := range name {
		score = score + c - asciiStartUppercase
	}
	return count * score
}

//readNames read the file and send the name to be added to the in-memory structure
//readNames calls addName for every name read from the file
func readNames(f *os.File) (map[string]*LinkedList, []string) {
	sortedNames := map[string]*LinkedList{
		"A": nil,
		"B": nil,
		"C": nil,
		"D": nil,
		"E": nil,
		"F": nil,
		"G": nil,
		"H": nil,
		"I": nil,
		"J": nil,
		"K": nil,
		"L": nil,
		"M": nil,
		"N": nil,
		"O": nil,
		"P": nil,
		"Q": nil,
		"R": nil,
		"S": nil,
		"T": nil,
		"U": nil,
		"V": nil,
		"W": nil,
		"X": nil,
		"Y": nil,
		"Z": nil,
	}

	recording := false
	const quotes = 34
	runeQuotes := rune(quotes)
	var name string
	names := []string{}

	for {
		buf := make([]byte, 1)
		_, err := f.Read(buf)
		if err == io.EOF {
			sort.Strings(names)
			return sortedNames, names
		}

		c := bytes.Runes(buf)[0]
		if c != runeQuotes && recording {
			name = name + string(c)
		}

		if c == runeQuotes && recording {
			//word detected, go insert it
			addName(sortedNames, name)
			names = append(names, name)

			//re-initialize
			name = ""
			recording = false
		} else if c == runeQuotes {
			recording = true
		}
	}
}

//addName adds the read-from-file name passed as the second parameter
//to the corresponding linked list (one linked list for each letter)
func addName(sortedNames map[string]*LinkedList, name string) {
	if sortedNames[string(name[0])] == nil {
		sortedNames[string(name[0])] = NewLinkedList(NewNode(name))
		return
	}

	ptr := sortedNames[string(name[0])].root
	var prev *Node
	idx := 1
	for {
		toInsert := NewNode(name)
		if ptr == nil {
			prev.next = toInsert
			break
		}
		//cases when we insert ADALBERTO (ADA, ADAH). We have to go to the
		//next name (ADAH) before we can insert ADALBERTO
		if idx == len(ptr.name) {
			prev = ptr
			ptr = ptr.next
			idx = 1
			continue
		}
		if idx > len(ptr.name) {
			toInsert.next = ptr.next
			ptr.next = toInsert
			break
		}
		if idx >= len(name) {
			if prev != nil {
				toInsert.next = ptr
				prev.next = toInsert
			} else {
				//replace root
				toInsert.next = ptr
				sortedNames[string(name[0])].root = toInsert
			}
			break
		}

		existingWordChar := ptr.name[idx]
		newWordChar := name[idx]
		if newWordChar < existingWordChar {
			toInsert.next = ptr
			if prev != nil {
				toInsert.next = ptr
				prev.next = toInsert
			} else {
				//reassing root
				toInsert.next = ptr
				sortedNames[string(name[0])].root = toInsert
			}
			break
		} else if newWordChar == existingWordChar {
			//for e.g Camille and Caryl for idx==1
			idx++
		} else {
			prev = ptr
			ptr = ptr.next
			idx = 1
		}
	}
}

//computeTruth computes the right score given an array sorted by the go programming language
func computeTruth(sortedNames []string) int32 {
	var score int32
	var count int32 = 1
	for _, v := range sortedNames {
		score = score + nameScore(count, v)
		count++
	}
	return score
}

//Was used for verification purposes
func printCounts(ll map[string]*LinkedList) map[string]int {
	res := make(map[string]int)
	for k, v := range ll {
		ptr := v.root
		count := 0
		for {
			if ptr == nil {
				break
			}
			ptr = ptr.next
			count++
		}
		res[string(k)] = count
	}
	return res
}

//Was used for verification purposes too
func printNames(ll map[string]*LinkedList) {
	for k, v := range ll {
		fmt.Printf("For letter %v: \n", string(k))
		ptr := v.root
		for {
			if ptr == nil {
				break
			}
			if ptr.next == nil {
				fmt.Printf("%v", ptr.name)
			} else {
				fmt.Printf("%v,", ptr.name)
			}
			ptr = ptr.next
		}
		fmt.Println()
	}
}

//Was used for verification purposes too
func verif(verifMap map[string]int, name string) {
	verifMap[string(name[0])]++
}

func verifMySort(mine map[string]*LinkedList, goSort []string) {
	count := 0
	for i := asciiStartUppercase + 1; i < asciiEndUppercase; i++ {
		ptr := mine[string(i)].root
		for {
			if ptr == nil {
				break
			}
			if goSort[count] != ptr.name {
				fmt.Printf("was %v and should be %v\n", ptr.name, goSort[count])
				return
			}
			ptr = ptr.next
			count++
		}
	}
}

//LinkedList starts at the root node
type LinkedList struct {
	root *Node
}

//NewLinkedList creates a new linked list with a node
func NewLinkedList(node *Node) *LinkedList {
	return &LinkedList{
		root: node,
	}
}

//Node contains a name
type Node struct {
	name string
	next *Node
}

//NewNode creates a new node
func NewNode(name string) *Node {
	return &Node{
		name: name,
		next: nil,
	}
}
