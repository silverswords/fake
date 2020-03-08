package err

import "log"

//Err is a check
func Err(err error) {
	if err != nil {
		log.Println("error:", err)
		return
	}
}
