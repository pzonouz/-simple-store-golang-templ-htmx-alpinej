CREATE TYPE parameter_type AS ENUM ('Boolean', 'Selectable', 'Text');

CREATE TABLE "parameters_groups" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" varchar UNIQUE NOT NULL,
    "category" uuid,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "parameters" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" varchar UNIQUE NOT NULL,
    "parameters_group" uuid,
    "parameter_type" parameter_type NOT NULL DEFAULT 'Text',
    "selectable_options" text [] DEFAULT '{}',  -- fixed typo
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "products_parameters_values" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "parameter" uuid,
    "product" uuid,
    "selectable_values" text [],
    "boolean_value" boolean,
    "text_value" text,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    UNIQUE ("parameter", "product") -- one value per pair
);

ALTER TABLE
    "parameters_groups"
ADD
    FOREIGN KEY ("category") REFERENCES "categories"("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE
    "parameters"
ADD
    FOREIGN KEY ("parameters_group") REFERENCES "parameters_groups"("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE
    "products_parameters_values"
ADD
    FOREIGN KEY ("parameter") REFERENCES "parameters"("id") ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE
    "products_parameters_values"
ADD
    FOREIGN KEY ("product") REFERENCES "products"("id") ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE;
