package pkg

import "github.com/thanhpk/randstr"

func Base64Encode() string {
	return randstr.String(30)
}
