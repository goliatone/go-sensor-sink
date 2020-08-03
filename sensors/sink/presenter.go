package sink

//ReadResponse format a single record response
func ReadResponse(record DHT22Reading) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   record,
	}
}

//ListResponse format a list of records
func ListResponse(records []DHT22Reading) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   records,
	}
}

func AggregateResponse(records []SensorReadingAggregate) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   records,
	}
}
