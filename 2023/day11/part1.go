package day11

import (
	"bufio"
	"log"
	"os"
    "fmt"
)

type galaxy struct {
    i int
    j int
}

func distance(a, b galaxy, vmark, hmark []int) int {
    di := (a.i + vmark[a.i]) - (b.i + vmark[b.i])
    if di < 0 {
        di = -di
    }
    dj := (a.j + hmark[a.j]) - (b.j + hmark[b.j])
    if dj < 0 {
        dj = -dj
    }
    
    return di + dj
}

func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    galaxies := make([]galaxy, 0)
    hmark := make([]int, 0) 
    vmark := make([]int, 0)
    for i := 0; scanner.Scan(); i++ {
        line := scanner.Text()

        vmark = append(vmark, 1)
        // make sure hmark has line size
        if len(hmark) < len(line) {
            for k := 0; k < len(line); k++ {
                hmark = append(hmark, 1)
            }
        }

        for j, v := range line {
            if v == '#' {
                vmark[i] = 0
                hmark[j] = 0
                galaxies = append(galaxies, galaxy{i, j})
            }
        }
    }

    for i := 1; i < len(hmark); i++ {
        hmark[i] += hmark[i-1]
    }
    
    for i := 1; i < len(vmark); i++ {
        vmark[i] += vmark[i-1]
    }
    
    sum := 0
    for i := 0; i < len(galaxies); i++ {
        for j := i+1; j < len(galaxies); j++ {
            sum += distance(galaxies[i], galaxies[j], vmark, hmark)
        }
    }
    
    fmt.Println(sum)
}
