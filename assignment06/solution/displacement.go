package solution

// returns a function that calculates displacement based on time t
func GenDisplaceFn(a, vo, so float32) func(t float32) float32 {
	return func(t float32) float32 {
		return 0.5*a*t*t + vo*t + so
	}
}
