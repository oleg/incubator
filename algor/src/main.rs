use std::collections::HashMap;

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
        steps += 1;
        if num % 2 == 0 {
            num /= 2;
        } else {
            num -= 1;
        }
    }
    steps
}

fn merge_alternately(word1: String, word2: String) -> String {
    let mut i1 = word1.chars();
    let mut i2 = word2.chars();
    let mut v = Vec::with_capacity(word1.len() + word2.len());
    loop {
        let n1 = i1.next();
        let n2 = i2.next();
        if let Some(c) = n1 {
            v.push(c);
        }
        if let Some(c) = n2 {
            v.push(c);
        }
        if n1.is_none() && n2.is_none() {
            break;
        }
    }
    v.into_iter().collect()
}


#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }
    fn from(vec: Vec<i32>) -> Self {
        let mut head = ListNode::new(vec[0]);
        let mut current = &mut head;
        for i in 1..vec.len() {
            let node = ListNode::new(vec[i]);
            current.next = Some(Box::new(node));
            current = current.next.as_mut().unwrap();
        }
        head
    }
}

fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let (mut step_op, mut half_step_op) = (head.clone(), head.clone());

    while let Some(step) = step_op {
        step_op = step.next;
        if let Some(step) = step_op {
            step_op = step.next;
            if let Some(half_step) = half_step_op {
                half_step_op = half_step.next;
            }
        }
    }

    half_step_op
}

fn can_construct_ransom_note(ransom_note: String, magazine: String) -> bool {
    let mut letters = HashMap::new();
    for c in magazine.chars() {
        letters.entry(c)
            .and_modify(|e| { *e += 1 })
            .or_insert(1);
    }

    for c in ransom_note.chars() {
        let v = letters.entry(c)
            .and_modify(|e| { *e -= 1 })
            .or_insert(-1);
        if *v < 0 {
            return false;
        };
    }

    true
}

fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
    let (mut max, mut cur) = (0, 0);
    for v in nums {
        match v {
            0 => {
                if cur > max { max = cur; }
                cur = 0;
            }

            1 => { cur += 1 }
            _ => {}
        }
    }
    if cur > max { max = cur; }
    max
}

fn sorted_squares(mut nums: Vec<i32>) -> Vec<i32> {
    for i in 0..nums.len() {
        nums[i] *= nums[i];
    }
    nums.sort();
    nums
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

fn merge_sorted(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
    let mut i1: usize = m as usize;
    let mut i2: usize = n as usize;
    let mut t: usize = (m + n) as usize;

    while i1 > 0 && i2 > 0 {
        if nums1[i1 - 1] >= nums2[i2 - 1] {
            nums1[t - 1] = nums1[i1 - 1];
            i1 -= 1;
        } else {
            nums1[t - 1] = nums2[i2 - 1];
            i2 -= 1;
        }
        t -= 1;
    }

    while i1 > 0 {
        nums1[t - 1] = nums1[i1 - 1];
        i1 -= 1;
        t -= 1;
    }

    while i2 > 0 {
        nums1[t - 1] = nums2[i2 - 1];
        i2 -= 1;
        t -= 1;
    }
}


fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut i = 0;
    let mut equal = 0;

    while i < nums.len() {
        if nums[i] != val {
            nums[i - equal] = nums[i];
        } else {
            equal += 1;
        }
        i += 1;
    }

    (nums.len() - equal) as i32
}

fn remove_element2(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut i_write = 0;

    for i_read in 0..nums.len() {
        if nums[i_read] != val {
            nums.swap(i_write, i_read);
            i_write += 1;
        }
    }

    i_write as i32
}

fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    let mut dup = 0;

    let mut i = 1;
    let mut prev = nums[0];

    while i < nums.len() {
        let last = nums[i];
        if last != prev {
            nums[i - dup] = last;
        } else {
            dup += 1;
        }
        prev = last;
        i += 1;
    }

    (nums.len() - dup) as i32
}

fn remove_duplicates2(nums: &mut Vec<i32>) -> i32 {
    let mut i_write = 1;

    for i_read in 1..nums.len() {
        if nums[i_read] != nums[i_read - 1] {
            nums[i_write] = nums[i_read];
            i_write += 1;
        }
    }

    i_write as i32
}

