package model

type ReturnMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Type    string `json:"type"`
	Data    any    `json:"data"`
}

func GetErrorMessage(message string) ReturnMessage {
	var msg ReturnMessage
	msg.Code = 200
	msg.Message = message
	msg.Success = false
	msg.Type = "error"
	msg.Data = nil
	return msg
}

func GetFatalMessaage(message string) ReturnMessage {
	var msg ReturnMessage
	msg.Code = 500
	msg.Message = message
	msg.Success = false
	msg.Type = "error"
	msg.Data = nil
	return msg
}

// func GetSuccessMessage(message string) ReturnMessage {
// 	var msg ReturnMessage
// 	msg.Code = 200
// 	msg.Message = message
// 	msg.Success = false
// 	msg.Type = "success"
// 	msg.Data = nil
// 	return msg
// }

func GetSuccessMessage(message string, data any) ReturnMessage {
	var msg ReturnMessage
	msg.Code = 200
	msg.Message = message
	msg.Success = true
	msg.Type = "success"
	msg.Data = data
	return msg
}

func GetWarningMessage(message string) ReturnMessage {
	var msg ReturnMessage
	msg.Code = 200
	msg.Message = message
	msg.Success = false
	msg.Type = "warnng"
	msg.Data = nil
	return msg
}
