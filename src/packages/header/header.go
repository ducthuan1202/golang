package header

func ContentType (contentType string) (key string, value string){

	key = "Content-Type"
	
	value = ""

	switch contentType {
	case "javascript":
		value = "application/javascript"
	case "text":
		value = "text/plain"
	default:
		value = "text/html"
	}

	return
}