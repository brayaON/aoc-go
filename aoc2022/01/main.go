package aoc2022

import ( 
    "os"
    "fmt"
    "bufio"
    "strconv"
    "log"
    "sort"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/01"
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

func SolveP1() {

    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    scanner := bufio.NewScanner(fileReader)

    var maxCals, currCals int64

    for scanner.Scan() {
	str := scanner.Text()
	if len(str) > 0 {
	    cal, _ := strconv.ParseInt(str, 10, 32)
	    currCals += cal
	} else {
	    maxCals = max(maxCals, currCals)
	    currCals = 0
	}
    }

    fileOut, err := os.Create(relativePath+"/data1.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, maxCals)

    if err := scanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}

func SolveP2() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)

    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)


    var elfCalories [] int
    currentCals := 0

    for fileScanner.Scan() {
	str := fileScanner.Text()
	if len(str) > 0 {
	    cal, err := strconv.Atoi(str)
	    checkError(err)

	    currentCals += cal
	} else {
	    elfCalories = append(elfCalories, currentCals)
	    currentCals = 0
	}
    }

    sort.Ints(elfCalories)

    totalCals := 0
    for i := len(elfCalories)-3; i < len(elfCalories); i++ {
	totalCals += elfCalories[i]
    }

    fileOut, err := os.Create(relativePath+"/data2.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, totalCals)
}
