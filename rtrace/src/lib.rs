#[derive(Debug, Copy, Clone)]
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

    fn magnitude(&self) -> f32 {
        (self.x.powi(2)
            + self.y.powi(2)
            + self.z.powi(2)
            + self.w.powi(2)
        ).sqrt()
    }

    fn normalize(&self) -> Tuple {
        *self / self.magnitude()
    }

    fn dot(&self, other: Tuple) -> f32 {
        self.x * other.x
            + self.y * other.y
            + self.z * other.z
            + self.w * other.w
    }

    fn cross(&self, other: Tuple) -> Tuple {
        Tuple::vector(
            self.y * other.z - self.z * other.y,
            self.z * other.x - self.x * other.z,
            self.x * other.y - self.y * other.x,
        )
    }
}

const EPSILON: f32 = 0.00001;

fn close_enough(a: f32, b: f32) -> bool {
    (a - b).abs() < EPSILON
}

impl PartialEq for Tuple {
    fn eq(&self, other: &Self) -> bool {
        close_enough(self.x, other.x)
            && close_enough(self.y, other.y)
            && close_enough(self.z, other.z)
            && close_enough(self.w, other.w)
    }

    fn ne(&self, other: &Self) -> bool {
        !self.eq(other)
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

    fn sub(self, other: Self) -> Self::Output {
        Tuple::tuple(
            self.x - other.x,
            self.y - other.y,
            self.z - other.z,
            self.w - other.w,
        )
    }
}

impl std::ops::Neg for Tuple {
    type Output = Tuple;

    fn neg(self) -> Self::Output {
        Tuple::tuple(-self.x, -self.y, -self.z, -self.w)
    }
}

impl std::ops::Mul<f32> for Tuple {
    type Output = Tuple;

    fn mul(self, other: f32) -> Tuple {
        Tuple::tuple(
            self.x * other,
            self.y * other,
            self.z * other,
            self.w * other,
        )
    }
}

impl std::ops::Div<f32> for Tuple {
    type Output = Tuple;

    fn div(self, other: f32) -> Tuple {
        Tuple::tuple(
            self.x / other,
            self.y / other,
            self.z / other,
            self.w / other,
        )
    }
}

#[cfg(test)]
mod tests {
    use approx::assert_relative_eq;
    use crate::Tuple;
    use crate::EPSILON;

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

    #[test]
    fn test_subtracting_a_vector_from_the_zero_vector() {
        let zero = Tuple::vector(0., 0., 0.);
        let v = Tuple::vector(1., -2., 3.);

        assert_eq!(zero - v, Tuple::vector(-1., 2., -3.));
    }

    #[test]
    fn test_negating_a_tuple() {
        let a = Tuple::tuple(1., -2., 3., -4.);

        assert_eq!(-a, Tuple::tuple(-1., 2., -3., 4.));
    }

    #[test]
    fn test_multiplying_a_tuple_by_a_scalar() {
        let a = Tuple::tuple(1., -2., 3., -4.);

        assert_eq!(a * 3.5, Tuple::tuple(3.5, -7., 10.5, -14.));
    }


    #[test]
    fn test_multiplying_a_tuple_by_a_fraction() {
        let a = Tuple::tuple(1., -2., 3., -4.);

        assert_eq!(a * 0.5, Tuple::tuple(0.5, -1., 1.5, -2.));
    }

    #[test]
    fn test_dividing_a_tuple_by_a_scalar() {
        let a = Tuple::tuple(1., -2., 3., -4.);

        assert_eq!(a / 2., Tuple::tuple(0.5, -1., 1.5, -2.));
    }

    #[test]
    fn test_magnitude_of_vector_1_0_0() {
        let v = Tuple::vector(1., 0., 0.);

        assert_eq!(v.magnitude(), 1.);
    }

    #[test]
    fn test_magnitude_of_vector_0_1_0() {
        let v = Tuple::vector(0., 1., 0.);

        assert_eq!(v.magnitude(), 1.);
    }

    #[test]
    fn test_magnitude_of_vector_0_0_1() {
        let v = Tuple::vector(0., 0., 1.);

        assert_eq!(v.magnitude(), 1.);
    }

    #[test]
    fn test_magnitude_of_vector_1_2_3() {
        let v = Tuple::vector(1., 2., 3.);

        assert_eq!(v.magnitude(), 14_f32.sqrt());
    }

    #[test]
    fn test_magnitude_of_vector_neg_1_neg_2_neg_3() {
        let v = Tuple::vector(-1., -2., -3.);

        assert_eq!(v.magnitude(), 14_f32.sqrt());
    }

    #[test]
    fn test_normalizing_vector_4_0_0_gives_1_0_0() {
        let v = Tuple::vector(4., 0., 0.);

        assert_eq!(v.normalize(), Tuple::vector(1., 0., 0.));
    }

    #[test]
    fn test_normalizing_vector_1_2_3() {
        let v = Tuple::vector(1., 2., 3.);
        let v = v.normalize();

        // vector(1/√14,    2/√14,   3/√14)
        assert_relative_eq!(v.x, 0.26726, epsilon = EPSILON);
        assert_relative_eq!(v.y, 0.53452, epsilon = EPSILON);
        assert_relative_eq!(v.z, 0.80178, epsilon = EPSILON);
        assert_relative_eq!(v.w, 0.)
    }

    #[test]
    fn test_magnitude_of_a_normalized_vector() {
        let v = Tuple::vector(1., 2., 3.);
        let norm = v.normalize();

        assert_relative_eq!(norm.magnitude(), 1., epsilon = EPSILON);
    }

    #[test]
    fn test_dot_product_of_two_tuples() {
        let a = Tuple::vector(1., 2., 3.);
        let b = Tuple::vector(2., 3., 4.);

        assert_eq!(a.dot(b), 20.);
    }

    #[test]
    fn the_cross_product_of_two_vectors() {
        let a = Tuple::vector(1., 2., 3.);
        let b = Tuple::vector(2., 3., 4.);

        assert_eq!(a.cross(b), Tuple::vector(-1., 2., -1.));
        assert_eq!(b.cross(a), Tuple::vector(1., -2., 1.));
    }
}
