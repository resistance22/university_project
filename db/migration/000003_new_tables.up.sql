CREATE TYPE app_user_roles AS ENUM ('admin', 'owner', 'sales_manager', 'warehouse_manager', 'sales_agent');
ALTER TABLE app_user ADD COLUMN role app_user_roles NOT NULL; 