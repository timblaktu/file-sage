package content

import (
	"net/http"
)

func GetType(hdr []byte) string {
	return http.DetectContentType(hdr)
}

func IsVideo(hdr []byte) bool {
	switch ft := GetType(hdr); ft {
	case "video":
		return true
	default:
		return false
	}
}

func IsImage(hdr []byte) bool {
	switch ft := GetType(hdr); ft {
	case "image":
		return true
	default:
		return false
	}
}
