-- Drop FKs first
ALTER TABLE "products_parameters_values" DROP CONSTRAINT IF EXISTS "products_parameters_values_product_fkey";
ALTER TABLE "products_parameters_values" DROP CONSTRAINT IF EXISTS "products_parameters_values_parameter_fkey";
ALTER TABLE "parameters"                 DROP CONSTRAINT IF EXISTS "parameters_parameters_group_fkey";
ALTER TABLE "parameters_groups"          DROP CONSTRAINT IF EXISTS "parameters_groups_category_fkey";

-- Drop unique constraint added on values table
ALTER TABLE "products_parameters_values"
    DROP CONSTRAINT IF EXISTS "products_parameters_values_parameter_product_key";

-- Drop tables (order: child → parent)
DROP TABLE IF EXISTS "products_parameters_values";
DROP TABLE IF EXISTS "parameters";
DROP TABLE IF EXISTS "parameters_groups";

-- Drop enum
DROP TYPE IF EXISTS parameter_type;
