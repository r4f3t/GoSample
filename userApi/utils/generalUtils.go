package utils

func CheckErrors(err error) {
	if err != nil {
		err.Error()
	}
}
