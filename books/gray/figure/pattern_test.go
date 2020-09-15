package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"testing"
)

func Test_creating_stripe_pattern(t *testing.T) {
	pattern := MakeStripePattern(oned.White, oned.Black)

	assert.Equal(t, oned.White, pattern.A)
	assert.Equal(t, oned.Black, pattern.B)
}

func Test_stripe_pattern_is_constant_in_y(t *testing.T) {
	pattern := MakeStripePattern(oned.White, oned.Black)

	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 1, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 2, 0}))
}

func Test_stripe_pattern_is_constant_in_z(t *testing.T) {
	pattern := MakeStripePattern(oned.White, oned.Black)

	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 1}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 2}))
}

func Test_stripe_pattern_alternates_in_x(t *testing.T) {
	pattern := MakeStripePattern(oned.White, oned.Black)

	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0.9, 0, 0}))

	assert.Equal(t, oned.Black, pattern.StripeAt(oned.Point{1, 0, 0}))
	assert.Equal(t, oned.Black, pattern.StripeAt(oned.Point{-0.1, 0, 0}))

	assert.Equal(t, oned.Black, pattern.StripeAt(oned.Point{-1, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{-1.1, 0, 0}))
}

func Test_stripes_with_object_transformation(t *testing.T) {
	object := MakeSphereT(multid.Scaling(2, 2, 2))
	pattern := MakeStripePattern(oned.White, oned.Black)

	c := pattern.StripeAtObject(object, oned.Point{1.5, 0, 0})

	assert.Equal(t, oned.White, c)
}

func Test_stripes_with_pattern_transformation(t *testing.T) {
	object := MakeSphere()
	pattern := MakeStripePatternT(oned.White, oned.Black, multid.Scaling(2, 2, 2))

	c := pattern.StripeAtObject(object, oned.Point{1.5, 0, 0})

	assert.Equal(t, oned.White, c)
}

func Test_stripes_with_both_object_and_pattern_transformation(t *testing.T) {
	object := MakeSphereT(multid.Scaling(2, 2, 2))
	pattern := MakeStripePatternT(oned.White, oned.Black, multid.Translation(0.5, 0, 0))

	c := pattern.StripeAtObject(object, oned.Point{2.5, 0, 0})

	assert.Equal(t, oned.White, c)
}
