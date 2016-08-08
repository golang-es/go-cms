--
-- Restricciones y Relaciones de la tabla users
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla users.
--

ALTER TABLE users ADD CONSTRAINT users_id_primary PRIMARY KEY (id);
ALTER TABLE users ADD CONSTRAINT users_email_unique UNIQUE (email);

-- Completed
