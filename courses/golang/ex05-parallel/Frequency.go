package letter

func _chFrequency(str string, ch chan map[string]int) {
	freqMap := make(map[string]int)

	for _, char := range str {
		strCh, _ := Downcase(string(char))
		_, ok := freqMap[strCh]

		if !ok {
			freqMap[strCh] = 1
		} else {
			freqMap[strCh]++
		}
	}

	ch <- freqMap
}

func Frequency(str string) map[string]int {
	ch := make(chan map[string]int)
	go _chFrequency(str, ch)
	return <-ch
}

func ConcurrentFrequency(strings []string) map[string]int {
	ch := make(chan map[string]int)

	for i := 0; i < len(strings); i++ {
		go _chFrequency(strings[i], ch)
	}

	freqMapRes := make(map[string]int)

	for i := 0; i < len(strings); i++ {
		freqMap := <-ch
		for letter, freq := range freqMap {
			_, ok := freqMapRes[letter]

			if !ok {
				freqMapRes[letter] = freq
			} else {
				freqMapRes[letter] += freq
			}
		}
	}

	return freqMapRes
}