func findRedundantConnection(edges [][]int) []int {
    n:=len(edges)

    parent := make([]int,n+1)
    for i:=1;i<=n;i++{
        parent[i]=i
    }

    var find func(x int) int
    find = func(x int) int {
        if parent[x] == x {
            return x
        }
        parent[x] = find(parent[x])
        return parent[x]
    }
    var union func(a,b int) bool
     union = func (a,b int) bool {
        rootA:=find(a)
        rootB:=find(b)
        if rootA == rootB {
            return false
        }
        parent[rootB]=rootA
        return true
    }
    arr:=[]int{}
    for i:=range edges{
        u,v :=edges[i][0],edges[i][1]

        if !union(u,v) {
            arr=append(arr,u,v)
        }
    }
    return arr 
}

