package apibanco

var code []ApiBanco

func CalculaCode() int64 {
	var soma int64
	soma = 0
	for _, v := range code {
		if v.Code != 0 {
			soma += v.Code
		}
	}
	return soma
}
