--
-- Restricciones y Relaciones de la tabla module_role
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla module_role.
--

ALTER TABLE module_role ADD CONSTRAINT module_role_id_primary PRIMARY KEY (id);
ALTER TABLE module_role ADD CONSTRAINT module_role_module_id_role_id_unique UNIQUE (module_id, role_id);
ALTER TABLE module_role ADD CONSTRAINT module_role_module_id_foreign FOREIGN KEY (module_id) REFERENCES modules (id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE module_role ADD CONSTRAINT module_role_role_id_foreign FOREIGN KEY (role_id) REFERENCES roles (id) ON UPDATE RESTRICT ON DELETE RESTRICT;

-- Completed
