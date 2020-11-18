package gorexos

// Transformation is used for any transformation in the REXos system. It
// uses quaternions for rotation and only one scale value.
// Please do not use this struct directly, but use the NewTransformation() instead!
type Transformation struct {
	Translation Vec3f `json:"translation"` // x/y/z
	Rotation    Vec4f `json:"rotation"`    // x/y/z/w
}

// TransformationWithScale takes a normal transformation and adds the scale
// value to it
type TransformationWithScale struct {
	Transformation
	Scale float32 `json:"scale" example:"0.5"` // one scale for all axis
}

// NewTransformation generates a valid transformation where W is set properly
func NewTransformation() Transformation {
	return Transformation{
		Translation: Vec3f{0.0, 0.0, 0.0},
		Rotation:    Vec4f{0.0, 0.0, 0.0, 1.0},
	}
}

// NewTransformationWithScale generates a valid transformation where scale and W is set properly
func NewTransformationWithScale() TransformationWithScale {
	t := NewTransformation()
	return TransformationWithScale{
		Transformation: t,
		Scale:          1.0,
	}
}
