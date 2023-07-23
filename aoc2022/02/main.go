package aoc2022

import ( 
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/02"
    pointsChoice = map[string]int { "X": 1, "Y": 2, "Z": 3 }
)

func max(a, b int64) int64 {
    if a > b {
	return a
    }
    return b
}

func checkError(err error) {
    if err != nil {
	log.Fatal(nil)
    }
}

// Rock -> A, X
// Paper -> B, Y
// Scissors -> C, Z

// Loss: 0
// Draw: 3
// Win: 6
// Rock: 1
// Paper: 2
// Scissor: 3

func SolveP1() {

    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    totalPoints := 0
    for fileScanner.Scan() {
	str := fileScanner.Text()
	strategy := strings.Split(str, " ")

	op, me := strategy[0], strategy[1]
	totalPoints += pointsChoice[me]

	switch me {
	case "X":
	    if op == "C" {
		totalPoints += 6
	    } else if op == "A" {
		totalPoints += 3
	    }
	case "Y":
	    if op == "A" {
		totalPoints += 6
	    } else if op == "B" {
		totalPoints += 3
	    }
	case "Z":
	    if op == "B" {
		totalPoints += 6
	    } else if op == "C" {
		totalPoints += 3
	    }
	}
    }

    fileOut, err := os.Create(relativePath+"/data1.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, totalPoints)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}

func SolveP2() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    totalPoints := 0
    for fileScanner.Scan() {
	str := fileScanner.Text()
	strategy := strings.Split(str, " ")

	op, me := strategy[0], strategy[1]

	switch me {
	case "X":
	    if op == "C" {
		totalPoints += 2
	    } else if op == "A" {
		totalPoints += 3
	    } else {
		totalPoints += 1
	    }
	case "Y":
	    totalPoints += 3
	    if op == "C" {
		totalPoints += 3
	    } else if op == "A" {
		totalPoints += 1
	    } else {
		totalPoints += 2
	    }
	case "Z":
	    totalPoints += 6
	    if op == "C" {
		totalPoints += 1
	    } else if op == "A" {
		totalPoints += 2
	    } else {
		totalPoints += 3
	    }
	}
    }

    fileOut, err := os.Create(relativePath+"/data2.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, totalPoints)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}

