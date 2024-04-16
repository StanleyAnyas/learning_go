package others

import (
	"errors"
	"fmt"
)

func checkIfValid(params interface{}) (bool, error) {
	_, ok := params.(string)
	if ok {
		return true, nil
	}
	return false, errors.New("it is not a valid")
}

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d: %v", e.arg, e.message)
}

func checkArg(num int) (int, error) {
	if num == 21 {
		return -1, &argError{num, "is not a valid number"}
	}
	return num + 3, nil
}

func MainFuncForError() {
	checkIfValidString, err := checkIfValid(12)
	if err != nil {
		fmt.Println(checkIfValidString)
	} else {
		fmt.Println(err)
	}

	checkArg, erro := checkArg(21)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(checkArg)
	}

	fmt.Println("done")

}
