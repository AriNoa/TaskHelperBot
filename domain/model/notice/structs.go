package notice

// Table is a struct that represents a collection of Notices
type Table struct {
	ID      string
	Notices map[int]*Notice
}

// User is a struct that represents a collection of user's Tables
type User struct {
	ID     string
	Tables map[string]*Table
}
