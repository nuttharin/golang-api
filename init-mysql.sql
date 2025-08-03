-- Create database

CREATE DATABASE IF NOT EXISTS my_db;
CREATE USER IF NOT EXISTS 'admin'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON my_db.* TO 'admin'@'%';
FLUSH PRIVILEGES;

-- USE my_db;

-- -- Create table
-- CREATE TABLE IF NOT EXISTS user (
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     name VARCHAR(255) NOT NULL,
--     email VARCHAR(255) NOT NULL UNIQUE
-- );

-- -- Seed data
-- INSERT INTO
--     users (name, email)
-- VALUES
--     ('user1', 'user1@test.com'),
--     ('user2', 'user2@test.com');