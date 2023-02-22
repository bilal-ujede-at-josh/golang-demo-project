package auth

type SendOtpRequest struct {
	Mobile string
}

type VerifyOtpRequest struct {
	Mobile string
	Otp    int
}
