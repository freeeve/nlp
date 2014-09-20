package util

func Ngrams(seq []string, n int) [][]string {
	ret := [][]string{}
	cur := []string{}
	for _, x := range seq {
		cur = append(cur, x)
		if len(cur) == n {
			ret = append(ret, cur)
			cur = cur[1:]
		}
	}
	return ret
}
