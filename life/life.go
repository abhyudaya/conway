package life

// Slices are pass by reference. Pass a current and next grid of equal sizes
// and NextLife fills 'next' grid with values based on Game of Life rules.
func NextLife(current, next [][]uint8)[][]uint8 {
  for i := range current {
    for j := range current[i] {
      count := getCount(current, i, j)
      switch {
         case current[i][j] == 1 && count < 2:
		next[i][j] = 0
         case current[i][j] == 1 && (count == 2 || count == 3):
		next[i][j] = 1
	 case current[i][j] == 1 && count > 3:
		next[i][j] = 0
	 case current[i][j] == 0 && count == 3:
		next[i][j] = 1
      }
    }
  }
  return next
}

func getCount(grid [][]uint8, i, j int) int {
  x := len(grid)
  y := len(grid[i])
  sum := 0

  for m := i - 1; m <= i + 1; m++ {
	for n := j - 1; n <= j + 1; n++ {
		switch {
		  case m < 0 || m >= x:
			break
		  case n < 0 || n >= y:
			break
		  default:
			sum += int(grid[m][n])
		}
	}
  }	

  return sum - int(grid[i][j])
}
