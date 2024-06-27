CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name varchar(255),
	email varchar(255),
	age integer,
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE messages (
	id SERIAL PRIMARY KEY,
	user_id integer,
	message varchar(255)
);

-- INSERT INTO users (name, email, age, created_at, updated_at) VALUES ('test', 'test@test.com', 25, NOW(), NOW());