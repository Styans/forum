CREATE TABLE IF NOT EXISTS comments (
    id          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    userName    TEXT,
    user_id     TEXT,
    comment     TEXT,
    post_id     int,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);