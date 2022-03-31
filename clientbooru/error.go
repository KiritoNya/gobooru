package clientbooru

type Error interface {
	Error() string
	StatusCode() int
}