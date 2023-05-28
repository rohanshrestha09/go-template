package enums

type Provider int

const (
	Facebook Provider = iota
	Google
)

type Model string

const (
	USER    Model = "user"
	PROJECT Model = "project"
	REQUEST Model = "request"
)

type ProjectStatus string

const (
	OPEN        ProjectStatus = "OPEN"
	IN_PROGRESS ProjectStatus = "IN_PROGRESS"
	COMPLETED   ProjectStatus = "COMPLETED"
)

type RequestStatus string

const (
	PENDING  RequestStatus = "PENDING"
	APPROVED RequestStatus = "APPROVED"
	REJECTED RequestStatus = "REJECTED"
)

type FileType string

const (
	IMAGE FileType = "image"
)

type FileDIR string

const (
	USER_DIR    FileDIR = "users/"
	PROJECT_DIR FileDIR = "projects/"
)
