fn main() {
    println!("Hello, world!");
}

fn running_sum(mut nums: Vec<i32>) -> Vec<i32> {
    for i in 1..nums.len() {
        nums[i] += nums[i - 1];
    }
    nums
}

fn maximum_wealth(accounts: Vec<Vec<i32>>) -> i32 {
    let mut max = 0;
    for banks in accounts {
        let sum = banks.iter().sum();
        if sum > max {
            max = sum;
        }
    }
    max
}

#[cfg(test)]
mod tests {
    use crate::{maximum_wealth, running_sum};

    #[test]
    fn test_running_sum() {
        assert_eq!(running_sum(vec![1, 2, 3, 4]), vec![1, 3, 6, 10]);
        assert_eq!(running_sum(vec![1, 1, 1, 1, 1]), vec![1, 2, 3, 4, 5]);
        assert_eq!(running_sum(vec![3, 1, 2, 10, 1]), vec![3, 4, 6, 16, 17]);
    }

    #[test]
    fn test_maximum_wealth() {
        assert_eq!(maximum_wealth(vec![vec![1, 2, 3], vec![3, 2, 1]]), 6);
        assert_eq!(maximum_wealth(vec![vec![1, 5], vec![7, 3], vec![3, 5]]), 10);
        assert_eq!(maximum_wealth(vec![vec![2, 8, 7], vec![7, 1, 3], vec![1, 9, 5]]), 17);
    }
}
