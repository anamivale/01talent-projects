package middlewares

import "fmt"

func CheckErr(err error, errorMsg string) {
	if err != nil {

		fmt.Println(errorMsg + ";" + err.Error())
		return
	}
}
