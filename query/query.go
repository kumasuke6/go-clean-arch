package query

type userQuery struct{}
type messageQuery struct{}

func NewUserQuery() *userQuery {
	return new(userQuery)
}

func (q *userQuery) Read() string {
	return "SELECT id, name, age, email, created_at, updated_at FROM users WHERE id = $1"
}

func (q *userQuery) Create() string {
	return "INSERT INTO users (name, age, email) VALUES ($1, $2, $3)"
}

func (q *userQuery) Update() string {
	return "UPDATE users SET name = $1, age = $2, email = $3 WHERE id = $4"
}

func (q *userQuery) Delete() string {
	return "DELETE FROM users WHERE id = $1"
}

func NewMessageQuery() *messageQuery {
	return new(messageQuery)
}

func (q *messageQuery) Read() string {
	return "SELECT id, user_id, message FROM messages WHERE user_id = $1"
}

func (q *messageQuery) Create() string {
	return "INSERT INTO messages (user_id, message) VALUES ($1, $2)"
}

func (q *messageQuery) Delete() string {
	return "DELETE FROM messages WHERE user_id = $1"
}
