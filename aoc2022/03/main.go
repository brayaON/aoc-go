package aoc2022

import ( 
    "os"
    "fmt"
    "bufio"
    "log"
    "unicode"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/03"
)

func checkError(err error) {
    if err != nil {
	log.Fatal(nil)
    }
}

func SolveP1() {

    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)


    var totalPriority int32 = 0
    for fileScanner.Scan() {
	str := fileScanner.Text()
	left, right := str[:len(str)/2], str[len(str)/2:]

	exists := make(map[rune]int32)
	for _, runeValue := range left {
	    var priority int32 = 0
	    if unicode.IsUpper(runeValue) {
		runeA := 'A'
		priority = runeValue - runeA + 27
	    } else {
		runeA := 'a'
		priority = runeValue - runeA + 1
	    }
	    exists[runeValue] = priority
	}
	
	for _, runeValue := range right {
	    if value, is := exists[runeValue]; is {
		totalPriority += value
		exists[runeValue] = 0
	    }
	}
    }

    fileOut, err := os.Create(relativePath+"/data1.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, totalPriority)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}

func SolveP2() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    var totalPriority int32 = 0
    counter := 0

    elfGroup := make([]string, 0)

    for fileScanner.Scan() {
	str := fileScanner.Text()
	elfGroup = append(elfGroup, str)
	counter += 1

	if counter == 3 {
	    elf1, elf2, elf3 := elfGroup[0], elfGroup[1], elfGroup[2]

	    exists := make(map[rune]int32)
	    for _, runeValue := range elf1 {
		var priority int32 = 0
		if unicode.IsUpper(runeValue) {
		    runeA := 'A'
		    priority = runeValue - runeA + 27
		} else {
		    runeA := 'a'
		    priority = runeValue - runeA + 1
		}
		exists[runeValue] = priority
	    }

	    exists2 := make(map[rune]int32)

	    for _, runeValue := range elf2 {
		if value, is := exists[runeValue]; is {
		    exists2[runeValue] = value
		}
	    }

	    for _, runeValue := range elf3 {
		if value, is := exists2[runeValue]; is {
		    totalPriority += value
		    break
		}
	    }
	    counter = 0
	    elfGroup = nil
	}
    }

    fileOut, err := os.Create(relativePath+"/data2.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, totalPriority)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}

