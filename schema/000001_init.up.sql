CREATE TABLE users
(
    id VARCHAR(20) not null UNIQUE,
    username varchar(255) not null UNIQUE,
    firstname varchar(20) not null,
    lastname varchar(20),
    password_hash varchar(255) not null,
    phone varchar(11) not null UNIQUE,
    email varchar(255) not null UNIQUE,
    userstatus varchar(300),
    imageURL varchar(255),
    createdAt timestamp not null
);

CREATE TABLE messages
(
    id VARCHAR(20) not null UNIQUE,
    message_text VARCHAR(1000) not null,
    user_id_from VARCHAR(10) not NULL,
    conversation_id_To VARCHAR(10) not NULL,
    CreatedAt timestamp not null,
    TypeOfMessage VARCHAR(10) not null
);

CREATE TABLE conversation
(
    id VARCHAR(20) not null UNIQUE,
    title varchar(40),
    creatorID VARCHAR(10),
    created_at timestamp
);

CREATE TABLE participants
(
    id VARCHAR(20) not null UNIQUE,
    conversation_id VARCHAR(20),
    users_id VARCHAR(20)
);