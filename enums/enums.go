package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION production environment
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// DEVELOP development environment
	DEVELOP = ENVIRONMENT("DEVELOP")
	// TEST test environment
	TEST = ENVIRONMENT("TEST")
)

const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)

// Command kafka command
type Command string

const (
	// Kube object ADD command
	ADD = Command("ADD")
	// Kube object UPDATE command
	UPDATE = Command("UPDATE")
	// Kube object DELETE command
	DELETE = Command("DELETE")
)
