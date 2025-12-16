package utils //?agar bisa direuse

import "github.com/abu-umair/lms-be-microservice/pb/common"

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message,
		// IsError:    false, //?defaultnya false
	}
}
