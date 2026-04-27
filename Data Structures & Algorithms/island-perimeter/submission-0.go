func islandPerimeter(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	peri := 0

	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if grid[i][j] == 1{
				if j-1 < 0 || grid[i][j-1] == 0{
					peri++
				}
				if i-1 < 0 || grid[i-1][j] == 0{
					peri++
				}
				if i+1 == n || grid[i+1][j] == 0{
					peri++
				}
				if j+1 == m || grid[i][j+1] == 0{
					peri++
				} 
			}
		}
	}

	return peri
}
