package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey %= 26
	res := make([]rune, 0, len(plain)) // 建议使用 slice 追加或预分配

	for _, p := range plain {
		var s rune
		if p >= 'a' && p <= 'z' {
			s = 'a' + (p-'a'+rune(shiftKey))%26
		} else if p >= 'A' && p <= 'Z' {
			s = 'A' + (p-'A'+rune(shiftKey))%26
		} else {
			s = p
		}
		res = append(res, s)
	}
	return string(res)
}
