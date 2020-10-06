package querys

// InsertNewMessage query
func InsertNewMessage() string {
	return `INSERT INTO message(Body, Processed) VALUES (?, ?)`
}
