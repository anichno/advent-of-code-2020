package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/set"
)

type matchDirection int

const (
	normalOrientation = iota
	flipHorizontalOrientation
	flipVerticalOrientation
	matchDirectionRight  matchDirection = 0
	matchDirectionBottom matchDirection = 1
)

func flipHorizontal(input [][]bool) [][]bool {
	flippedImage := make([][]bool, len(input))

	for i, row := range input {
		flippedImage[len(flippedImage)-1-i] = row
	}

	return flippedImage
}

func flipVertical(input [][]bool) [][]bool {
	flippedImage := make([][]bool, len(input))

	for y, row := range input {
		newRow := make([]bool, len(row))
		for x, col := range row {
			newRow[len(row)-1-x] = col
		}

		flippedImage[y] = newRow
	}

	return flippedImage
}

func rotateRight(input [][]bool) [][]bool {
	rotatedImage := make([][]bool, len(input))
	for i := 0; i < len(input); i++ {
		rotatedImage[i] = make([]bool, len(input[0]))
	}

	for i, row := range input {
		for j, col := range row {
			rotatedImage[j][len(row)-1-i] = col
		}
	}

	return rotatedImage
}

func matchRight(leftImage, rightImage [][]bool) bool {
	for y := 0; y < len(leftImage); y++ {
		if leftImage[y][len(leftImage[0])-1] != rightImage[y][0] {
			return false
		}
	}

	return true
}

func matchBottom(topImage, bottomImage [][]bool) bool {
	for x := 0; x < len(topImage[0]); x++ {
		if topImage[len(topImage)-1][x] != bottomImage[0][x] {
			return false
		}
	}

	return true
}

func matchImages(image1, image2 [][]bool, matchDir matchDirection) [][]bool {
	for _, orientation := range []int{normalOrientation, flipHorizontalOrientation, flipVerticalOrientation} {
		flippedImage := image2
		if orientation == flipHorizontalOrientation {
			flippedImage = flipHorizontal(image2)
		} else if orientation == flipVerticalOrientation {
			flippedImage = flipVertical(image2)
		}

		for rotate := 0; rotate < 4; rotate++ {
			// check if match
			var match bool

			if matchDir == matchDirectionRight {
				match = matchRight(image1, flippedImage)
			} else if matchDir == matchDirectionBottom {
				match = matchBottom(image1, flippedImage)
			} else {
				log.Fatal("Invalid matchDir")
			}

			if match {
				return flippedImage
			}

			flippedImage = rotateRight(flippedImage)
		}
	}

	return nil
}

