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

func parentedBFS(i, j int, marked [][]int, maze []string, parents map[node][]node) int {
    next := make([]node, 0)
    current := []node{{i, j}}
    level := 0

    var ne []node
    for len(current) > 0 {
        for _, v := range current {
            if marked[v.i][v.j] == -1 {
                marked[v.i][v.j] = level
                ne = findUnmarkedNeighbors(v.i, v.j , marked, maze)
                for _, n := range ne {
                    parents[n] = append(parents[n], v)
                }
                next = append(next, ne...)
            }
        }

        current = next
        next = make([]node, 0)
        level++
    }

    return level
}

func SolvePart2(filesrc string) {
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

    parents := make(map[node][]node)
    target := parentedBFS(si, sj, marked, maze, parents)-1 
    
    var ti , tj int 
    loop := make([][]int, len(marked))
    for i := 0; i < len(marked); i++ {
        loop[i] = make([]int, len(marked[i]))
        for j := 0; j < len(marked[0]); j++ {
            if marked[i][j] == target {
                ti, tj = i , j
            }
        }
    }
    
    currnet := []node{{ti, tj}}
    next := make([]node, 0)
    level := 0
    for len(currnet) > 0 {
        for i, v := range currnet {
            if i == 0 {
                loop[v.i][v.j] = target + level
            } else {
                loop[v.i][v.j] = target - level
            }
            next = append(next, parents[v]...)
        }
        
        currnet = next
        next = make([]node, 0)
        level++
    }

    for {
        v := loop[si][sj]
        if v == 0 {
            loop[si][sj] = target + level - 1
        }

        switch {
        case si-1 >=0 && loop[si-1][sj] == v + 1: {
            for i := 1; loop[si][sj+i] == 0 || loop[si-1][sj+i] == 0; i++ {
                if loop[si][sj+i] == 0 {
                    loop[si][sj+i] = -1
                }
                if loop[si-1][sj+i] == 0 {
                    loop[si-1][sj+i] = -1
                }
            }
            si--
        }
        case sj-1 >= 0 && loop[si][sj-1] == v + 1: {
            for i := 1; loop[si-i][sj] == 0 || loop[si-i][sj-1] == 0; i++ {
                if loop[si-i][sj] == 0 {
                    loop[si-i][sj] = -1
                } 
                if loop[si-i][sj-1] == 0 {
                    loop[si-i][sj-1] = -1
                }
            }
            sj--
        }
        case si+1 < len(loop) && loop[si+1][sj] == v + 1: si++
        case sj+1 < len(loop[0]) && loop[si][sj+1] == v + 1: sj++


    } 

        if loop[si][sj] == target + level - 1 {
            break
        }
    } 

    encNum := 0
    for i := 0; i < len(loop); i++ {
        for j := 0; j < len(loop[0]); j++ {
            if loop[i][j] == -1 {
                encNum++
            }
        }
    }

    fmt.Println(encNum)
}
