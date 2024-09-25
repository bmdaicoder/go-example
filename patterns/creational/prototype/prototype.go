package prototype

type Prototype interface {
	Clone() Prototype
	GetInfo() string
}

type User struct {
	Name string
}

func (u *User) Clone() Prototype {
	clone := *u
	return &clone
}

func (u *User) GetInfo() string {
	return u.Name
}
