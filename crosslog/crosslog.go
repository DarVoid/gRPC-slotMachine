package crosslog

//checks for error and calls callback
func Check(err error, handlerCallback func(error)) {
	if err != nil {
		handlerCallback(err)
	}
}

//add diferent handlers or a factory of errors below
