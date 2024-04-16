package result

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) Result {
	return Result{
		Success: true,
		Message: "Success",
		Data:    data,
	}
}

func Fail(message string) Result {
	return Result{
		Success: false,
		Message: message,
		Data:    nil,
	}
}
