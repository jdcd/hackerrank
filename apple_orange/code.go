package apple_orange

/* context of the problem in https://www.hackerrank.com/challenges/apple-and-orange/problem?isFullScreen=true*/
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) []int32 {
	var appleCounter, orangeCounter int32

	for _, apple := range apples {
		applePosition := apple + a
		if applePosition >= s && applePosition <= t {
			appleCounter = appleCounter + 1
		}
	}

	for _, orange := range oranges {
		orangePosition := orange + b
		if orangePosition >= s && orangePosition <= t {
			orangeCounter = orangeCounter + 1
		}
	}

	return []int32{appleCounter, orangeCounter}
}
