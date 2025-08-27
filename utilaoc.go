package util

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func Read(file string) string {
	return strings.TrimSpace(string(Must(os.ReadFile("input.txt"))))
}

func Grid(input string) [][]uint8 {
	chars := [][]uint8{}
	for _, line := range NonEmptyLines(input) {
		chars = append(chars, []uint8(line))
	}
	return chars
}

// Deprecated: use [grid.Grid].
func GridStr(G [][]uint8) string {
	var sb strings.Builder
	for i, row := range G {
		if i != 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(string(row))
	}
	return sb.String()
}

// returns non-empty lines
func NonEmptyLines(s string) (lines []string) {
	s = strings.TrimSpace(s)
	for _, line := range strings.Split(s, "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func Blocks(s string) (blocks []string) {
	rawBlocks := strings.Split(s, "\n\n")
	for _, block := range rawBlocks {
		block = strings.TrimSpace(block)
		if block != "" {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func Num(s string) int {
	return Must(strconv.Atoi(s))
}

func Numbers(input string, separator string) [][]int {
	numbers := [][]int{}
	for _, line := range NonEmptyLines(input) {
		tokens := strings.Split(line, separator)
		var row []int
		for i := range tokens {
			// unlike [strings.Fields], splitting by " " keeps empty strings
			if tokens[i] != "" {
				row = append(row, Num(tokens[i]))
			}
		}
		numbers = append(numbers, row)
	}
	return numbers
}

func Transposed[V any](m [][]V) [][]V {
	rows := len(m[0])
	cols := len(m)

	res := make([][]V, rows)
	for r := range rows {
		res[r] = make([]V, cols)
	}

	for x := range m {
		for y := range m[x] {
			res[y][x] = m[x][y]
		}
	}
	return res
}

// find Least Common Multiple (LCM) via GCD
// https://go.dev/play/p/SmzvkDjYlb
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	integers = integers[2:]
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
// https://go.dev/play/p/SmzvkDjYlb
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}
