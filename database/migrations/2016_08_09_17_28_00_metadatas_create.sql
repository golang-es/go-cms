--
-- Creación de la tabla metadatas
-- Author: Alexys Lozada
-- Description: Creación de la tabla metadatas, que son los contenidos de tipos de meta que aparece en el head de html.
--

CREATE TABLE metadatas (
    id SERIAL NOT NULL,
    post_id INTEGER NOT NULL,
    metadata_type_id SMALLINT NOT NULL,
    content VARCHAR(500) NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL
);

-- Completed
