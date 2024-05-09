package utils

func GetLinks(chunk []string) int {
	var links int

	for i := 0; i < len(chunk)-1; i++ {

		switch chunk[i] {
		case "A":
			if chunk[i+1] == "T" {
				links++
				i++
			}
		case "T":
			if chunk[i+1] == "A" {
				links++
				i++
			}
		case "G":
			if chunk[i+1] == "C" {
				links++
				i++
			}
		case "C":
			if chunk[i+1] == "G" {
				links++
				i++
			}
		}
	}

	return links
}
