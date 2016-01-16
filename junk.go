package main

type junk struct{}

func (j *junk) Write(p []byte) (n int, err error) {
	return 0, nil
}