func highlightSeaMonsters(image [][]bool) [][]bool {
	monsterPatternString := `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

	var monsterPattern [][]bool
	for _, line := range strings.Split(monsterPatternString, "\n") {
		var row []bool
		for _, char := range line {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		monsterPattern = append(monsterPattern, row)
	}

	var monsterHighlight [][]bool
	for y := 0; y < len(image); y++ {
		row := make([]bool, len(image[y]))
		monsterHighlight = append(monsterHighlight, row)
	}

	monsterFound := false
	for y := 0; y < len(image)-3; y++ {
		for scanX := 0; scanX < len(image[y])-len(monsterPattern[0]); scanX++ {
			monster := true
			for i := 0; i < len(monsterPattern[0]); i++ {
				if monsterPattern[0][i] && !image[y][scanX+i] {
					monster = false
					break
				}

				if monsterPattern[1][i] && !image[y+1][scanX+i] {
					monster = false
					break
				}

				if monsterPattern[2][i] && !image[y+2][scanX+i] {
					monster = false
					break
				}
			}

			if monster {
				monsterFound = true
				for i := 0; i < len(monsterPattern[0]); i++ {
					monsterHighlight[y][scanX+i] = monsterPattern[0][i]
					monsterHighlight[y+1][scanX+i] = monsterPattern[1][i]
					monsterHighlight[y+2][scanX+i] = monsterPattern[2][i]
				}
			}
		}
	}

	if monsterFound {
		return monsterHighlight
	} else {
		return nil
	}
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	images := make(map[int][][]bool)
	var curImage [][]bool
	var curImageID int
	for scanner.Scan() {
		text := scanner.Text()
		if curImage == nil {
			curImageID, _ = strconv.Atoi(strings.TrimSuffix(strings.Split(text, " ")[1], ":"))
			curImage = make([][]bool, 0)
		} else if len(text) == 0 {
			images[curImageID] = curImage
			curImage = nil
		} else {
			newRow := make([]bool, len(text))
			for i, char := range text {
				if char == '#' {
					newRow[i] = true
				}
			}
			curImage = append(curImage, newRow)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tileMatches := make(map[int]*set.Set)
	for idLeft, image1 := range images {
		matches := set.New()
		for _, orientation := range []int{normalOrientation, flipHorizontalOrientation, flipVerticalOrientation} {
			image := image1
			if orientation == flipHorizontalOrientation {
				image = flipHorizontal(image1)
			} else if orientation == flipVerticalOrientation {
				image = flipVertical(image1)
			}

			for i := 0; i < 4; i++ {
				for idRight, image2 := range images {
					if idLeft == idRight {
						continue
					}

					testImage := image2
					for j := 0; j < 4; j++ {
						image2Bottom := testImage[len(testImage)-1]
						imageTop := image[0]

						match := true
						for idx := 0; idx < len(imageTop); idx++ {
							if imageTop[idx] != image2Bottom[idx] {
								match = false
								break
							}
						}

						if match {
							matches.Insert(idRight)
							break
						}

						testImage = rotateRight(testImage)
					}

				}

				image = rotateRight(image)
			}
		}

		tileMatches[idLeft] = matches
	}

	// Part 1

	tot := 1
	for tileID, matches := range tileMatches {
		if matches.Len() == 2 {
			tot *= tileID
		}
	}

	fmt.Println("Part 1:", tot)

	// Part 2
	gridSize := int(math.Sqrt(float64(len(images))))

	var startCornerID int
	for tileID, matches := range tileMatches {
		if matches.Len() == 2 {
			startCornerID = tileID
			break
		}
	}

	cornerImage := images[startCornerID]

	imageOrientations := make(map[int][][]bool)

	var imageIDGrid [][]int
	for y := 0; y < gridSize; y++ {
		row := make([]int, gridSize)
		imageIDGrid = append(imageIDGrid, row)
	}

	for _, orientation := range []int{normalOrientation, flipHorizontalOrientation, flipVerticalOrientation} {
		flippedImage := cornerImage
		if orientation == flipHorizontalOrientation {
			flippedImage = flipHorizontal(cornerImage)
		} else if orientation == flipVerticalOrientation {
			flippedImage = flipVertical(cornerImage)
		}

		for cornerRotate := 0; cornerRotate < 4; cornerRotate++ {
			// find right side
			var rightID int
			var rightImage [][]bool

			tileMatches[startCornerID].Do(func(imageID interface{}) {
				if rightImage != nil {
					return
				}
				rightID = imageID.(int)
				rightImage = matchImages(flippedImage, images[rightID], matchDirectionRight)
			})

			// find bottom side
			var bottomID int
			var bottomImage [][]bool

			if rightImage != nil {
				tileMatches[startCornerID].Do(func(imageID interface{}) {
					if bottomImage != nil {
						return
					}
					bottomID = imageID.(int)
					if bottomID == rightID {
						return
					}
					bottomImage = matchImages(flippedImage, images[bottomID], matchDirectionBottom)
				})
			}

			if rightImage != nil && bottomImage != nil {
				imageOrientations[startCornerID] = flippedImage
				imageOrientations[rightID] = rightImage
				imageOrientations[bottomID] = bottomImage

				imageIDGrid[0][0] = startCornerID
				imageIDGrid[0][1] = rightID
				imageIDGrid[1][0] = bottomID

				break
			}

			flippedImage = rotateRight(flippedImage)
		}

		if imageIDGrid[0][0] != 0 {
			break
		}
	}

	for y := 0; y < gridSize; y++ {
		if y > 0 && y < gridSize-1 {
			// find next one below
			pTile := imageIDGrid[y][0]
			pTileImage := imageOrientations[pTile]
			var bottomID int
			var bottomImage [][]bool
			tileMatches[pTile].Do(func(imageID interface{}) {
				if bottomImage != nil {
					return
				}
				bottomID = imageID.(int)
				bottomImage = matchImages(pTileImage, images[bottomID], matchDirectionBottom)
			})

			imageIDGrid[y+1][0] = bottomID
			imageOrientations[bottomID] = bottomImage
		}
		for x := 0; x < gridSize-1; x++ {
			// already solved
			if y == 0 && x == 0 {
				continue
			}

			pTile := imageIDGrid[y][x]
			pTileImage := imageOrientations[pTile]
			var rightID int
			var rightImage [][]bool
			tileMatches[pTile].Do(func(imageID interface{}) {
				if rightImage != nil {
					return
				}
				rightID = imageID.(int)
				rightImage = matchImages(pTileImage, images[rightID], matchDirectionRight)
			})

			imageIDGrid[y][x+1] = rightID
			imageOrientations[rightID] = rightImage
		}
	}

	var fullImage [][]bool
	for y := 0; y < (len(cornerImage)-2)*gridSize; y++ {
		row := make([]bool, (len(cornerImage[0])-2)*gridSize)
		fullImage = append(fullImage, row)
	}

	for y, row := range imageIDGrid {
		for x, imageID := range row {
			image := imageOrientations[imageID]

			for j := 1; j < len(image)-1; j++ {
				for i := 1; i < len(image[j])-1; i++ {
					fullImage[y*(len(image)-2)+j-1][x*(len(image)-2)+i-1] = image[j][i]
				}
			}
		}
	}

	part2Answer := 0
	for _, orientation := range []int{normalOrientation, flipHorizontalOrientation, flipVerticalOrientation} {
		flippedImage := fullImage
		if orientation == flipHorizontalOrientation {
			flippedImage = flipHorizontal(fullImage)
		} else if orientation == flipVerticalOrientation {
			flippedImage = flipVertical(fullImage)
		}

		for i := 0; i < 4; i++ {
			monsterHighlight := highlightSeaMonsters(flippedImage)
			if monsterHighlight != nil {
				for y, row := range monsterHighlight {
					for x, col := range row {
						if !col && flippedImage[y][x] {
							part2Answer++
						}
					}
				}
				break
			}
			flippedImage = rotateRight(flippedImage)
		}

		if part2Answer > 0 {
			break
		}
	}

	fmt.Println("Part 2:", part2Answer)
}
