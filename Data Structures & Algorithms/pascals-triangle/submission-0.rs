impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let num_rows = num_rows as usize;
        let mut res: Vec<Vec<i32>> = Vec::new();
        for i in 0..num_rows {
            let mut row = vec![1;i+1];
            for j in 1..i{
                row[j] = res[i-1][j-1] + res[i-1][j] ;
           }
           res.push(row);
        }
        res
    }
}
