// mask
package qrgen

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
