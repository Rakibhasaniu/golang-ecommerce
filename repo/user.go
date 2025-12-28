package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	CreateUser(u User) (*User, error)
	GetUserByEmail(email string, password string) (*User, error)
}
type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	return repo
}

func (u *userRepo) CreateUser(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}
	user.ID = len(u.users) + 1
	u.users = append(u.users, user)
	return &user, nil
}

func (u *userRepo) GetUserByEmail(email string, password string) (*User, error) {
	for _, u := range u.users {
		if u.Email == email && u.Password == password {
			return &u, nil
		}
	}
	return nil, nil
}
