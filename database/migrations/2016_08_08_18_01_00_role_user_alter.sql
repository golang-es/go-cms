--
-- Restricciones y Relaciones de la tabla users
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla users.
--

ALTER TABLE role_user ADD CONSTRAINT role_user_id_primary PRIMARY KEY (id);
ALTER TABLE role_user ADD CONSTRAINT role_user_role_id_user_id_unique UNIQUE (role_id, user_id);
ALTER TABLE role_user ADD CONSTRAINT role_user_role_id_foreign FOREIGN KEY (role_id) REFERENCES roles (id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE role_user ADD CONSTRAINT role_user_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT;

-- Completed
