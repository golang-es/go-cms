--
-- Trigger que crea los registros en module_role desde la tabla modules
-- Author: Alexys Lozada
-- Description: Cuando se registra un nuevo elemento en modules, se pobla la tabla module_role con todos los roles actuales.
--

-- Función trigger
CREATE FUNCTION trigger_modules_module_role()
RETURNS trigger AS
$BODY$
BEGIN
    INSERT INTO module_role (module_id, role_id)
        SELECT NEW.id, roles.id
        FROM roles;

    RETURN NEW;
END;
$BODY$
LANGUAGE plpgsql;

-- Comentario de la función
COMMENT ON FUNCTION trigger_modules_module_role() IS 'Cuando se registra un nuevo elemento en modules, se pobla la tabla module_role con todos los roles actuales';

-- Ejecución del trigger
CREATE TRIGGER trigger_modules_module_role AFTER INSERT ON modules FOR EACH ROW EXECUTE PROCEDURE trigger_modules_module_role();

-- Completed
