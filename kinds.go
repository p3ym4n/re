package re

import "net/http"

const (
	KindInvalid     Kind = "invalid"
	KindForbidden   Kind = "forbidden"
	KindNotFound    Kind = "notfound"
	KindUnexpected  Kind = "unexpected"
	KindUnavailable Kind = "unavailable"
)

var kinds = map[Kind]int{
	KindInvalid:     http.StatusUnprocessableEntity,
	KindForbidden:   http.StatusForbidden,
	KindNotFound:    http.StatusNotFound,
	KindUnexpected:  http.StatusInternalServerError,
	KindUnavailable: http.StatusServiceUnavailable,
}

func AddKind(k Kind, status int) {
	kinds[k] = status
}

func HttpCode(bag Error) int {
	if code, has := kinds[bag.Kind()]; has {
		return code
	}
	return http.StatusInternalServerError
}
