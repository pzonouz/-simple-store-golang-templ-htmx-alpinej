-- name: GetParameter :one
SELECT
   p.id,
   p.name,
   p.parameters_group,
   p.parameter_type,
   p.selectable_options,
   p.created_at,
   pg.name as parameters_group_name
FROM
    parameters as p LEFT JOIN
    parameters_groups as pg ON p.parameters_group = pg.id 
WHERE
    p.id = $1;

-- name: GetParameters :many
SELECT
   p.id,
   p.name,
   p.parameters_group,
   p.parameter_type,
   p.selectable_options,
   p.created_at,
   pg.name as parameters_group_name
FROM
    parameters as p LEFT JOIN
    parameters_groups as pg ON p.parameters_group = pg.id 
ORDER BY
    p.created_at
LIMIT $1 OFFSET $2;

-- name: GetParametersByParametersGroup :many
SELECT
   p.id,
   p.name,
   p.parameters_group,
   p.parameter_type,
   p.selectable_options,
   p.created_at,
   pg.name as parameters_group_name
FROM
    parameters as p LEFT JOIN
    parameters_groups as pg ON p.parameters_group = pg.id 
WHERE
    pg.id = $1
ORDER BY
    p.created_at;

-- name: CreateParameter :one
INSERT INTO parameters(name, parameters_group, parameter_type, selectable_options)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: UpdateParameter :one
UPDATE
    parameters
SET
    name = $2,
    parameters_group = $3,
    parameter_type = $4,
    selectable_options =$5
WHERE
    id = $1
RETURNING
    *;

-- name: DeleteParameter :exec
DELETE FROM parameters
WHERE id = $1;

