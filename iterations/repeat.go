package iteration

var repeatCount = 5

func Repeat(character string, limit int) string {
	var repeated string
	if limit>0{
	repeatCount = limit
	}
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
