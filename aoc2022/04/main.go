package aoc2022

import ( 
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
    "strconv"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/04"
)

func checkError(err error) {
    if err != nil {
	log.Fatal(nil)
    }
}

type pair struct {
    first int
    second int
}

func SolveP1() {

    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    var totalPairs int32 = 0
    for fileScanner.Scan() {
	str := fileScanner.Text()

	pairs := strings.Split(str, ",")
	first, second := pairs[0], pairs[1]

	psli1 := strings.Split(first, "-")
	psli2 := strings.Split(second, "-")

	pair1 := pair{}
	pair2 := pair{}

	pair1.first, _ = strconv.Atoi(psli1[0])
	pair1.second, _ = strconv.Atoi(psli1[1])

	pair2.first, _ = strconv.Atoi(psli2[0])
	pair2.second, _ = strconv.Atoi(psli2[1])

	if pair1.first >= pair2.first && pair1.second <= pair2.second {
	    totalPairs += 1
	} else if pair2.first >= pair1.first && pair2.second <= pair1.second {
	    totalPairs += 1
	}
    }

    fileOut, err := os.Create(relativePath+"/data1.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, totalPairs)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}

func SolveP2() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)
    
    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    var totalPairs int32 = 0
    
    for fileScanner.Scan() {
	str := fileScanner.Text()

	pairs := strings.Split(str, ",")
	first, second := pairs[0], pairs[1]

	psli1 := strings.Split(first, "-")
	psli2 := strings.Split(second, "-")

	pair1 := pair{}
	pair2 := pair{}

	pair1.first, _ = strconv.Atoi(psli1[0])
	pair1.second, _ = strconv.Atoi(psli1[1])

	pair2.first, _ = strconv.Atoi(psli2[0])
	pair2.second, _ = strconv.Atoi(psli2[1])

	if pair1.first >= pair2.first && pair1.first <= pair2.second {
	    totalPairs += 1
	} else if pair2.first >= pair1.first && pair2.first <= pair1.second {
	    totalPairs += 1
	}
    }

    fileOut, err := os.Create(relativePath+"/data2.out")
    checkError(err)

    defer fileOut.Close()

    fmt.Fprintln(fileOut, totalPairs)

    if err := fileScanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading stdin:", err)
    }
}

