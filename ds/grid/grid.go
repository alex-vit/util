package grid

import (
	"fmt"
	"iter"
	"strings"
	"unicode"
)

// Grid is an experimental class to work with Advent of Code grid type puzzles.
type Grid struct {
	A    [][]byte
	R, C int
}

func Parse(s string) *Grid {
	lines := strings.FieldsFunc(s, func(r rune) bool { return r == '\n' })
	a := make([][]byte, 0, len(lines))
	for _, line := range lines {
		a = append(a, []byte(line))
	}
	return New(a)
}

func Size(r, c int) *Grid {
	A := make([][]byte, 0, r)
	for range r {
		A = append(A, make([]byte, c))
	}
	return New(A)
}

func New(A [][]byte) *Grid {
	if len(A) == 0 || len(A[0]) == 0 {
		panic("empty grid")
	}
	for i, bs := range A {
		for _, b := range bs {
			if b > unicode.MaxASCII {

				panic(fmt.Sprintf("line contains unicode, use runes instead: %q", string(bs)))
			}
		}
		if i > 0 {
			prevBs := A[i-1]
			if len(prevBs) != len(bs) {
				prevLine, line := string(prevBs), string(bs)
				panic(fmt.Sprintf(`lines of diffrent length:
%d. (%d) %q
%d. (%d) %q
`,
					i, len(prevLine), prevLine,
					i+1, len(line), line))
			}
		}
	}

	R, C := len(A), len(A[0])
	g := &Grid{A: A, R: R, C: C}
	return g
}

func (g Grid) Clone() *Grid {
	a := make([][]byte, 0, g.R)
	for r := range g.R {
		a = append(a, make([]byte, g.C))
		copy(a[r], g.A[r])
	}
	return New(a)
}

func (g *Grid) Walk(visit func(r, c int, b byte) (continue_ bool)) {
	for r := range g.R {
		for c := range g.C {
			continue_ := visit(r, c, g.A[r][c])
			if !continue_ {
				return
			}
		}
	}
}

func (g *Grid) Count(b byte) (count int) {
	for r := range g.R {
		for c := range g.C {
			if b == g.A[r][c] {
				count++
			}
		}
	}
	return count
}

func (g Grid) InBounds(r, c int) bool {
	return r >= 0 && r < g.R && c >= 0 && c < g.C
}

// Isb tells you whether B is at the coordinates R, C. Checks bounds.
//
// Deprecated: use [Grid.Is].
func (g Grid) Isb(r, c int, b byte) bool {
	return g.InBounds(r, c) && g.A[r][c] == b
}

func (g Grid) Is(r, c int, oneOf string) bool {
	return g.InBounds(r, c) && strings.Contains(oneOf, string(g.A[r][c]))
}

func (g *Grid) Place(r, c int, b byte) (placed bool) {
	if g.InBounds(r, c) {
		g.A[r][c] = b
		return true
	} else {
		return false
	}
}

func (g *Grid) Clear(with byte) {
	g.Walk(func(r, c int, b byte) (continue_ bool) {
		g.A[r][c] = with
		return true
	})
}

func (g *Grid) Replace(what, with byte) {
	g.Walk(func(r, c int, b byte) (continue_ bool) {
		if g.Isb(r, c, what) {
			g.A[r][c] = with
		}
		return true
	})
}

func (g Grid) String() string {
	var sb strings.Builder
	for i, row := range g.A {
		if i != 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(string(row))
	}
	return sb.String()
}

func (g Grid) Find(anyOf string) (r, c int) {
	r, c = -1, -1
	g.Walk(func(r_, c_ int, b byte) (continue_ bool) {
		if strings.Contains(anyOf, string(b)) {
			r, c = r_, c_
			return false
		}
		return true
	})

	if r == -1 {
		panic(anyOf + " not found in:\n" + g.String())
	}
	return r, c
}

var deltas = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (g Grid) Coord(r, c, dir int) (nr, nc int) {
	dir = norm(dir)
	dr, dc := deltas[dir][0], deltas[dir][1]
	nr, nc = r+dr, c+dc
	return nr, nc
}

type Tile struct {
	R, C int
	B    byte
}

func (t Tile) String() string {
	return fmt.Sprintf("{%d %d %s}", t.R, t.C, string(t.B))
}

func (g Grid) Direction(r, c, dir int) iter.Seq[Tile] {
	return func(yield func(Tile) bool) {
		dir := norm(dir)
		for nr, nc := r, c; g.InBounds(nr, nc); nr, nc = g.Coord(nr, nc, dir) {
			if !yield(Tile{nr, nc, g.A[nr][nc]}) {
				return
			}
		}
	}
}

// Returns a slice of {row, col, direction}
func (g *Grid) Neighbors(r int, c int) (rcds [][3]int) {
	for direction := range 4 {
		nr, nc := g.Coord(r, c, direction)
		if g.InBounds(nr, nc) {
			rcds = append(rcds, [3]int{nr, nc, direction})
		}
	}
	return rcds
}

func norm(dir int) int {
	if dir == -1 {
		dir = 3
	} else if dir > 3 {
		dir = dir % 4
	}
	return dir
}
