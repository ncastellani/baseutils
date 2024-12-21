package baseutils

// FillString
// check if the current string is empty, if so, return the new string
func FillString(currVal, newVal string) string {
	if currVal == "" {
		return newVal
	}
	return currVal
}
