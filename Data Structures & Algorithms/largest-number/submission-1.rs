impl Solution {
    pub fn largest_number(nums: Vec<i32>) -> String {
        let mut arr: Vec<String> = nums.iter().map(|n|n.to_string()).collect();
        arr.sort_by(|a , b| {
            let ab = format!("{}{}",b,a);
            let ba = format!("{}{}",a,b);
            ab.cmp(&ba)
        });
        let res:String = arr.join("");
        if res.starts_with('0') {"0".to_string()} else { res}
    }
}
