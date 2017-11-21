package result

type Result struct {
	inner interface{}
	err   error
}