fn check_if_exist(arr: Vec<i32>) -> bool {
    arr.iter()
        .map(|&v| v * 2)
        .enumerate()
        .any(|(i, v)| {
            arr.iter()
                .enumerate()
                .any(|(j, &v2)| {
                    i != j && v == v2
                })
        })
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

fn sort_array_by_parity(mut nums: Vec<i32>) -> Vec<i32> {
    let mut i_write = 0;

    for i_read in 0..nums.len() {
        if nums[i_read] % 2 == 0 {
            nums.swap(i_write, i_read);
            i_write += 1;
        }
    }

    nums
}

#[cfg(test)]
mod tests {
    use crate::{
        can_construct_ransom_note, check_if_exist, duplicate_zeros, find_max_consecutive_ones,
        fizz_buzz, ListNode, maximum_wealth, merge_alternately, merge_sorted, middle_node,
        move_zeroes, number_of_steps, remove_duplicates, remove_duplicates2, remove_element, remove_element2,
        replace_elements, running_sum, sort_array_by_parity, sorted_squares, valid_mountain_array,
    };

    #[test]
    fn test_sort_array_by_parity() {
        assert_eq!(sort_array_by_parity(vec![3, 1, 2, 4]), vec![2, 4, 3, 1]);
        assert_eq!(sort_array_by_parity(vec![1, 2, 3, 4]), vec![2, 4, 3, 1]);
        assert_eq!(sort_array_by_parity(vec![0]), vec![0]);
        assert_eq!(sort_array_by_parity(vec![1]), vec![1]);
        assert_eq!(sort_array_by_parity(vec![2, 4]), vec![2, 4]);
        assert_eq!(sort_array_by_parity(vec![1, 3]), vec![1, 3]);
    }

    #[test]
    fn test_move_zeros() {
        let mut vec1 = vec![0, 1, 0, 3, 12];
        move_zeroes(&mut vec1);
        assert_eq!(vec1, vec![1, 3, 12, 0, 0]);
    }

    #[test]
    fn test_replace_elements() {
        assert_eq!(replace_elements(vec![17, 18, 5, 4, 6, 1]), vec![18, 6, 6, 6, 1, -1]);
    }

    #[test]
    fn test_valid_mountain_array() {
        assert_eq!(valid_mountain_array(vec![2, 1]), false);
        assert_eq!(valid_mountain_array(vec![3, 5, 5]), false);
        assert_eq!(valid_mountain_array(vec![0, 3, 2, 1]), true);
        assert_eq!(valid_mountain_array(vec![0, 1, 2, 3, 4, 5, 6, 7, 8, 9]), false);
        assert_eq!(valid_mountain_array(vec![9, 8, 7, 6, 5, 4, 3, 2, 1, 0]), false);
    }

    #[test]
    fn test_check_if_exist() {
        assert_eq!(check_if_exist(vec![10, 2, 5, 3]), true);
        assert_eq!(check_if_exist(vec![3, 1, 7, 11]), false);
    }

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

    #[test]
    fn test_middle_node() {
        assert_eq!(middle_node(Some(Box::new(ListNode::from(vec![1, 2, 3, 4, 5])))),
                   Some(Box::new(ListNode::from(vec![3, 4, 5]))));

        assert_eq!(middle_node(Some(Box::new(ListNode::from(vec![1, 2, 3, 4, 5, 6])))),
                   Some(Box::new(ListNode::from(vec![4, 5, 6]))));
    }

    #[test]
    fn test_can_construct_ransom_note() {
        assert_eq!(can_construct_ransom_note("a".to_string(), "b".to_string()), false);
        assert_eq!(can_construct_ransom_note("aa".to_string(), "ab".to_string()), false);
        assert_eq!(can_construct_ransom_note("aa".to_string(), "aab".to_string()), true);
    }

    #[test]
    fn test_merge_alternately() {
        assert_eq!(merge_alternately("abc".to_string(), "pqr".to_string()), "apbqcr".to_string());
        assert_eq!(merge_alternately("ab".to_string(), "pqrs".to_string()), "apbqrs".to_string());
        assert_eq!(merge_alternately("abcd".to_string(), "pq".to_string()), "apbqcd".to_string());
    }

    #[test]
    fn test_find_max_consecutive_ones() {
        assert_eq!(find_max_consecutive_ones(vec![1, 1, 0, 1, 1, 1]), 3);
        assert_eq!(find_max_consecutive_ones(vec![1, 0, 1, 1, 0, 1]), 2);
    }

    #[test]
    fn test_sorted_squares() {
        assert_eq!(sorted_squares(vec![-4, -1, 0, 3, 10]), vec![0, 1, 9, 16, 100]);
        assert_eq!(sorted_squares(vec![-7, -3, 2, 3, 11]), vec![4, 9, 9, 49, 121]);
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

    #[test]
    fn test_merge_sorted() {
        let mut v1 = vec![1, 2, 3, 0, 0, 0];
        merge_sorted(&mut v1, 3, &mut vec![2, 5, 6], 3);
        assert_eq!(v1, vec![1, 2, 2, 3, 5, 6]);

        let mut v2 = vec![1];
        merge_sorted(&mut v2, 1, &mut vec![], 0);
        assert_eq!(v2, vec![1]);

        let mut v3 = vec![0];
        merge_sorted(&mut v3, 0, &mut vec![1], 1);
        assert_eq!(v3, vec![1]);

        let mut v4 = vec![2, 0];
        merge_sorted(&mut v4, 1, &mut vec![1], 1);
        assert_eq!(v4, vec![1, 2]);
    }

    #[test]
    fn test_remove_element() {
        let mut v1 = vec![3, 2, 2, 3];
        let r1 = remove_element(&mut v1, 3);
        assert_eq!(r1, 2);
        assert_eq!(v1[0], 2);
        assert_eq!(v1[1], 2);

        let mut v2 = vec![0, 1, 2, 2, 3, 0, 4, 2];
        let r2 = remove_element(&mut v2, 2);
        assert_eq!(r2, 5);
        assert_eq!(v2[0], 0);
        assert_eq!(v2[1], 1);
        assert_eq!(v2[2], 3);
        assert_eq!(v2[3], 0);
        assert_eq!(v2[4], 4);
    }

    #[test]
    fn test_remove_element2() {
        let mut v1 = vec![3, 2, 2, 3];
        let r1 = remove_element2(&mut v1, 3);
        assert_eq!(r1, 2);
        assert_eq!(v1[0], 2);
        assert_eq!(v1[1], 2);

        let mut v2 = vec![0, 1, 2, 2, 3, 0, 4, 2];
        let r2 = remove_element2(&mut v2, 2);
        assert_eq!(r2, 5);
        assert_eq!(v2[0], 0);
        assert_eq!(v2[1], 1);
        assert_eq!(v2[2], 3);
        assert_eq!(v2[3], 0);
        assert_eq!(v2[4], 4);
    }

    #[test]
    fn test_remove_duplicates() {
        let mut v1 = vec![1, 1, 2];
        let r1 = remove_duplicates(&mut v1);
        assert_eq!(r1, 2);
        assert_eq!(v1[0], 1);
        assert_eq!(v1[1], 2);

        let mut v2 = vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
        let r2 = remove_duplicates(&mut v2);
        assert_eq!(r2, 5);
        assert_eq!(v2[0], 0);
        assert_eq!(v2[1], 1);
        assert_eq!(v2[2], 2);
        assert_eq!(v2[3], 3);
        assert_eq!(v2[4], 4);
    }

    #[test]
    fn test_remove_duplicates2() {
        let mut v1 = vec![1, 1, 2];
        let r1 = remove_duplicates2(&mut v1);
        assert_eq!(r1, 2);
        assert_eq!(v1[0], 1);
        assert_eq!(v1[1], 2);

        let mut v2 = vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
        let r2 = remove_duplicates2(&mut v2);
        assert_eq!(r2, 5);
        assert_eq!(v2[0], 0);
        assert_eq!(v2[1], 1);
        assert_eq!(v2[2], 2);
        assert_eq!(v2[3], 3);
        assert_eq!(v2[4], 4);
    }

    // #[test]
    // fn test_find_max_consecutive_ones() {
    //     assert_eq!(find_max_consecutive_ones2(vec![1]), 1);
    //     assert_eq!(find_max_consecutive_ones2(vec![0]), 1);
    //     assert_eq!(find_max_consecutive_ones2(vec![1, 0]), 2);
    //     assert_eq!(find_max_consecutive_ones2(vec![1, 0, 1]), 3);
    //     assert_eq!(find_max_consecutive_ones2(vec![0, 0, 0]), 1);
    //     assert_eq!(find_max_consecutive_ones2(vec![1, 1]), 2);
    //     assert_eq!(find_max_consecutive_ones2(vec![1, 0, 1, 1, 0]), 4);
    //     assert_eq!(find_max_consecutive_ones2(vec![1, 0, 1, 1, 0, 1]), 4);
    // }
}
