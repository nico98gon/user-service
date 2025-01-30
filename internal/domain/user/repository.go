package user

type Repository interface {
	FindAll() ([]User, error)
	FindByID(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
	OptOut(id int) error
}
