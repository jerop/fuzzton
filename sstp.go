package main

func IsSSTP(data []byte) bool {
	if len(data) == 3 {
		if data[0] == 'S' && data[1] == 'S' && data[2] == 'T' && data[3] == 'P' {
			return true
		}
	}
	return false
}
