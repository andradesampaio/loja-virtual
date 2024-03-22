package exception

func Error(err interface{}) {
	// Error
	if err != nil {
		panic(err)
	}
}