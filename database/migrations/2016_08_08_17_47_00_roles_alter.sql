--
-- Restricciones y Relaciones de la tabla roles
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla roles.
--

ALTER TABLE roles ADD CONSTRAINT roles_id_primary PRIMARY KEY (id);
ALTER TABLE roles ADD CONSTRAINT roles_id_unique UNIQUE (name);

-- Completed
