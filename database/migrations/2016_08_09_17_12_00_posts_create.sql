--
-- Creación de la tabla posts
-- Author: Alexys Lozada
-- Description: Creación de la tabla posts.
--

CREATE TABLE posts (
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    title VARCHAR(140) NOT NULL,
    slug VARCHAR(140) NOT NULL,
    content TEXT,
    published_at TIMESTAMP,
    poster VARCHAR(500),
    banner VARCHAR(500),
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL
);

-- Completed
