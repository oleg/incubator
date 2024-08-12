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
#[test]
fn test_sort_array_by_parity() {
    assert_eq!(sort_array_by_parity(vec![3, 1, 2, 4]), vec![2, 4, 3, 1]);
    assert_eq!(sort_array_by_parity(vec![1, 2, 3, 4]), vec![2, 4, 3, 1]);
    assert_eq!(sort_array_by_parity(vec![0]), vec![0]);
    assert_eq!(sort_array_by_parity(vec![1]), vec![1]);
    assert_eq!(sort_array_by_parity(vec![2, 4]), vec![2, 4]);
    assert_eq!(sort_array_by_parity(vec![1, 3]), vec![1, 3]);
}

fn sorted_squares(mut nums: Vec<i32>) -> Vec<i32> {
    for i in 0..nums.len() {
        nums[i] *= nums[i];
    }
    nums.sort();
    nums
}
#[test]
fn test_sorted_squares() {
    assert_eq!(
        sorted_squares(vec![-4, -1, 0, 3, 10]),
        vec![0, 1, 9, 16, 100]
    );
    assert_eq!(
        sorted_squares(vec![-7, -3, 2, 3, 11]),
        vec![4, 9, 9, 49, 121]
    );
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
