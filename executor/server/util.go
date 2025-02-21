package server

import "encoding/base64"

func decodeBase64(in string) string {
	out, _ := base64.StdEncoding.DecodeString(in)
	return string(out)
}
