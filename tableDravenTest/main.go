package main

func clamp(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func main() {

}
