package exporter

//GetGreeter Gets the greeter prefix ...
func getGreeter(lang string) string {
	if lang == "es" {
		return esGreeter
	}
	if lang == "en" {
		return enGreeter
	}
	return defaultGreeter
}
