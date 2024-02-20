CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    position INTEGER NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK(length(name) <= 255)
);

CREATE TABLE boards (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL ,
    description TEXT,
    position INTEGER NOT NULL,
    hidden INTEGER NOT NULL DEFAULT 0,
    locked INTEGER NOT NULL DEFAULT 0,
    redirect TEXT,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    category_id INTEGER NOT NULL,
    parent_board_id INTEGER,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (parent_board_id) REFERENCES boards(id),
    UNIQUE(category_id, id),
    CHECK(length(name) <= 255),
    CHECK(length(description) <= 255)
);

CREATE TABLE threads (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    subject TEXT NOT NULL,
    slug TEXT NOT NULL,
    body TEXT NOT NULL,
    pinned INTEGER NOT NULL DEFAULT 0,
    locked INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT,
    updated_reason TEXT,
    board_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (board_id) REFERENCES boards(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE(board_id, id),
    CHECK(length(subject) <= 255)
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_number INTEGER NOT NULL,
    body TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT,
    updated_reason TEXT,
    thread_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    reply_to_post_id INTEGER,
    FOREIGN KEY (thread_id) REFERENCES threads(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (reply_to_post_id) REFERENCES posts(id),
    UNIQUE(thread_id, id)
);

CREATE INDEX posts_thread_id_post_number ON posts (thread_id, post_number);

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK(length(username) <= 50)
);