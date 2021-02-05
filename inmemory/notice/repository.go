package notice

import (
	"encoding/json"
	"errors"
	"log"

	model "github.com/AriNoa/TaskHelperBot/domain/notice/model"
)

// Repository is a struct that implements domain/notice/model/Repository interface.
// It stores data only in memory
type Repository struct {
	Users map[string]([]byte)
}

// Create creates a new user
func (r *Repository) Create(ID string) error {
	if _, ok := r.Users[ID]; ok {
		return errors.New("already created")
	}

	user := model.User{}
	b, err := json.Marshal(user)

	if err != nil {
		return nil
	}

	r.Users[ID] = b

	return nil
}

// Read reads the user
func (r *Repository) Read(ID string) (model.User, bool) {
	user := model.User{}

	b, ok := r.Users[ID]

	if !ok {
		return user, false
	}

	if err := json.Unmarshal(b, &user); err != nil {
		log.Fatal(err)
	}

	return user, true
}

// Update updates the user
func (r *Repository) Update(ID string, user model.User) error {
	if _, ok := r.Users[ID]; !ok {
		return errors.New("not created")
	}

	b, err := json.Marshal(user)

	if err != nil {
		log.Fatal(err)
	}

	r.Users[ID] = b

	return nil
}

// Delete deletes the user
func (r *Repository) Delete(ID string) error {
	if _, ok := r.Users[ID]; !ok {
		return errors.New("not created")
	}

	delete(r.Users, ID)

	return nil
}
