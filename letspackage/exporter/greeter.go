package exporter

//Greeter function that greets ...
func Greeter(lang string, whom string) string {
	var greetMessage = getGreeter(lang) + ", " + whom
	//fmt.Println(greetMessage)
	return greetMessage
}