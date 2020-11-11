package request

import (
	"Food/helpers/e"
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid bind data and return error if exist.
// @Return int errCode
func BindAndValid(c *gin.Context, form interface{}) (int) {
	err := c.Bind(form)
	if err != nil {
		return e.INVALID_REQUEST
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return e.INVALID_REQUEST
	}

	return e.SUCCESS
}

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		// logging.Info(err.Key, err.Message)
		fmt.Printf("MarkErrors : %v %v", err.Key, err.Message)
	}
}
