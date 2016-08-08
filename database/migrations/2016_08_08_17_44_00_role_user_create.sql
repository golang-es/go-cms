--
-- Creación de la tabla role_user
-- Author: Alexys Lozada
-- Description: Creación de la tabla role_user.
--

CREATE TABLE role_user (
    id SERIAL NOT NULL,
    role_id SMALLINT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

-- Completed
