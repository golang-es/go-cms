--
-- Trigger que crea los registros en module_role desde la tabla roles
-- Author: Alexys Lozada
-- Description: Cuando se registra un nuevo elemento en roles, se pobla la tabla module_role con todos los modulos actuales.
--

-- Función trigger
CREATE FUNCTION trigger_roles_module_role()
RETURNS trigger AS
$BODY$
BEGIN
    INSERT INTO module_role (module_id, role_id)
        SELECT NEW.id, modules.id
        FROM modules;

    RETURN NEW;
END;
$BODY$
LANGUAGE plpgsql;

-- Comentario de la función
COMMENT ON FUNCTION trigger_roles_module_role() IS 'Cuando se registra un nuevo elemento en roles, se pobla la tabla module_role con todos los modulos actuales';

-- Ejecución del trigger
CREATE TRIGGER trigger_roles_module_role AFTER INSERT ON roles FOR EACH ROW EXECUTE PROCEDURE trigger_roles_module_role();

-- Completed
