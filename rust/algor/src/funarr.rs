fn height_checker(heights: Vec<i32>) -> i32 {
    let mut heights_sorted = heights.clone();
    heights_sorted.sort();
    let mut count = 0;
    for i in 0..heights.len() {
        if heights[i] != heights_sorted[i] {
            count += 1;
        }
    }
    count
}
#[test]
fn test_height_checker() {
    assert_eq!(height_checker(vec![5, 1, 2, 3, 4]), 5);
    assert_eq!(height_checker(vec![1, 2, 3, 4, 5]), 0);
}

fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
    let mut next_zero_done = 0;
    let mut next_zero_continue = 0;

    let mut max = 0;
    for num in nums {
        if num == 0 {
            max = max.max(next_zero_done);
            next_zero_done = next_zero_continue + 1;
            next_zero_continue = 0;
        } else {
            next_zero_done += 1;
            next_zero_continue += 1;
        }
    }
    max.max(next_zero_done)
}
#[test]
fn test_find_max_consecutive_ones() {
    assert_eq!(find_max_consecutive_ones(vec![1, 0, 1, 1, 0]), 4);
    assert_eq!(find_max_consecutive_ones(vec![1, 0, 1, 1, 0, 1]), 4);
}

fn find_max_consecutive_ones2(nums: Vec<i32>) -> i32 {
    let mut max = 0;
    let mut zeros = 0;
    let mut l = 0;
    let mut r = 0;
    let len = nums.len();
    while r < len {
        if nums[r] == 0 {
            zeros += 1;
        }
        while zeros == 2 {
            if nums[l] == 0 {
                zeros -= 1;
            }
            l += 1;
        }
        max = max.max(r - l + 1);
        r += 1;
    }
    max as i32
}
#[test]
fn test_find_max_consecutive_ones2() {
    assert_eq!(find_max_consecutive_ones2(vec![1, 0, 1, 1, 0]), 4);
    assert_eq!(find_max_consecutive_ones2(vec![1, 0, 1, 1, 0, 1]), 4);
}

fn third_max(nums: Vec<i32>) -> i32 {
    let mut max1 = nums[0];
    for &n in &nums {
        if n > max1 {
            max1 = n;
        }
    }
    let mut max2 = nums[0];
    for &n in &nums {
        if (n > max2 && n != max1) || max2 == max1 {
            max2 = n
        }
    }
    let mut max3 = nums[0];
    for &n in &nums {
        if (n > max3 && n != max1 && n != max2) || max3 == max1 || max3 == max2 {
            max3 = n
        }
    }
    if max1 != max2 && max2 != max3 {
        return max3;
    }
    max1
}
#[test]
fn test_third_max() {
    assert_eq!(third_max(vec![3, 2, 1]), 1);
    assert_eq!(third_max(vec![1, 2]), 2);
    assert_eq!(third_max(vec![2, 2, 3, 1]), 1);
    assert_eq!(third_max(vec![1, 2, 2, 5, 3, 5]), 2);
}


fn find_disappeared_numbers(nums: Vec<i32>) -> Vec<i32> {
    let size = nums.len();
    let mut v = vec![0; size];
    for n in nums {
        v[(n - 1) as usize] = 1;
    }
    let mut r: Vec<i32> = vec![];
    for i in 0..v.len() {
        if v[i] == 0 {
            r.push((i + 1) as i32)
        }
    }
    r
}
#[test]
fn test_find_disappeared_numbers() {
    assert_eq!(find_disappeared_numbers(vec![4, 3, 2, 7, 8, 2, 3, 1]), vec![5, 6]);
    assert_eq!(find_disappeared_numbers(vec![1, 1]), vec![2]);
}