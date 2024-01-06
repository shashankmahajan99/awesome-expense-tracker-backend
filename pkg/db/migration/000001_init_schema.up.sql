-- init_schema.up.sql

-- Users table
CREATE TABLE IF NOT EXISTS Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
    -- Add other fields as necessary
);

-- Expenses table
CREATE TABLE IF NOT EXISTS Expenses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    amount DECIMAL(10, 2),
    description VARCHAR(255),
    -- Add other fields as necessary
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- Reports table
CREATE TABLE IF NOT EXISTS Reports (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    title VARCHAR(255),
    -- Add other fields as necessary
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- Profiles table
CREATE TABLE IF NOT EXISTS Profiles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    bio TEXT,
    -- Add other fields as necessary
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- Settings table
CREATE TABLE IF NOT EXISTS Settings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    theme VARCHAR(255),
    currency VARCHAR(255),
    -- Add other fields as necessary
    FOREIGN KEY (user_id) REFERENCES Users(id)
);