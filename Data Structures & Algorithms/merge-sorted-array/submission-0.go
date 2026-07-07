func merge(nums1 []int, m int, nums2 []int, n int) {
	l,r:=m-1,n-1
	p:=m+n-1 
	for l>=0 && r>=0 {
		if nums1[l] <= nums2[r] {
			nums1[p]=nums2[r]
			r--
		}else{
			nums1[p]=nums1[l]
			l--
		}
		p--
	}
	for ;l>=0 ; l--{
		nums1[p]=nums1[l]
		p--
	}
	for ;r>=0 ; r-- {
		nums1[p]=nums2[r]
		p--
	}
}
