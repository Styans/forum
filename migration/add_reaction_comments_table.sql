CREATE TABLE IF NOT EXISTS reactions_comments (
    id          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    comment_id  TEXT,
    user_id     TEXT,
    like        boolean,
    dislike     boolean,  
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);