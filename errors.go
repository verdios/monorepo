package main

type noInputFileError struct{}

type fileNotFoundError struct{}

func (err noInputFileError) Error() string {
	return "No file was inputed"
}

func (err fileNotFoundError) Error() string {
	return "File not found"
}
