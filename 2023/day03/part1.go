package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)


func isSymbol(r rune) bool {
    return !unicode.IsDigit(r) && r != '.' 
}

func isNearSymbol(i int, linebuf [3]string) bool {
    switch {
    case linebuf[0] != "" && isSymbol(rune(linebuf[0][i])): return true
    case linebuf[0] != "" && i > 0 && isSymbol(rune(linebuf[0][i-1])): return true
    case linebuf[0] != "" && i < len(linebuf[0])-1 && isSymbol(rune(linebuf[0][i+1])): return true
    case linebuf[2] != "" && isSymbol(rune(linebuf[2][i])): return true
    case linebuf[2] != "" && i > 0 && isSymbol(rune(linebuf[2][i-1])): return true
    case linebuf[2] != "" && i < len(linebuf[2])-1 && isSymbol(rune(linebuf[2][i+1])): return true
    case i > 0 && isSymbol(rune(linebuf[1][i-1])): return true
    case i < len(linebuf[1])-1 && isSymbol(rune(linebuf[1][i+1])): return true
    }

    return false
}

func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var linebuf [3]string
    var num, sum int64
    var nearSymbol bool

    for scanner.Scan() || linebuf[1] != ""{
        linebuf[0] = linebuf[1]
        linebuf[1] = linebuf[2]
        linebuf[2] = scanner.Text()

        if linebuf[1] != "" {
            num = 0
            nearSymbol = false  
            for i, v := range linebuf[1] {
                if unicode.IsDigit(v) {
                    digit, _ := strconv.Atoi(string(v))
                    num = num * 10 + int64(digit)
                    nearSymbol = nearSymbol || isNearSymbol(i, linebuf)
                } else {
                    if nearSymbol {
                        sum += num
                        fmt.Println(num)
                    }
                    num = 0
                    nearSymbol = false
                }
                if nearSymbol && i == len(linebuf[1]) - 1 { 
                    sum += num
                }

            }
        }
    }

    fmt.Println(sum)
}
