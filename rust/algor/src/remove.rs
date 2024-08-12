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



