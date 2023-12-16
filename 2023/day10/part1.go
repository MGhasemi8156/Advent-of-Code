package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
    i int
    j int
}

func isValid(i, j int, h, w int) bool {
    if i < 0 || i >= h {
        return false
    }
    if j < 0 || j >= w {
        return false
    }

    return true
}

func leftValid(r rune) bool {
    switch r {
    case '-', 'J', '7', 'S': return true
    }
    return false
}

func rightValid(r rune) bool {
    switch r {
    case '-', 'L', 'F', 'S': return true
    }
    return false
}

func upValid(r rune) bool {
    switch r {
    case '|', 'L', 'J', 'S': return true
    }
    return false
}

func downValid(r rune) bool {
    switch r {
    case '|', 'F', '7', 'S': return true
    }
    return false
}

func findUnmarkedNeighbors(i, j int, marked [][]int, maze []string) []node {
    ne := make([]node, 0)

    h, w := len(maze), len(maze[0])
    r := rune(maze[i][j])
    if leftValid(r) && j-1 >= 0 && marked[i][j-1] == -1 && rightValid(rune(maze[i][j-1])) {
        ne = append(ne, node{i, j-1})
    } 
    if upValid(r) && i-1 >= 0 && marked[i-1][j] == -1 && downValid(rune(maze[i-1][j])) {
        ne = append(ne, node{i-1, j})
    } 
    if rightValid(r) && j+1 < w && marked[i][j+1] == -1 && leftValid(rune(maze[i][j+1])) {
        ne = append(ne, node{i, j+1}) 
    } 
    if downValid(r) && i+1 < h && marked[i+1][j] == -1 && upValid(rune(maze[i+1][j])) {
        ne = append(ne, node{i+1, j})
    } 
    return ne
}

func BFS(i, j int, marked [][]int, maze []string) int {
    next := make([]node, 0)
    current := []node{{i, j}}
    level := 0

    for len(current) > 0 {
        fmt.Println("level: ", level)
        for _, v := range current {
            fmt.Print(v, ", ")
            if marked[v.i][v.j] == -1 {
                marked[v.i][v.j] = level
                next = append(next, findUnmarkedNeighbors(v.i, v.j , marked, maze)...)
            }
        }
        fmt.Println()
        current = next
        next = make([]node, 0)
        level++
    }

    return level
}

func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    maze := make([]string, 0)

    for scanner.Scan() {
        maze = append(maze, scanner.Text())
    }

    var si, sj int
    marked := make([][]int, len(maze))
    for i := 0; i < len(maze); i++ {
        marked[i] = make([]int, len(maze[i]))
        for j, v := range maze[i] {
            marked[i][j] = -1
            if v == 'S' {
                si, sj = i, j
            }
        }
    }

    fmt.Println(BFS(si, sj, marked, maze)-1)
}
