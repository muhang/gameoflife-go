package main

import (
    "fmt"
    "bufio"
    "os"
    "os/exec"
    "time"
    "errors"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func countLiveNeighbors(y int, x int, grid [][]rune) int {
    count := 0

    // get wrapped indices for 4 directions
    up := y - 1
    if up < 0 {
        up = len(grid) - 1
    }

    down := y + 1
    if down >= len(grid) {
        down = 0
    }

    left := x - 1
    if left < 0 {
        left = len(grid) - 1
    }

    right := x + 1
    if right >= len(grid) {
        right = 0
    }

    count = checkForLife(grid[up][left], count)
    count = checkForLife(grid[y][left], count)
    count = checkForLife(grid[down][left], count)

    count = checkForLife(grid[up][x], count)
    count = checkForLife(grid[down][x], count)

    count = checkForLife(grid[up][right], count)
    count = checkForLife(grid[y][right], count)
    count = checkForLife(grid[down][right], count)

    return count
}

func checkForLife(r rune, c int) int {
    if string(r) == "#" {
        c++
    }
    return c
}

func handleCount(c int, r rune) rune {
    if string(r) == "." {   // dead
        if c == 3 {
            r = []rune("#")[0]
        }
    } else {    // alive
        if c < 2 {
            r = []rune(".")[0]
        } else if c > 3 {
            r = []rune(".")[0]
        }
    }

    return r
}

func tick(grid [][]rune) [][]rune {
    clone := make([][]rune, len(grid))
    for i, row := range grid {
        clone[i] = make([]rune, len(grid))
        copy(clone[i], row)
    }
    for i, row := range grid {
        for j, char := range row {
            c := countLiveNeighbors(i, j, grid)
            clone[i][j] = handleCount(c, char)
        }
    }
    return clone
}

func readInput(path string) [][]rune {
    grid := [][]rune{}

    file, err := os.Open(path)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        row := []rune(scanner.Text())
        grid = append(grid, row)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return grid
}

func printGrid(grid [][]rune) {
    for _, row := range grid {
        for _, char := range row {
            fmt.Print(string(char))
        }
        fmt.Print("\n")
    }
}

func resetOutput() {
    c := exec.Command("cmd", "/c", "cls")
    c.Stdout = os.Stdout
    c.Run()
}

func runGame(iterations int, grid [][]rune) {
    resetOutput()
    fmt.Println("Gen 0")
    printGrid(grid)

    framegap := time.Second/4

    time.Sleep(framegap)

    for i := 0; i < iterations; i++ {
        resetOutput()
        grid = tick(grid)
        fmt.Println("Gen", i+1)
        printGrid(grid)

        time.Sleep(framegap)
    }
}

func parseArgs (args []string) (string, error) {
    if len(args) == 0 {
        return "", errors.New("Need file path")
    }

    return args[0], nil
}

func main() {
    // takes arg fp, path to .txt with initial structure
    fp, err := parseArgs(os.Args[1:])
    if err != nil {
        panic(err)
    }
    grid := readInput(fp)
    runGame(7, grid)
}