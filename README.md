# go-cms
Sistema manejador de contenidos (CMS) creado con el lenguaje de programación Go (golang).

## Convenciones para la base de datos

### El nombre de las tablas
El nombre de las tablas deberán estar en inglés, plural y en minúsculas.

Si el nombre de una tabla tiene más de una palabra, éstas deberán estar separadas por un guión bajo (`_`).

### El primary key
El campo primary key de toda tabla deberá llamarse `id`.

### Los campos created_at y updated_at
Cada tabla deberá tener los campos timestamp `created_at` y `updated_at` que almacenarán la fecha de creación y modificación de cada registro.

### El nombre de los campos
El nombre de los campos deberán estar en inglés.

### Las tablas pivote
Las tablas pivote son aquellas que representan una relación de muchos a muchos entre dos tablas.

El nombre de una tabla pivote deberá estar formada por el nombre en singular de las tablas involucradas separadas por un guión bajo, en orden alfabético y en minúsculas. Por ejemplo, si tenemos dos tablas `users` y `roles`, la tabla pivote se llamará `role_user`.

La tabla pivote deberá tener sus propios campos `id`, `created_at` y `updated_at`.

### Nombre de las relaciones
En nombre de los campos que sean llaves foráneas deberán formarse con el nombre en singular de la otra tabla seguido de `_id`.

Ejemplo: Si la tabla `posts` tiene una llave foránea que apunta a la tabla `users`, el nombre del campo en la tabla `posts` será `user_id`.

El nombre de las relaciones e índices deberá estar compuesto por el nombre de la tabla actual seguido de un guión bajo (`_`), el nombre del campo de la tabla actual, un guión bajo y el tipo de relación/índice (foreign, unique, primary).

Ejemplo: `posts_user_id_foreign`

### Las migraciones
Cada tabla deberá tener su propio archivo `.sql` nombrado con el formato `YYY_MM_DD_hh_mm_ss_nombre_de_archivo.sql`. El nombre deberá estar en minúsculas usando guiones bajos para separar las palabras en caso de haber más de una.

#### Restricciones de tablas
Todas las relaciones y restricciones deberán ser creadas con el comando ALTER y en su propio archivo SQL en la medida de lo posible.

#### Los trigger
Los triggers deberán ser nombrados con el formato `trigger_tabla_origen_tabla_destino`.
