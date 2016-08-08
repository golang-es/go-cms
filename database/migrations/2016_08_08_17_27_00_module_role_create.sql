--
-- Creación de la tabla module_role
-- Author: Alexys Lozada
-- Description: Creación de la tabla module_role.
--

CREATE TABLE module_role (
    id SERIAL NOT NULL,
    module_id SMALLINT NOT NULL,
    role_id SMALLINT NOT NULL,
    append boolean DEFAULT false NOT NULL,
    modify boolean DEFAULT false NOT NULL,
    remove boolean DEFAULT false NOT NULL,
    read boolean DEFAULT false NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

-- Completed
