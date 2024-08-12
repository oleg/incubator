use std::collections::HashMap;

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
#[test]
fn test_merge_alternately() {
    assert_eq!(
        merge_alternately("abc".to_string(), "pqr".to_string()),
        "apbqcr".to_string()
    );
    assert_eq!(
        merge_alternately("ab".to_string(), "pqrs".to_string()),
        "apbqrs".to_string()
    );
    assert_eq!(
        merge_alternately("abcd".to_string(), "pq".to_string()),
        "apbqcd".to_string()
    );
}


fn can_construct_ransom_note(ransom_note: String, magazine: String) -> bool {
    let mut letters = HashMap::new();
    for c in magazine.chars() {
        letters.entry(c).and_modify(|e| *e += 1).or_insert(1);
    }

    for c in ransom_note.chars() {
        let v = letters.entry(c).and_modify(|e| *e -= 1).or_insert(-1);
        if *v < 0 {
            return false;
        };
    }

    true
}

#[test]
fn test_can_construct_ransom_note() {
    assert_eq!(
        can_construct_ransom_note("a".to_string(), "b".to_string()),
        false
    );
    assert_eq!(
        can_construct_ransom_note("aa".to_string(), "ab".to_string()),
        false
    );
    assert_eq!(
        can_construct_ransom_note("aa".to_string(), "aab".to_string()),
        true
    );
}

