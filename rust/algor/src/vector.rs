fn running_sum(mut nums: Vec<i32>) -> Vec<i32> {
    for i in 1..nums.len() {
        nums[i] += nums[i - 1];
    }
    nums
}
#[test]
fn test_running_sum() {
    assert_eq!(running_sum(vec![1, 2, 3, 4]), vec![1, 3, 6, 10]);
    assert_eq!(running_sum(vec![1, 1, 1, 1, 1]), vec![1, 2, 3, 4, 5]);
    assert_eq!(running_sum(vec![3, 1, 2, 10, 1]), vec![3, 4, 6, 16, 17]);
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
#[test]
fn test_maximum_wealth() {
    assert_eq!(maximum_wealth(vec![vec![1, 2, 3], vec![3, 2, 1]]), 6);
    assert_eq!(maximum_wealth(vec![vec![1, 5], vec![7, 3], vec![3, 5]]), 10);
    assert_eq!(
        maximum_wealth(vec![vec![2, 8, 7], vec![7, 1, 3], vec![1, 9, 5]]),
        17
    );
}


fn fizz_buzz(n: i32) -> Vec<String> {
    let mut result = Vec::with_capacity(n as usize);
    for i in 1..=n {
        let v = match i {
            _ if i % 15 == 0 => "FizzBuzz".to_string(),
            _ if i % 3 == 0 => "Fizz".to_string(),
            _ if i % 5 == 0 => "Buzz".to_string(),
            _ => i.to_string(),
        };
        result.push(v);
    }
    result
}
#[test]
fn test_fizz_buzz() {
    assert_eq!(fizz_buzz(3), vec!["1", "2", "Fizz"]);
    assert_eq!(fizz_buzz(5), vec!["1", "2", "Fizz", "4", "Buzz"]);
    assert_eq!(
        fizz_buzz(15),
        vec![
            "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13",
            "14", "FizzBuzz"
        ]
    );
}


fn duplicate_zeros(arr: &mut Vec<i32>) {
    let mut i = 0;
    while i < arr.len() {
        if arr[i] == 0 {
            i += 1;
            arr.insert(i, 0);
            arr.pop();
        }
        i += 1;
    }
}
#[test]
fn test_duplicate_zeros() {
    let mut v1 = vec![1, 0, 2, 3, 0, 4, 5, 0];
    duplicate_zeros(&mut v1);
    assert_eq!(v1, vec![1, 0, 0, 2, 3, 0, 0, 4]);

    let mut v2 = vec![1, 2, 3];
    duplicate_zeros(&mut v2);
    assert_eq!(v2, vec![1, 2, 3]);
}


fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
    let (mut max, mut cur) = (0, 0);
    for v in nums {
        match v {
            0 => {
                if cur > max {
                    max = cur;
                }
                cur = 0;
            }

            1 => cur += 1,
            _ => {}
        }
    }
    if cur > max {
        max = cur;
    }
    max
}
#[test]
fn test_find_max_consecutive_ones() {
    assert_eq!(find_max_consecutive_ones(vec![1, 1, 0, 1, 1, 1]), 3);
    assert_eq!(find_max_consecutive_ones(vec![1, 0, 1, 1, 0, 1]), 2);
}


fn check_if_exist(arr: Vec<i32>) -> bool {
    arr.iter()
        .map(|&v| v * 2)
        .enumerate()
        .any(|(i, v)| arr.iter().enumerate().any(|(j, &v2)| i != j && v == v2))
}
#[test]
fn test_check_if_exist() {
    assert_eq!(check_if_exist(vec![10, 2, 5, 3]), true);
    assert_eq!(check_if_exist(vec![3, 1, 7, 11]), false);
}


fn valid_mountain_array(arr: Vec<i32>) -> bool {
    if arr.len() < 3 || arr[0] >= arr[1] {
        return false;
    }
    let mut up = true;
    for i in 1..arr.len() {
        let curr = arr[i];
        let prev = arr[i - 1];
        match curr {
            _ if curr == prev => return false,
            _ if up && curr < prev => up = false,
            _ if !up && curr > prev => return false,
            _ => {}
        }
    }
    !up
}
#[test]
fn test_valid_mountain_array() {
    assert_eq!(valid_mountain_array(vec![2, 1]), false);
    assert_eq!(valid_mountain_array(vec![3, 5, 5]), false);
    assert_eq!(valid_mountain_array(vec![0, 3, 2, 1]), true);
    assert_eq!(
        valid_mountain_array(vec![0, 1, 2, 3, 4, 5, 6, 7, 8, 9]),
        false
    );
    assert_eq!(
        valid_mountain_array(vec![9, 8, 7, 6, 5, 4, 3, 2, 1, 0]),
        false
    );
}


fn replace_elements(mut arr: Vec<i32>) -> Vec<i32> {
    let i_last = arr.len() - 1;
    let mut max = arr[i_last];
    arr[i_last] = -1;
    for i in (0..i_last).rev() {
        let curr = arr[i];
        arr[i] = max;
        if curr > max {
            max = curr;
        }
    }
    arr
}
#[test]
fn test_replace_elements() {
    assert_eq!(
        replace_elements(vec![17, 18, 5, 4, 6, 1]),
        vec![18, 6, 6, 6, 1, -1]
    );
}


fn move_zeroes(nums: &mut Vec<i32>) {
    let mut i_write = 0;

    for i_read in 0..nums.len() {
        if nums[i_read] != 0 {
            nums[i_write] = nums[i_read];
            i_write += 1;
        }
    }

    for i in i_write..nums.len() {
        nums[i] = 0;
    }
}
#[test]
fn test_move_zeros() {
    let mut vec1 = vec![0, 1, 0, 3, 12];
    move_zeroes(&mut vec1);
    assert_eq!(vec1, vec![1, 3, 12, 0, 0]);
}