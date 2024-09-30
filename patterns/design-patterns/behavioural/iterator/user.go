package iterator

type Collection interface {
	CreateIterator() Iterator
}

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		Users: u.Users,
	}
}

type User struct {
	Name string
	Age  int
}
