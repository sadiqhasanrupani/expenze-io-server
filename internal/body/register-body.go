package body

type RegistrationBody struct {
	Firstname     string `binding:"required" json:"firstname"`
	Lastname      string `binding:"required" json:"lastname"`
	EmailID       string `binding:"required" json:"emailId"`
	MobilieNumber string `binding:"required" json:"mobileNumber"`
	PhoneCode     string `binding:"required" json:"phonecode"`
	Password      string `binding:"required" json:"password"`
}
