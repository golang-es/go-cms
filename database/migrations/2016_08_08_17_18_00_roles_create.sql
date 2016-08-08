--
-- Creación de la tabla roles
-- Author: Alexys Lozada
-- Description: Creación de la tabla roles.
--

CREATE TABLE roles (
    id SMALLSERIAL NOT NULL,
    name VARCHAR(25) NOT NULL,
    active boolean DEFAULT true NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

-- Completed
