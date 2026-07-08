func minimumRecolors(blocks string, k int) int {
	wc:=0
	for i:=0;i<k;i++{
		if blocks[i] == 'W' {
			wc++
		}
	}
	m:=wc
	for i :=1; i<len(blocks)-k+1; i++ {
		if blocks[i+k-1] == 'W' {
			wc++
		}
		if  blocks[i-1] == 'W' {
			wc--
		}
		m=min(m,wc)
	}
	return m
}
