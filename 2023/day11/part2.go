package day11

import (
	"bufio"
	"log"
	"os"
    "fmt"
)

func SolvePart2(filesrc string) {
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

        vmark = append(vmark, 999999)
        // make sure hmark has line size
        if len(hmark) < len(line) {
            for k := 0; k < len(line); k++ {
                hmark = append(hmark, 999999)
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
