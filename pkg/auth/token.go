package auth

import "regexp"

const (
	TOKEN_PREFIX = "stt_"
)

var (
	TOKEN_REGEX = regexp.MustCompile(`^` + TOKEN_PREFIX + `[0-7][0-9A-HJKMNP-TV-Z]{25}$`)
)
