ALTER TABLE IF EXISTS users ADD CONSTRAINT unique_email UNIQUE (email);
ALTER TABLE IF EXISTS users ADD CONSTRAINT unique_username UNIQUE (username);
