package schoolerrors

import (
	"fmt"
)

type SchoolError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *SchoolError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

const (
	ErrCodeNotFound           = 404
	ErrCodeExists             = 409
	ErrCodeDecode             = 400
	ErrCodeUnauthorizedAccess = 401
	ErrCodeAccessStudent      = 403

)

var (
	ErrorUnauthorizedAccess  = &SchoolError{Code: ErrCodeUnauthorizedAccess, Message: "Unauthorized access"}
	ErrorInvalidHeaderFormat = &SchoolError{Code: ErrCodeUnauthorizedAccess, Message: "Invalid authorization header format"}
	ErrorInvalidToken        = &SchoolError{Code: ErrCodeUnauthorizedAccess, Message: "Invalid token"}
	ErrorStudentAccessOnly   = &SchoolError{Code: ErrCodeAccessStudent, Message: "Only Students have this access"}
	ErrorTeacherAccessOnly   = &SchoolError{Code: ErrCodeAccessStudent, Message: "Only Teachers have this access"}
	ErrorParentAccessOnly   = &SchoolError{Code: ErrCodeAccessStudent, Message: "Only parents have this access"}
)
