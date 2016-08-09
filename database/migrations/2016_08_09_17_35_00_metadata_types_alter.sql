--
-- Restricciones y Relaciones de la tabla metadata_types
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla metadata_types.
--

ALTER TABLE metadata_types ADD CONSTRAINT metadata_types_id_primary PRIMARY KEY (id);
ALTER TABLE metadata_types ADD CONSTRAINT metadata_types_type_unique UNIQUE (type);

-- Completed
