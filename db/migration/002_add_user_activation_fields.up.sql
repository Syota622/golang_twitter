ALTER TABLE users
ADD COLUMN is_active BOOLEAN DEFAULT FALSE,
ADD COLUMN activation_token VARCHAR(255);