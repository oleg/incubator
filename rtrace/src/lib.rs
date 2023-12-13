#[derive(Debug, PartialEq)]
pub struct Tuple {
    x: f32,
    y: f32,
    z: f32,
    w: f32,
}

impl Tuple {
    fn tuple(x: f32, y: f32, z: f32, w: f32) -> Tuple {
        Tuple { x, y, z, w }
    }

    fn point(x: f32, y: f32, z: f32) -> Tuple {
        Tuple { x, y, z, w: 1.0 }
    }

    fn vector(x: f32, y: f32, z: f32) -> Tuple {
        Tuple { x, y, z, w: 0.0 }
    }

    fn is_point(&self) -> bool {
        self.w == 1.0
    }
    fn is_vector(&self) -> bool {
        self.w == 0.0
    }
}

impl std::ops::Add for Tuple {
    type Output = Tuple;

    fn add(self, other: Tuple) -> Tuple {
        Tuple::tuple(
            self.x + other.x,
            self.y + other.y,
            self.z + other.z,
            self.w + other.w,
        )
    }
}

impl std::ops::Sub for Tuple {
    type Output = Tuple;

    fn sub(self, rhs: Self) -> Self::Output {
        Tuple::tuple(
            self.x - rhs.x,
            self.y - rhs.y,
            self.z - rhs.z,
            self.w - rhs.w,
        )
    }
}

#[cfg(test)]
mod tests {
    use crate::Tuple;

    #[test]
    fn test_a_tuple_is_a_point() {
        let a = Tuple::tuple(4.3, -4.2, 3.1, 1.0);

        assert_eq!(a.x, 4.3);
        assert_eq!(a.y, -4.2);
        assert_eq!(a.z, 3.1);
        assert_eq!(a.w, 1.0);
        assert!(a.is_point());
        assert!(!a.is_vector());
    }

    #[test]
    fn test_a_tuple_is_a_vector() {
        let a = Tuple::tuple(4.3, -4.2, 3.1, 0.0);

        assert_eq!(a.x, 4.3);
        assert_eq!(a.y, -4.2);
        assert_eq!(a.z, 3.1);
        assert_eq!(a.w, 0.0);
        assert!(!a.is_point());
        assert!(a.is_vector());
    }

    #[test]
    fn test_point_creates_tuples_with_w_1() {
        let p = Tuple::point(4.0, -4.0, 3.0);

        assert_eq!(p.x, 4.0);
        assert_eq!(p.y, -4.0);
        assert_eq!(p.z, 3.0);
        assert_eq!(p.w, 1.0);
        assert!(p.is_point());
        assert!(!p.is_vector());
    }

    #[test]
    fn test_vector_creates_tuples_with_w_0() {
        let v = Tuple::vector(4.0, -4.0, 3.0);

        assert_eq!(v.x, 4.0);
        assert_eq!(v.y, -4.0);
        assert_eq!(v.z, 3.0);
        assert_eq!(v.w, 0.0);
        assert!(!v.is_point());
        assert!(v.is_vector());
    }

    #[test]
    fn test_adding_two_tuples() {
        let a1 = Tuple::tuple(3.0, -2.0, 5.0, 1.0);
        let a2 = Tuple::tuple(-2.0, 3.0, 1.0, 0.0);

        assert_eq!(a1 + a2, Tuple::tuple(1., 1., 6., 1.));
    }

    #[test]
    fn test_subtracting_two_points() {
        let p1 = Tuple::point(3., 2., 1.);
        let p2 = Tuple::point(5., 6., 7.);

        assert_eq!(p1 - p2, Tuple::vector(-2., -4., -6.))
    }

    #[test]
    fn test_subtracting_vector_from_point() {
        let p = Tuple::point(3., 2., 1.);
        let v = Tuple::vector(5., 6., 7.);

        assert_eq!(p - v, Tuple::point(-2., -4., -6.))
    }


    #[test]
    fn test_subtracting_two_vectors() {
        let v1 = Tuple::vector(3., 2., 1.);
        let v2 = Tuple::vector(5., 6., 7.);

        assert_eq!(v1 - v2, Tuple::vector(-2., -4., -6.))
    }
}




/*
    //todo add this to partial_eq?
 	constant EPSILON ← 0.00001
 	function equal(a, b)
 	  if abs(a - b) < EPSILON
 	    return true
 	  else
 	    return false
 	  end if
 	end function
 */