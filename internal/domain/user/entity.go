package user

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID       	 	int    `json:"id"`
	Name      	string `json:"name"`
	Email     	string `json:"email"`
	OptOut    	bool   `json:"opt_out"`
	LocalityID 	*int   `json:"locality_id"`
	CreatedAt  	time.Time `json:"created_at"`
	UpdatedAt  	time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("nombre es requerido")
	}
	if len(u.Name) < 3 {
		return errors.New("nombre debe tener al menos 3 caracteres")
	}
	if u.Email == "" {
		return errors.New("email es requerido")
	}
	if !isValidEmail(u.Email) {
		return errors.New("email no es válido")
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
