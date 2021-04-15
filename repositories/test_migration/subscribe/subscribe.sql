TRUNCATE TABLE users, friends, subscribe, blocks;

INSERT INTO users(id, email) VALUES (1, 'andy@example.com');
INSERT INTO users(id, email) VALUES (2, 'john@example.com');
INSERT INTO users(id, email) VALUES (3, 'kate@example.com');
INSERT INTO users(id, email) VALUES (4, 'lisa@example.com');

INSERT INTO subscribe(requestor, target) VALUES (1, 2);
