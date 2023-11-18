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

fn fizz_buzz(n: i32) -> Vec<String> {
    let mut result = Vec::with_capacity(n as usize);
    for i in 1..=n {
        let v = match i {
            _ if i % 15 == 0 => "FizzBuzz".to_string(),
            _ if i % 3 == 0 => "Fizz".to_string(),
            _ if i % 5 == 0 => "Buzz".to_string(),
            _ => i.to_string()
        };
        result.push(v);
    }
    result
}

fn number_of_steps(mut num: i32) -> i32 {
    let mut steps = 0;
    while num != 0 {
        steps+= 1;
        if num % 2 == 0 {
            num /= 2;
        } else {
            num -= 1;
        }
    }
    steps
}

#[cfg(test)]
mod tests {
    use crate::{fizz_buzz, maximum_wealth, number_of_steps, running_sum};

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

    #[test]
    fn test_fizz_buzz() {
        assert_eq!(fizz_buzz(3), vec!["1", "2", "Fizz"]);
        assert_eq!(fizz_buzz(5), vec!["1", "2", "Fizz", "4", "Buzz"]);
        assert_eq!(fizz_buzz(15), vec!["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"]);
    }

    #[test]
    fn test_number_of_steps() {
        assert_eq!(number_of_steps(14), 6);
        assert_eq!(number_of_steps(8), 4);
        assert_eq!(number_of_steps(123), 12);
    }
}
