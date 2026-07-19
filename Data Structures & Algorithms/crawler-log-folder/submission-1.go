func minOperations(logs []string) int {
	s:=[]string{}
	for i:=range logs{
		if logs[i] == "../" {
			if  len(s) > 0  {
				s=s[:len(s)-1]
			}
			continue
		}
		if logs[i] == "./"{
			continue
		}
		s=append(s,logs[i])
	}
	return len(s)
}
