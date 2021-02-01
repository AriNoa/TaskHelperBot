package notice

// Table is a struct that represents a collection of Notices
type Table struct {
	ID      string
	Notices map[int]*Notice
}

// NewTable is a constructor for Table
func NewTable(id string, notices []*Notice) *Table {
	hash := map[int]*Notice{}

	for _, n := range notices {
		hash[n.ID] = n
	}

	return &Table{
		id,
		hash,
	}
}

// User is a struct that represents a collection of user's Tables
type User struct {
	ID     string
	Tables map[string]*Table
}

// NewUser is a constructor for User
func NewUser(id string, tables []*Table) *User {
	hash := map[string]*Table{}

	for _, t := range tables {
		hash[t.ID] = t
	}

	return &User{
		id,
		hash,
	}
}
