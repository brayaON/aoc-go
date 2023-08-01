package aoc2022

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
    wd, err = os.Getwd()
    relativePath = wd + "/aoc2022/08"
)

const (
    N = 99
)

func checkError(err error) {
    if err != nil {
	log.Fatal(err)
    }
}

func checkLX(i, j int, grid [][]int) (int, bool) {
    totalTrees := 0
    visible := true

    for x := j-1; x >= 0; x-- {
	if grid[i][x] < grid[i][j] {
	    totalTrees++
	}
	if grid[i][x] >= grid[i][j] {
	    visible = false
	    totalTrees++
	    break
	}
    }
    return totalTrees, visible
}

func checkRX(i, j int, grid [][]int) (int, bool) {
    totalTrees := 0
    visible := true

    for x := j+1; x < N; x++ {
	if grid[i][x] < grid[i][j] {
	    totalTrees++
	}

	if grid[i][x] >= grid[i][j] {
	    visible = false
	    totalTrees++
	    break
	}
    }
    return totalTrees, visible
}

func checkUY(i, j int, grid [][]int) (int, bool) {
    totalTrees := 0
    visible := true

    for y := i-1; y >= 0; y-- {
	if grid[y][j] < grid[i][j] {
	    totalTrees++
	} 
	if grid[y][j] >= grid[i][j] {
	    visible = false
	    totalTrees++
	    break
	}
    }

    return totalTrees, visible
}

func checkDY(i, j int, grid [][]int) (int, bool) {
    totalTrees := 0
    visible := true

    for y := i+1; y < N; y++ {
	if grid[y][j] < grid[i][j] {
	    totalTrees++
	} 
	if grid[y][j] >= grid[i][j] {
	    visible = false
	    totalTrees++
	    break
	}
    }
    return totalTrees, visible
}

func SolveP1() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)

    fileOut, err := os.Create(relativePath+"/data1.out")
    checkError(err)
    defer fileOut.Close()

    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    grid := make([][]int, N)
    row := 0

    for fileScanner.Scan() {
	str := fileScanner.Text()

	for _, num := range str {
	    intNum, err := strconv.Atoi(string(num))
	    checkError(err)

	    grid[row] = append(grid[row], intNum)
	}
	row++
    }

    ans1 := 4 * (N - 1)

    for i := 1; i < row - 1; i++ {
	for j := 1; j < row - 1; j++ {
	    if _, visible := checkLX(i, j, grid); visible {
		ans1++
	    } else if _, visible := checkRX(i, j, grid); visible {
		ans1++
	    } else if _, visible := checkUY(i, j, grid);  visible {
		ans1++
	    } else if _, visible := checkDY(i, j, grid); visible {
		ans1++
	    }
	}
    }

    fmt.Fprintln(fileOut, ans1)

    if err := fileScanner.Err(); err != nil {
	log.Fatal(err)
    }
}

func max(a, b int) int {
    if a > b {
	return a
    }

    return b
}

func SolveP2() {
    fileIn, err := os.Open(relativePath+"/data.in")
    checkError(err)

    fileOut, err := os.Create(relativePath+"/data2.out")
    checkError(err)
    defer fileOut.Close()

    fileReader := bufio.NewReader(fileIn)
    fileScanner := bufio.NewScanner(fileReader)

    grid := make([][]int, N)
    row := 0

    for fileScanner.Scan() {
	str := fileScanner.Text()

	for _, num := range str {
	    intNum, err := strconv.Atoi(string(num))
	    checkError(err)

	    grid[row] = append(grid[row], intNum)
	}
	row++
    }

    var ans2 int = 0

    for i := 1; i < row - 1; i++ {
	for j := 1; j < row - 1; j++ {
	    treeViewLX, _ := checkLX(i, j, grid)
	    treeViewRX, _ := checkRX(i, j, grid)
	    treeViewUY, _ := checkUY(i, j, grid)
	    treeViewDY, _ := checkDY(i, j, grid)

	    ans2 = max(ans2, treeViewDY * treeViewLX * treeViewRX * treeViewUY)
	}
    }

    fmt.Fprintln(fileOut, ans2)

    if err := fileScanner.Err(); err != nil {
	log.Fatal(err)
    }
}
