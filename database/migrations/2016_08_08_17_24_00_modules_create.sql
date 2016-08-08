--
-- Creación de la tabla modules
-- Author: Alexys Lozada
-- Description: Creación de la tabla modules.
--

CREATE TABLE modules (
    id SMALLSERIAL NOT NULL,
    name VARCHAR(30) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

-- Completed
