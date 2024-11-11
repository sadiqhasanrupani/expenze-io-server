package body

type VerifyOtpBody struct {
	MobileOtp string
	EmailOtp  string
	Token     string
	UserId    int64
}
