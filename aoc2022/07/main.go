package aoc2022

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
    "math"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/07"
)

func checkError(err error) {
    if err != nil {
	log.Fatal(err)
    }
}

func min(a, b int64) int64 {
    if a < b {
	return a
    }

    return b
}

const (
    MAX_SIZE = 70000000
    MAX_UNUSED = 30000000
    MAX_ALLOWED = 100000
)

func SolveP1() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)

    fileOut, err := os.Create(relativePath+"/data1.out")
    checkError(err)
    defer fileOut.Close()

    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)


    sizes := make(map[string]int64) // Directory ID -> Size
    prefixes := make([]string, int64(0)) // stack to identify directories

    prefix := ""
    oldPrefix := ""

    for fileScanner.Scan() {
	str := fileScanner.Text()
	tokens := strings.Split(str, " ")

	if tokens[0] == "$" && tokens[1] == "cd" {
	    out := false
	    if tokens[2] != ".." {
		prefixes = append(prefixes, tokens[2])
	    } else {
		out = true
		prefixes = prefixes[:len(prefixes)-1]
	    }

	    // Add the children directory's size to the parent's
	    oldPrefix = prefix
	    prefix = ""
	    for _, value := range prefixes {
		prefix += value
	    }
	    if out {
		sizes[prefix] += sizes[oldPrefix]
	    }

	} else if tokens[0] != "$" {
	    // Check if the first token is a number and add it to the diirectory's size
	    if size, err := strconv.Atoi(tokens[0]); err == nil {
		sizes[prefix] += int64(size)
	    }
	}
    }

    var ans1 int64
    for _, v := range sizes {
	if v <= MAX_ALLOWED {
	    ans1 += v
	}
    }

    fmt.Fprintln(fileOut, ans1)

    if err := fileScanner.Err(); err != nil {
	log.Fatal(err)
    }
}

func SolveP2() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)

    fileOut, err := os.Create(relativePath+"/data2.out")
    checkError(err)
    defer fileOut.Close()

    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)


    sizes := make(map[string]int64) // Directory ID -> Size
    prefixes := make([]string, int64(0)) // stack to identify directories

    prefix := ""
    oldPrefix := ""

    for fileScanner.Scan() {
	str := fileScanner.Text()
	tokens := strings.Split(str, " ")

	if tokens[0] == "$" && tokens[1] == "cd" {
	    out := false
	    if tokens[2] != ".." {
		prefixes = append(prefixes, tokens[2])
	    } else {
		out = true
		prefixes = prefixes[:len(prefixes)-1]
	    }

	    // Add the children directory's size to the parent's
	    oldPrefix = prefix
	    prefix = ""
	    for _, value := range prefixes {
		prefix += value
	    }
	    if out {
		sizes[prefix] += sizes[oldPrefix]
	    }

	} else if tokens[0] != "$" {
	    // Check if the first token is a number and add it to the diirectory's size
	    if size, err := strconv.Atoi(tokens[0]); err == nil {
		sizes[prefix] += int64(size)
	    }
	}
    }

    // Process the remaining directories to compute the total size
    for len(prefixes) > 1 {
	oldPrefix = prefix
	prefixes = prefixes[:len(prefixes)-1]

	prefix = ""
	for _, value := range prefixes {
	    prefix += value
	}
	sizes[prefix] += sizes[oldPrefix]
    }

    var ans2 int64 = math.MaxInt64
    var totalSpace int64 = sizes["/"] 

    freeSpace := MAX_SIZE - totalSpace
    requiredSpace := MAX_UNUSED - freeSpace

    // Find the min directory to delete that meets the condition
    for _, val := range sizes {
	if val >= requiredSpace {
	    ans2 = min(val, ans2)
	}
    }

    fmt.Fprintln(fileOut, ans2)

    if err := fileScanner.Err(); err != nil {
	log.Fatal(err)
    }
}
