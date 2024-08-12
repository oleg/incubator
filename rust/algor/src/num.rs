
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
#[test]
fn test_number_of_steps() {
    assert_eq!(number_of_steps(14), 6);
    assert_eq!(number_of_steps(8), 4);
    assert_eq!(number_of_steps(123), 12);
}
