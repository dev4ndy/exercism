package triangle

import "math"

// Kind s
type Kind string

// possible values
const (
	NaT Kind = "Not a triangle"
	Equ Kind = "Equilateral"
	Iso Kind = "Isosceles"
	Sca Kind = "Scalene"
)

// KindFromSides return type of a triangle according to its sides a, b, c
func KindFromSides(a, b, c float64) Kind {
	if IsNotPosibleTriangle(a, b, c) {
		return NaT
	}
	if IsEquilateral(a, b, c) {
		return Equ
	} else if IsIsosceles(a, b, c) {
		return Iso
	} else if IsScalene(a, b, c) {
		return Sca
	}
	return NaT
}

// IsPosibleTriangle this function valid that the sides form a triangle
func IsPosibleTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) { // No accept infinite
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 { // All slides greater than zero
		return false
	}
	if (a+b) < c || (a+c) < b || (b+c) < a { //the sum of the lengths of any two sides must be greater
		return false
	}
	return true
}

// IsNotPosibleTriangle this function valid that the sides NOT form a triangle
func IsNotPosibleTriangle(a, b, c float64) bool {
	return !IsPosibleTriangle(a, b, c)
}

// IsEquilateral Valid that the sides form a equilateral triangle
func IsEquilateral(a, b, c float64) bool {
	return a == b && b == c && c == a
}

// IsIsosceles Valid that the sides form a isosceles triangle
func IsIsosceles(a, b, c float64) bool {
	return a == b || b == c || c == a
}

// IsScalene Valid that the sides form a scalene triangle
func IsScalene(a, b, c float64) bool {
	return a != b && b != c && c != a
}
