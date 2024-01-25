CREATE TABLE IF NOT EXISTS comments (
    id          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    userName    TEXT NOT NULL,
    user_id     TEXT NOT NULL,
    comment     TEXT NOT NULL,
    post_id     int NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);