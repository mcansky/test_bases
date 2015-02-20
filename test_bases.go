package test_bases

func Search(numbers []int, val int) int {
	t := len(numbers) / 2
	first_half := numbers[0:t]
	last_half := numbers[t:len(numbers)]
	res := -1
	if len(numbers) == 0 {
	} else if first_half[len(first_half)-1] > val {
		res = Search(first_half, val)
	} else if last_half[0] < val {
		res = Search(last_half, val)
	} else if numbers[t] == val {
		res = t
	}
	return res
}

/*
func main() {
	s := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		s[i] = i + 1
	}
	search(s, 3)
}*/
