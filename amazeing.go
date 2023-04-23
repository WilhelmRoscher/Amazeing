package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

const (
	empty    = byte(0)
	start    = byte(1)
	end      = byte(2)
	path     = byte(3)
	solution = byte(4)
)

type point struct {
	x int
	y int
}

func (p point) neighbors() []point {
	ns := make([]point, 4)

	ns[0] = point{p.x - 1, p.y}
	ns[1] = point{p.x + 1, p.y}
	ns[2] = point{p.x, p.y - 1}
	ns[3] = point{p.x, p.y + 1}

	return ns
}

var blue = color.RGBA{0, 0, 255, 255}
var green = color.RGBA{0, 255, 0, 255}
var red = color.RGBA{255, 0, 0, 255}
var black = color.RGBA{0, 0, 0, 255}
var white = color.RGBA{255, 255, 255, 255}

func cmpPixelToColor(p color.Color, c color.RGBA) bool {
	pR, pG, pB, pA := p.RGBA()
	cR, cG, cB, cA := c.RGBA()

	return cR == pR && cG == pG && cB == pB && cA == pA
}

func readMaze(p string) ([][]byte, error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	width, height := img.Bounds().Max.X, img.Bounds().Max.Y

	m := make([][]byte, width)
	for x := range m {
		m[x] = make([]byte, height)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch {
			case cmpPixelToColor(img.At(x, y), blue):
				m[x][y] = start
			case cmpPixelToColor(img.At(x, y), green):
				m[x][y] = end
			case cmpPixelToColor(img.At(x, y), black):
				m[x][y] = path
			default:
				m[x][y] = empty
			}
		}
	}

	return m, nil
}

func writeMaze(p string, m [][]byte) error {
	file, err := os.Create(p)
	if err != nil {
		return err
	}
	defer file.Close()

	width, height := len(m), len(m[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch m[x][y] {
			case start:
				img.Set(x, y, blue)
			case end:
				img.Set(x, y, green)
			case path:
				img.Set(x, y, black)
			case solution:
				img.Set(x, y, red)
			default:
				img.Set(x, y, white)
			}
		}
	}

	return png.Encode(file, img)
}

func printMaze(m [][]byte) {
	for y := range m[0] {
		for x := range m {
			switch m[x][y] {
			case empty:
				fmt.Print("  ")
			case start:
				fmt.Print("S ")
			case end:
				fmt.Print("E ")
			case path:
				fmt.Print(". ")
			case solution:
				fmt.Print("+ ")
			}
		}
		fmt.Print("\n")
	}
}

func copyMaze(m [][]byte) [][]byte {
	c := make([][]byte, len(m))
	for x := range m {
		c[x] = make([]byte, len(m[x]))
		copy(c[x], m[x])
	}

	return c
}

func nextTo(m [][]byte, p point) (paths []point, ends []point) {
	for _, np := range p.neighbors() {
		if np.y >= 0 && np.y < len(m) && np.x >= 0 && np.x < len(m[np.y]) {
			switch m[np.x][np.y] {
			case path:
				paths = append(paths, np)
			case end:
				ends = append(ends, np)
			}
		}
	}

	return paths, ends
}

func solveMaze(m [][]byte) (s [][]byte, l uint, err error) {
	type searchPath struct {
		head   point
		m      [][]byte
		length uint
	}

	// Start suchen
	startPoint := point{0, 0}
	countStarts := 0
	for x := range m {
		for y := range m[x] {
			if m[x][y] == start {
				startPoint.x, startPoint.y = x, y
				countStarts++
			}
		}
	}

	if countStarts != 1 {
		err := fmt.Errorf("Exakt 1 Startpunkt erlaubt. %d gefunden", countStarts)
		return nil, 0, err
	}

	// Suchpfade initialisieren
	searchPaths := make([]searchPath, 1)
	searchPaths[0] = searchPath{
		head:   startPoint,
		m:      copyMaze(m),
		length: uint(0),
	}

	minSeenLength := make([][]uint, len(m))
	for x := range minSeenLength {
		minSeenLength[x] = make([]uint, len(m[x]))
		for y := range minSeenLength[x] {
			minSeenLength[x][y] = ^uint(0)
		}
	}

	// Suche mit mehreren Köpfen gleichzeitig
	for {
		var newPaths []searchPath

		for _, sp := range searchPaths {
			sp.m[sp.head.x][sp.head.y] = solution

			paths, ends := nextTo(sp.m, sp.head)

			if len(ends) > 0 {
				// Endpunkt erreicht
				sp.m[startPoint.x][startPoint.y] = start // Start wieder eintragen
				return sp.m, sp.length, nil
			}

			for _, p := range paths {
				newPath := searchPath{
					head:   p,
					m:      copyMaze(sp.m),
					length: sp.length + 1,
				}

				if newPath.length < minSeenLength[p.x][p.y] {
					// kürzester bekannter Pfad zu diesem Punkt
					minSeenLength[p.x][p.y] = newPath.length
					newPaths = append(newPaths, newPath)
				}
			}
		}

		if len(newPaths) == 0 {
			err := fmt.Errorf("Es existiert keine Lösung.")
			return nil, 0, err
		}

		searchPaths = newPaths
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Aufruf: amazeing /path/to/maze.png")
		os.Exit(1)
	}

	path := args[0]

	m, err := readMaze(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s, l, err := solveMaze(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Lösung gefunden. Pfadlänge:", l)

	solvedPath := strings.TrimSuffix(path, ".png") + "_solved.png"
	err = writeMaze(solvedPath, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Lösung geschrieben:", solvedPath)
}
