fn main() {
    println!("Hello, world!");
}

fn running_sum(nums: Vec<i32>) -> Vec<i32> {
    todo!()
}

#[cfg(test)]
mod tests {
    use crate::running_sum;

    #[test]
    fn test1() {
        let input = vec![1, 2, 3, 4];
        let result = running_sum(input);
        assert_eq!(result, vec![1, 3, 6, 10]);
    }

    #[test]
    fn test2() {
        let input = vec![1, 1, 1, 1, 1];
        let result = running_sum(input);
        assert_eq!(result, vec![1, 2, 3, 4, 5]);
    }

    #[test]
    fn test3() {
        let input = vec![3, 1, 2, 10, 1];
        let result = running_sum(input);
        assert_eq!(result, vec![3, 4, 6, 16, 17]);
    }
}
