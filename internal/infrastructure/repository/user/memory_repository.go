package repository

import (
	"errors"
	"nilus-challenge-backend/internal/domain/user"
	"sync"
)

type UserMemoryRepository struct {
	users  []user.User
	nextID int
	mutex  sync.RWMutex
}

func NewUserRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users:  make([]user.User, 0),
		nextID: 1,
	}
}

func (r *UserMemoryRepository) FindAll() ([]user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.users, nil
}

func (r *UserMemoryRepository) FindByID(id int) (*user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserMemoryRepository) Create(u *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	u.ID = r.nextID
	r.nextID++
	r.users = append(r.users, *u)
	return nil
}

func (r *UserMemoryRepository) Update(u *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, existing := range r.users {
		if existing.ID == u.ID {
			r.users[i] = *u
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserMemoryRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserMemoryRepository) OptOut(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, u := range r.users {
		if u.ID == id {
			r.users[i].OptOut = true
			return nil
		}
	}
	return errors.New("user not found")
}

