func shipWithinDays(weights []int, days int) int {
	mx:=-1
	t:=0
	for i:=range weights{
		mx=max(mx,weights[i])
		t+=weights[i]
	}
	l:=mx
	h:=t
	for l <= h {
		m:= l+ (h-l)/2
		d:=1
		c:=0
		for i:=range weights{
			if c + weights[i] > m{
				d++
				c=0
			}
			c+=weights[i]
		}
		if d>days {
			l=m+1
		}else {
			h=m-1
		}
	}
	return l
}
