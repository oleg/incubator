#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
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
#[test]
fn test_middle_node() {
    assert_eq!(
        middle_node(Some(Box::new(ListNode::from(vec![1, 2, 3, 4, 5])))),
        Some(Box::new(ListNode::from(vec![3, 4, 5])))
    );

    assert_eq!(
        middle_node(Some(Box::new(ListNode::from(vec![1, 2, 3, 4, 5, 6])))),
        Some(Box::new(ListNode::from(vec![4, 5, 6])))
    );
}
