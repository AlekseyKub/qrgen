// mask
package qrgen

import (
	"fmt"
)

// маскирование QR кода
func ApplyMask(arr [][]int, numMask int) {
	if numMask == 0 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if (x+y)%2 == 0 {
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}
	if numMask == 2 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if y%3 == 0 {
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}

	if numMask == 1 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if x%3 == 0 {
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}

	if numMask == 3 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if (x+y)%3 == 0 {
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}

	if numMask == 4 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if (y/3+x/2)%2 == 0 {
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}

	if numMask == 5 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if (x*y)%2+(x*y)%3 == 0 {
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}

	if numMask == 6 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if ((x*y)%2+(x*y)%3)%2 == 0 { //((X*Y) % 2 + (X*Y) % 3) % 2
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}

	if numMask == 7 {
		for x := 0; x < (len(arr)); x++ {
			for y := 0; y < (len(arr)); y++ {
				if (arr[x][y] == 0) || (arr[x][y] == 1) {
					if ((x*y)%3+(x+y)%2)%2 == 0 { //((X*Y) % 3 + (X+Y) % 2) % 2
						if arr[x][y] == 0 {
							arr[x][y] = 1
						} else {
							arr[x][y] = 0
						}
					}

				}
			}
		}
	}
}

// здесь буду перебирать все маски
func MaskSelection(arr [][]int) {
	fmt.Println("ghbdtn")
	fiveAddMore(arr)
}

func fiveAddMore(arr [][]int) {
	blaсk := 0 //счетки для белых
	white := 0 //счетчик для черных
	//last :=0	//цвет прошлого элимента
	ball := 0 //количество баллов
	fmt.Println("ghbdtn11")

	for i := 0; i < (len(arr)); i++ {
		for x := 0; x < (len(arr)); x++ {
			if (arr[i][x] == 1) || (arr[i][x] == 3) {
				blaсk++
				white = 0
			} else {
				white++
				blaсk = 0
			}
			if blaсk >= 5 {
				if x == len(arr)-1 {
					ball += blaсk - 2
				} else {
					if arr[i][x+1] == 0 {
						ball += blaсk - 2
					}
				}
			}

			if white >= 5 {
				if x == len(arr)-1 {
					ball += white - 2
				} else {
					if arr[i][x+1] == 0 {
						ball += white - 2
					}
				}
			}
			fmt.Sprintf("Количество баллов %d", ball)
		}
	}
	fmt.Sprintf("Количество баллов %d", ball)
	//return ball
}
