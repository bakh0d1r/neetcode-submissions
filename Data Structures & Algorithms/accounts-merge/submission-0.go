func accountsMerge(accounts [][]string) [][]string {
	arr := [][]string{}
	emailToName := map[string]string{}
	for i := range accounts {
		for j := range accounts[i] {
			if j > 0 {
				emailToName[accounts[i][j]] = accounts[i][0]
			}
		}
	}
	graph := map[string][]string{}
	for i := range accounts {
		emails := []string{}
		for j := 1; j < len(accounts[i]); j++ {
			emails = append(emails, accounts[i][j])
		}

		for j := 1; j < len(emails); j++ {
			graph[emails[0]] = append(graph[emails[0]], emails[j])
			graph[emails[j]] = append(graph[emails[j]], emails[0])
		}
	}
	visited := map[string]bool{}
	merged := []string{}
    var dfs func(string)

	dfs = func(email string) {
		visited[email] = true
		merged = append(merged, email)
		for _, next := range graph[email] {
			if !visited[next] {
				dfs(next)
			}
		}
	}
	for email := range emailToName {
		if !visited[email] {
            merged = []string{}

			dfs(email)
			sort.Strings(merged)
			account := []string{emailToName[merged[0]]}
			account = append(account, merged...)
			arr = append(arr, account)
		}
	}

	return arr
}