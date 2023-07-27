package aoc2022

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/06"
)

func checkError(err error) {
    if err != nil {
	log.Fatal(nil)
    }
}

func isUnique(data []byte) bool {
    exists := make(map[byte]bool)

    ans := true

    for _, val := range data {
	if _, is := exists[val]; is {
	    ans = false
	    break
	}
	exists[val] = true
    }

    return ans
}

func SolveP1() {

    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    ans := 0

    for fileScanner.Scan() {
	str := fileScanner.Text()
	
	streamData := []byte(str)

	for i := 3; i < len(streamData); i++ {
	    if isUnique(streamData[i - 3: i + 1]) {
		ans = i + 1
		break
	    }
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

func SolveP2() {

    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    ans := 0

    for fileScanner.Scan() {
	str := fileScanner.Text()
	
	streamData := []byte(str)

	for i := 13; i < len(streamData); i++ {
	    if isUnique(streamData[i - 13: i + 1]) {
		ans = i + 1
		break
	    }
	}
    }
    
    fileOut, err := os.Create(relativePath+"/data2.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, ans)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}
