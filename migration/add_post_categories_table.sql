CREATE TABLE IF NOT EXISTS categories (
    id     INT PRIMARY KEY,
    category_name   VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS PostCategories (
    post_id INT,
    category_id INT,
    PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);