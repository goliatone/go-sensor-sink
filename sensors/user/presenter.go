package user

func ReadResponse(record User) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   record,
	}
}

func ListResponse(records []User) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   records,
	}
}
