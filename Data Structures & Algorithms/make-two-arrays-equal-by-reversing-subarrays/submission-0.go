func canBeEqual(target []int, arr []int) bool {
	    counts:=[1001]int{}; for _,x:=range target { counts[x]++ }; for _,x:=range arr { counts[x]-- }; for _,x:=range counts { if x!=0 { return false } }; return true

}


