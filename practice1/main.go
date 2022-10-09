package main

func main() {
	strs := []string{"I", "am", "stupid", "and", "weak"}

	var results []string
	for _, str := range strs {
		if str == "stupid" {
			str = "smart"
		}
		if str == "weak" {
			str = "strong"
		}
		results = append(results, str)
	}

}
