CREATE TABLE IF NOT EXISTS posts (
    id          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title       TEXT NOT NULL,
    content     TEXT NOT NULL,
    author_id   INTEGER NOT NULL,
    authorname  TEXT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NOT NULL,
    -- version     INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS images (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id     INTEGER,
    image_path TEXT UNIQUE,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
)