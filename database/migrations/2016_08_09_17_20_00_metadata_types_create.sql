--
-- Creación de la tabla metadata_types
-- Author: Alexys Lozada
-- Description: Creación de la tabla metadata_types, que son los tipos de meta que aparece en el head de html.
--

CREATE TABLE metadata_types (
    id SMALLSERIAL NOT NULL,
    type VARCHAR(50) NOT NULL,
    max_length SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Completed
