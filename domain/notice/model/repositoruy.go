package notice

// UserRepository is a repository interface for User
type UserRepository interface {
	Create(ID string) error
	Read(ID string) (User, bool)
	Update(ID string, user User) error
	Delete(ID string) error
}
