fn main() {
    println!("Hello, world!");
}

fn running_sum(mut nums: Vec<i32>) -> Vec<i32> {
    for i in 1..nums.len() {
        nums[i] += nums[i - 1];
    }
    return nums;
}

#[cfg(test)]
mod tests {
    use crate::running_sum;

    #[test]
    fn test_running_sum() {
        assert_eq!(running_sum(vec![1, 2, 3, 4]), vec![1, 3, 6, 10]);
        assert_eq!(running_sum(vec![1, 1, 1, 1, 1]), vec![1, 2, 3, 4, 5]);
        assert_eq!(running_sum(vec![3, 1, 2, 10, 1]), vec![3, 4, 6, 16, 17]);
    }
}
