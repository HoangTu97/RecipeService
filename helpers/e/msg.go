package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_REQUEST:                  "INVALID_REQUEST",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token timeout",
	ERROR_AUTH_TOKEN:                "Token invalid",
	ERROR_AUTH:                      "Token empty",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
