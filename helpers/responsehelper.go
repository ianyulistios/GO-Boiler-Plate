package helpers

//ErrResponseHelper Represented of function to handle Error Response
func ErrResponseHelper(err error, statusCode string) (map[string]interface{}, int) {
	responseJSONInterface := GetJSONResponse(statusCode)
	mapJSONInterface := responseJSONInterface.(map[string]interface{})
	messageBody := mapJSONInterface["body"].(map[string]interface{})["message"].(string)
	statusCodeResponse := mapJSONInterface["header"].(map[string]interface{})["code"].(int)
	errorMessage := err.Error()
	responseBody := map[string]interface{}{
		"code":    statusCodeResponse,
		"message": messageBody,
		"err":     errorMessage,
	}
	return responseBody, statusCodeResponse
}

// SuccessResponseHelper Represented of Function to handle transformation of JSON
func SuccessResponseHelper(data map[string]interface{}, statusCode string) (map[string]interface{}, int) {
	responseJSONInterface := GetJSONResponse(statusCode)
	mapJSONInterface := responseJSONInterface.(map[string]interface{})
	messageBody := mapJSONInterface["body"].(map[string]interface{})["message"].(string)
	statusCodeResponse := mapJSONInterface["header"].(map[string]interface{})["code"].(int)
	responseBody := map[string]interface{}{
		"code":    statusCodeResponse,
		"message": messageBody,
		"data":    data,
	}
	return responseBody, statusCodeResponse
}
