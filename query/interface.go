package query

type UserQuery interface {
	Read() string
	Create() string
	Update() string
	Delete() string
}

type MessageQuery interface {
	Read() string
	Create() string
	Delete() string
}
