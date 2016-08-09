--
-- Restricciones y Relaciones de la tabla metadatas
-- Author: Alexys Lozada
-- Description: Restricciones y Relaciones de la tabla metadatas.
--

ALTER TABLE metadatas ADD CONSTRAINT metadatas_id_primary PRIMARY KEY (id);
ALTER TABLE metadatas ADD CONSTRAINT metadatas_post_id_foreign FOREIGN KEY (post_id) REFERENCES posts (id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE metadatas ADD CONSTRAINT metadatas_metadata_type_id_foreign FOREIGN KEY (metadata_type_id) REFERENCES metadata_types (id) ON UPDATE RESTRICT ON DELETE RESTRICT;

-- Completed
