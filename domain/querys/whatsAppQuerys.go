package querys

// InsertNewMessage query
func InsertNewMessage() string {
	return `INSERT INTO message(Body, Processed) VALUES (?, ?)`
}

func UpdateMessage() string {
	return `UPDATE message SET Processed = TRUE WHERE Id = ?`
}
