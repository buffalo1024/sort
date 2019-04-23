package main



func ShellSort(s []int) {
	var a []int = make([]int, len(s))
	copy(a, s)
	for k := len(a) / 2; k >= 1; k /= 2 {
		for i := 0; i < len(a); i++ {
			m := a[i]
			for j := i; j > 0; j -= k {
				if (j-k > 0 && a[j-k] <= m) && m < a[j] {
					for l := i; l > j; l -= k {
						a[l] = a[l-k]
					}
					a[j] = m
				}
			}
		}
	}
	printSlice(a)
}
