package rest

func String2Json(s string) string {
	var (
		nam string
		val string
	)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			if s[i-1] == 58 {
				nam = string(34) + s[0:i-1] + string(34) + string(58)
				val = string(34) + s[i+1:] + string(34)
				break
			} else {
				nam = string(34) + s[0:i] + string(34) + string(58)
				val = string(34) + s[i+1:] + string(34)
				break
			}
		}
	}
	return "\t" + nam + " " + val
}
