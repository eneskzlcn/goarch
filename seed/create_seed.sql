CREATE TABLE IF NOT EXISTS users(
    id VARCHAR NOT NULL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    username VARCHAR NOT NULL UNIQUE,
    post_count INTEGER NOT NULL DEFAULT 0 CHECK ( post_count >= 0 ),
    follower_count INTEGER NOT NULL DEFAULT 0 CHECK (follower_count >= 0),
    followed_count INTEGER NOT NULL DEFAULT 0 CHECK ( followed_count >= 0 ),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS posts(
    id VARCHAR NOT NULL PRIMARY KEY,
    user_id VARCHAR NOT NULL REFERENCES users ON DELETE CASCADE ON UPDATE CASCADE,
    content TEXT NOT NULL,
    comment_count INT DEFAULT 0 CHECK ( comment_count >= 0 ),
    like_count INT DEFAULT 0 CHECK ( like_count >= 0 ),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);