package main

import (
	//textcolor "color"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	_ "io"
	"strings"
)

//Color used to display pattern
const (
	PatternColor = "\033[1;34m%s\033[0m"
)

const gopher = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`

func readerPNG() io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(gopher))

}

func main() {
	img, _ := png.Decode(readerPNG())
	levels := []string{" ", ".", ",", "-", "~", ":", ";", "=", "!", "*", "#", "$", "@"}
	println(len(levels))
	printASCIImage(img, levels)
	var pixels = getPixels(img, levels)
	// symbols on top of the left eye
	//example of horizontal symbol
	//eyeSymbols := []string{";", "=", "="}
	//example of vertical symbol
	eyeSymbols := []string{"~", ";", "~"}
	pattern := Pattern{pattern: eyeSymbols, horizontal: false}
	patternPositionList := pattern.findPatternLocation(pixels)
	for patternPosition := range patternPositionList {
		if patternPosition == -1 {
			println("no symbol here")
		}
	}
	printASCIImageColorizedPattern(img, levels, patternPositionList, PatternColor)
}

// Pattern is what will help find where the eyes are situated
type Pattern struct {
	pattern    []string
	horizontal bool
}

func (pattern Pattern) findPatternLocation(img [][]string) [][]int {
	lenY := len(img)
	lenX := len(img[0])
	position := 0
	patternPositionList := [][]int{}
	if pattern.horizontal == true {
		println("horizontal pattern")
		for i := 0; i < lenY; i++ {
			for j := 0; j < lenX; j++ {
				if img[i][j] == pattern.pattern[position] {
					position++
					if position == len(pattern.pattern) {
						for nb := len(pattern.pattern) - 1; nb >= 0; nb-- {
							if j-nb > 0 {
								newPosition := []int{i, j - nb}
								patternPositionList = append(patternPositionList, newPosition)
							}
						}
						return patternPositionList
					}
				} else {
					position = 0
				}
			}
		}
	} else {
		println("vertical pattern")
		for i := 0; i < lenX; i++ {
			for j := 0; j < lenY; j++ {
				if img[j][i] == pattern.pattern[position] {
					position++
					if position == len(pattern.pattern) {
						for nb := len(pattern.pattern) - 1; nb >= 0; nb-- {
							if j-nb > 0 {
								newPosition := []int{j - nb, i}
								patternPositionList = append(patternPositionList, newPosition)
							}
						}
						return patternPositionList
					}
				} else {
					position = 0
				}
			}

		}
	}
	patternPositionList = append(patternPositionList, []int{-1, -1})
	return patternPositionList
}

func printASCIImageColorizedPattern(img image.Image, levels []string, pattern [][]int, symbolColor string) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// you can use img.At to get a specific pixel
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			//level := c.Y / 51
			level := c.Y / 20
			if level == 13 {
				level--
			}
			element := []int{y, x}
			_, found := Find(pattern, element)
			if !found {
				fmt.Print(levels[level])
			} else {
				fmt.Printf(symbolColor, levels[level])
			}
		}
		fmt.Print("\n")
	}
}

//Find allows you to check in a array of elements if an element is inside
func Find(slice [][]int, val []int) (int, bool) {
	for i, item := range slice {
		if (item[0] == val[0]) && (item[1] == val[1]) {
			return i, true
		}
	}
	return -1, false
}

func printASCIImage(img image.Image, levels []string) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// you can use img.At to get a specific pixel
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			//level := c.Y / 51
			level := c.Y / 20
			if level == 13 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}

//transform the image into a 2d array of strings containing the symbols
func getPixels(img image.Image, levels []string) [][]string {
	bounds := img.Bounds()
	maxY := bounds.Max.Y
	maxX := bounds.Max.X
	charImage := make([][]string, maxY)
	for i := range charImage {
		charImage[i] = make([]string, maxX)
	}
	//charImage := [][]string{{}}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// you can use img.At to get a specific pixel
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			//level := c.Y / 51
			level := c.Y / 20
			if level == 13 {
				level--
			}
			charImage[y][x] = (levels[level])
		}
	}
	return charImage
}
