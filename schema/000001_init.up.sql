CREATE TABLE users
(
    id serial PRIMARY KEY,
    username varchar(255) not null UNIQUE,
    firstname varchar(20) not null,
    lastname varchar(20),
    password_hash varchar(255) not null,
    email varchar(255) UNIQUE,
    userstatus varchar(300),
    imageURL varchar(255),
    createdAt timestamp not null
);
CREATE TABLE conversation
(
    id serial PRIMARY KEY,
    channel_id int not null UNIQUE,
    title varchar(40),
    creator_id int references users (id) on delete cascade not null,
    created_at timestamp
);
CREATE TABLE messages
(
    id serial PRIMARY KEY,
    message_text VARCHAR(1000) not null,
    user_id_from int references users (id) on delete cascade not null,
    conversation_id_To INT not null,
    CreatedAt timestamp not null,
    TypeOfMessage VARCHAR(10) not null
);
CREATE TABLE participants
(
    id serial PRIMARY KEY,
    conversation_id int,
    users_id int
);