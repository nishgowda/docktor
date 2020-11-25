package suggestions


// DockerVars type is the attributes of the docker file
type DockerVars struct {
	userCount int
	workDirCount int
	environCount int
}

// ErrorMessages assigns messages to possible suggestions
type ErrorMessages struct {
	userMsg string
	workDirMsg string
	environMsg string
}