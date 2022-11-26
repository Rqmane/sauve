package piscine

func ConcatParams(args []string) string {
	mString := ""

	for v, str := range args {
		mString = mString + str
		if v < len(args)-1 {
			mString = mString + "\n"
		}
	}
	return mString
}
