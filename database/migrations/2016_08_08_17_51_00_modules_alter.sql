--
-- Restricciones y Relaciones de la tabla modules
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla modules.
--

ALTER TABLE modules ADD CONSTRAINT modules_id_primary PRIMARY KEY (id);
ALTER TABLE modules ADD CONSTRAINT modules_name_unique UNIQUE (name);

-- Completed
