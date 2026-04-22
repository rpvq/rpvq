package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

    /w

// ANSI codes
const (
	reset   = "\033[0m"
	bold    = "\033[1m"
	dim     = "\033[2m"
	gray0   = "\033[90m"
	gray1   = "\033[37m"
	white   = "\033[97m"
	clrscr  = "\033[2J\033[H"
	home    = "\033[H"
	hidecur = "\033[?25l"
	showcur = "\033[?25h"
	
	// a
)

// R 

// Character art generated directly from the uploaded image
// 0 = dark areas (hair, eyes, outlines)
// 1 = mid-tone areas (face, body)
// space = background
var art = []string{
	`          1  10000000001                `,
	`        11  10010000000011              `,
	`       111010000001000000111            `,
	`      1000000000000000000001            `,
	`     100000000000000000000011           `,
	`    10000001000100000000000011          `,
	`    1000001 000000000000010001          `,
	`    000001110000000000000100011         `,
	`   100001  10010000000000000001         `,
	`   10000001 11 0000000000000001         `,
	`   00001100  1 0011000000000000111      `,
	`  10000111     11 100000000000011       `,
	`  11001        10010000000000001        `,
	`   1001         000000000000000         `,
	`   100          100000000000000         `,
	`   001           1 111000000001         `,
	`   000             100000000001         `,
	`  11001            000000000011         `,
	`   1100           10001100001011        `,
	`    1101          100110001011          `,
	`      001        100000000 10111        `,
	`      10011000000000000001 01           `,
	`        011000001100000001             `,
	`           1001  101111001             `,
	`            
}

    /u
    
type entry struct {
	k string
	v string
}

    //y

var profile = []entry{
	{"name",    "noriol"},
	{"status",  "bored"},
	{"craving", "hamburger"},
	{"note",    "literally nothing to say"},
	{"reason",  "did not want to waste the account"},
	{"file",    "simple redmi file"},
	{"offer",   "eat a hamburger with me"},
	{"github",  "github.com/noriol"},
}

func ms(n int) { time.Sleep(time.Duration(n) * time.Millisecond) }

func typeWrite(s string, delay int) {
	for _, c := range s {
		fmt.Printf("%c", c)
		ms(delay)
	}
}

// matrixRain — 0/1 digits appear bright and fade to dark before vanishing
func matrixRain(rows, cols int, dur time.Duration) {
	fmt.Print(hidecur + clrscr)

	grid := make([][]byte, rows)
	bright := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]byte, cols)
		bright[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	deadline := time.Now().Add(dur)
	for time.Now().Before(deadline) {
		for k := 0; k < 8; k++ {
			r := rand.Intn(rows)
			c := rand.Intn(cols)
			grid[r][c] = byte('0' + rand.Intn(2))
			bright[r][c] = 6
		}
		for r := range grid {
			for c := range grid[r] {
				if bright[r][c] > 0 {
					bright[r][c]--
					if bright[r][c] == 0 {
						grid[r][c] = ' '
					}
				}
			}
		}
		fmt.Print(home)
		for r, row := range grid {
			for c, ch := range row {
				b := bright[r][c]
				switch {
				case b >= 5:
					fmt.Print(white + bold + string(ch) + reset)
				case b >= 3:
					fmt.Print(gray1 + string(ch) + reset)
				case b >= 1:
					fmt.Print(gray0 + dim + string(ch) + reset)
				default:
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		ms(50)
	}
}

func scanLine(width, delay int) {
	fmt.Print(gray0 + bold)
	for i := 0; i < width; i++ {
		if i%4 < 2 {
			fmt.Print("-")
		} else {
			fmt.Print(" ")
		}
		if delay > 0 {
			ms(delay)
		}
	}
	fmt.Print(reset)
	fmt.Println()
}

func pad(s string, w int) string {
	if len(s) >= w {
		return s
	}
	return s + strings.Repeat(" ", w-len(s))
}

func drawProfile() {
	fmt.Print(clrscr + hidecur)

	artW := 0
	for _, l := range art {
		if len(l) > artW {
			artW = len(l)
		}
	}

	const infoW = 50
	const sep = "  |  "
	const indent = " "

	totalW := len(indent) + artW + len(sep) + infoW

	scanLine(totalW, 3)
	ms(50)

	leftH  := indent + pad("  character", artW)
	rightH := pad("  profile", infoW)
	fmt.Print(gray1 + bold)
	typeWrite(leftH+sep+rightH+"\n", 5)
	fmt.Print(reset)

	fmt.Print(gray0 + dim)
	fmt.Println(indent + strings.Repeat("-", artW) + sep + strings.Repeat("-", infoW))
	fmt.Print(reset)

	maxLines := len(art)
	if len(profile)+3 > maxLines {
		maxLines = len(profile) + 3
	}

	for i := 0; i < maxLines; i++ {
		if i < len(art) {
			fmt.Print(gray0 + indent + art[i] + reset)
		} else {
			fmt.Print(strings.Repeat(" ", len(indent)+artW))
		}

		fmt.Print(sep)
		switch {
		case i == 0:
			fmt.Print(gray1 + bold + pad("  noriol", infoW) + reset)
		case i == 1:
			fmt.Print(gray0 + dim + pad("  "+strings.Repeat("-", infoW-4), infoW) + reset)
		case i-2 < len(profile):
			p := profile[i-2]
			line := fmt.Sprintf("  %-10s  %s", p.k, p.v)
			fmt.Print(gray1 + pad(line, infoW) + reset)
		}

		fmt.Println()
		ms(40)
	}

	fmt.Print(gray0 + dim)
	fmt.Println(indent + strings.Repeat("-", artW) + sep + strings.Repeat("-", infoW))
	fmt.Print(reset)
	scanLine(totalW, 2)
	ms(80)

	fmt.Println()
	fmt.Print(gray0 + bold)
	typeWrite(indent+" noriol", 55)
	fmt.Print(reset + gray0 + dim)
	typeWrite("  --  github.com/noriol  --  bored  --  want a hamburger\n", 22)
	fmt.Print(reset + "\n")
	fmt.Print(showcur)
}

    / a

func main() {
	rand.Seed(time.Now().UnixNano())
	matrixRain(23, 82, 2400*time.Millisecond)
	drawProfile()
}

// n