-- name: GetProductsParamatersValue :one
SELECT
    ppv.id,
    ppv.parameter,
    ppv.product,
    ppv.selectable_values,
    ppv.text_value,
    ppv.created_at,
    par.name AS parameter_name,
    p.name AS product_name
FROM
    products_parameters_values AS ppv
    JOIN parameters AS par ON ppv.parameter = par.id
    JOIN products AS p ON ppv.product = p.id
WHERE
    p.id = $1
LIMIT
    1;

-- name: GetProductsParamatersValues :many
SELECT
    ppv.id,
    ppv.parameter,
    ppv.product,
    ppv.selectable_values,
    ppv.boolean_value,
    ppv.text_value,
    ppv.created_at,
    par.name AS parameter_name,
    p.name AS product_name
FROM
    products_parameters_values AS ppv
    JOIN parameters AS par ON ppv.parameter = par.id
    JOIN products AS p ON ppv.product = p.id
ORDER BY
    p.created_at
LIMIT
    $1 OFFSET $2;

-- name: CreateProductsParamatersValue :one
INSERT INTO
    products_parameters_values(
        parameter,
        product,
        selectable_values,
        boolean_value,
        text_value
    )
VALUES
    ($1, $2, $3, $4, $5)
RETURNING
    *;

-- name: UpdateProductsParamatersValue :one
UPDATE
    products_parameters_values
SET
    parameter = $2,
    product = $3,
    selectable_values = $4,
    boolean_value = $5,
    text_value = $6
WHERE
    id = $1
RETURNING
    *;

-- name: DeleteProductsParamatersValue :exec
DELETE FROM
    products_parameters_values
WHERE
    id = $1;
