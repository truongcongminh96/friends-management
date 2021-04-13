CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    email text NOT NULL,
    CONSTRAINT unique_user_email UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS friends
(
    id SERIAL PRIMARY KEY,
    user1 int NOT NULL,
    user2 int NOT NULL,
    FOREIGN KEY (user1) REFERENCES users(id),
    FOREIGN KEY (user2) REFERENCES users(id),
    CONSTRAINT unique_friends_user1_user2 UNIQUE (user1, user2)
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    id SERIAL PRIMARY KEY,
    requestor int NOT NULL,
    target int NOT NULL,
    FOREIGN KEY (requestor) REFERENCES users(id),
    FOREIGN KEY (target) REFERENCES users(id),
    CONSTRAINT unique_subscriptions_requestor_target UNIQUE (requestor, target)
);

CREATE TABLE IF NOT EXISTS blockings
(
    id SERIAL PRIMARY KEY,
    requestor int NOT NULL,
    target int NOT NULL,
    FOREIGN KEY (requestor) REFERENCES users(id),
    FOREIGN KEY (target) REFERENCES users(id),
    CONSTRAINT unique_blockings_requestor_target UNIQUE (requestor, target)
);
