package service

func Repeat(s string, rep_num int) string {
	var repeated string
	for i := 0; i < rep_num; i++ {
		repeated = repeated + s
	}
	return repeated
}
