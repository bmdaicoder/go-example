package bridge

type Database interface {
	Create(input map[string]any, result any) error
	Find(id string, result any) error
}

type User struct {
	FirstName string
	LastName  string
	Email     string
}

type UserRepository struct {
	db Database
}

func NewUserRepository(db Database) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user User) error {
	data := map[string]any{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
	}

	err := r.db.Create(data, &user)

	return err
}

func (r *UserRepository) Find(id string) (User, error) {
	var result User

	err := r.db.Find(id, &result)

	return result, err
}
