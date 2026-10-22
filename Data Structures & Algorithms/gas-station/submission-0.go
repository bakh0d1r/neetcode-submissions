func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    sumGas , sumCost := 0,0
    for i:=range n {
        sumGas+=gas[i]
        sumCost+=cost[i]
    }
    if sumGas < sumCost {
        return -1
    }
    ans,currGas:=0,0
    for i:=range n{
        currGas+=gas[i]-cost[i]
        if currGas < 0 {
            currGas = 0
            ans=i+1
        }
    }
    return ans
}