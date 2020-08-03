package device

//ReadResponse format a single record response
func ReadResponse(record Device) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   record,
	}
}

//ListResponse format a list of records
func ListResponse(records []Device) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   records,
	}
}
