TRUNCATE TABLE users, friends, subscribe, blocks;

INSERT INTO users(id, email) VALUES (1, 'andy@example.com');
INSERT INTO users(id, email) VALUES (2, 'john@example.com');
INSERT INTO users(id, email) VALUES (3, 'kate@example.com');
INSERT INTO users(id, email) VALUES (4, 'lisa@example.com');

INSERT INTO friends(user1, user2) VALUES (1, 2);
INSERT INTO friends(user1, user2) VALUES (2, 3);