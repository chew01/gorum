CREATE TABLE IF NOT EXISTS users (
    ID INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    ID INTEGER PRIMARY KEY,
    creator TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    createdAt TEXT NOT NULL,
    updatedAt TEXT NOT NULL,
    FOREIGN KEY (creator) REFERENCES users (name)
);

CREATE TABLE IF NOT EXISTS comments (
    ID INTEGER PRIMARY KEY,
    post INTEGER NOT NULL,
    creator TEXT NOT NULL,
    content TEXT NOT NULL,
    createdAt TEXT NOT NULL,
    updatedAt TEXT NOT NULL,
    FOREIGN KEY (post) REFERENCES posts (ID) ON DELETE CASCADE,
    FOREIGN KEY (creator) REFERENCES users (name)
);