package repos

import "log"

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

/** ini func SayHello() sebelum major changes
func SayHello() string {
	return "Hello World"
}
*/

//Major changes
func SayHello(name string) string {
	return "Hello " + name
}

/**
Apabila ada major changes disarankan buat modul baru dan naikkan versi
*/
