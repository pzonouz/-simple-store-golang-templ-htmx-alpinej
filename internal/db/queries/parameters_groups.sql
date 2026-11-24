
-- name: GetParametersGroup :one
SELECT
   p.id,
   p.name,
   p.category,
   p.created_at,
   c.name as category_name
FROM
    parameters_groups as p JOIN
    categories as c ON p.category = c.id 
WHERE
    p.id = $1
LIMIT 1;

-- name: GetParametersGroups :many
SELECT
   p.id,
   p.name,
   p.category,
   p.created_at,
   c.name as category_name
FROM
    parameters_groups as p JOIN
    categories as c ON p.category = c.id 
ORDER BY
    p.created_at
LIMIT $1 OFFSET $2;

-- name: CreateParametersGroup :one
INSERT INTO parameters_groups(name, category)
    VALUES ($1, $2)
RETURNING
    *;

-- name: UpdateParametersGroup :one
UPDATE
    parameters_groups
SET
    name = $2,
    category = $3
WHERE
    id = $1
RETURNING
    *;

-- name: DeleteParametersGroup :exec
DELETE FROM parameters_groups
WHERE id = $1;

-- name: GetParametersGroupsByCategoryAndAncestors :many
WITH recursive ancestors AS(
SELECT c.id,c.name,c.parent FROM categories c WHERE c.id=$1
UNION ALL
SELECT c.id,c.name,c.parent FROM categories c INNER JOIN ancestors an ON c.id=an.parent 
)
SELECT pg.id,pg.name,pg.category,pg.created_at FROM ancestors an JOIN parameters_groups pg ON an.id=pg.category;
