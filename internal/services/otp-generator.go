package services

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type OTPService struct {
	length int
}

func NewOTPService(length int) *OTPService {
	return &OTPService{
		length: length,
	}
}

func (s *OTPService) GenerateOTP() (string, error) {
	if s.length <= 0 {
		return "", fmt.Errorf("invalid OTP length")
	}

	var otpBuilder strings.Builder

	for i := 0; i < s.length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("failed to generate OTP: %w", err)
		}
		otpBuilder.WriteString(num.String())
	}

	return otpBuilder.String(), nil
}
