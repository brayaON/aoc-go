package aoc2022

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/05"
)

func checkError(err error) {
    if err != nil {
	log.Fatal(nil)
    }
}

func reverseSlice(s []string) []string {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
	s[i], s[j] = s[j], s[i]
    }

    return s
}

func SolveP1() {

    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    cargo := make([][]string, 10)

    movesStarted := false

    for fileScanner.Scan() {
	str := fileScanner.Text()

	if len(str) > 0 && !movesStarted {
	    line := strings.Split(str, " ")
	    col, empty := 0, 0
	    for _, value := range line {
		if len(value) == 0 {
		    empty++
		} else if len(value) > 1 {
		    sliceLine := []byte(value)
		    cargo[col] = append(cargo[col], string(sliceLine[1]))
		    col++
		}
		if empty == 4 {
		    col++
		    empty = 0
		}
	    }
	} else {
	    movesStarted = true

	    if len(str) > 0 {
		line := strings.Split(str, " ")
		amnt, err := strconv.Atoi(line[1])
		checkError(err)
		from, err := strconv.Atoi(line[3])
		checkError(err)
		to, err := strconv.Atoi(line[5])
		checkError(err)

		stackAux := cargo[from - 1][len(cargo[from - 1])-amnt:]
		stackAux = reverseSlice(stackAux) // Remove this to solve the second part of the day

		cargo[from-1] = cargo[from - 1][:len(cargo[from - 1]) - amnt] // Remove amount from stackFrom
		cargo[to-1] = append(cargo[to-1], stackAux...) // Add amount to stackTo

	    } else {
		for _, stack := range cargo {
		    stack = reverseSlice(stack)
		}
	    }
	}
    }
    
    ans := ""
    for _, stack := range cargo {
	if len(stack) > 0 {
	    ans += stack[len(stack)-1]
	}
    }

    fileOut, err := os.Create(relativePath+"/data1.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, ans)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}
