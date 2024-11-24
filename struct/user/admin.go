package user

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string, user User) (Admin, error) {

	return Admin{
		email:    email,
		password: password,
		User:     user,
	}, nil
}
