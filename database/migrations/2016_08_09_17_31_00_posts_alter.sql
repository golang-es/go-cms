--
-- Restricciones y Relaciones de la tabla post
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla post.
--

ALTER TABLE posts ADD CONSTRAINT posts_id_primary PRIMARY KEY (id);
ALTER TABLE posts ADD CONSTRAINT posts_title_unique UNIQUE (title);
ALTER TABLE posts ADD CONSTRAINT posts_slug_unique UNIQUE (slug);
ALTER TABLE posts ADD CONSTRAINT posts_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT;

-- Completed
