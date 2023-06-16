-- Prepares the development database on the MySQL server.
-- Create the development and test databases if they don't already exist.
CREATE DATABASE IF NOT EXISTS wwapp_dev_db;
CREATE DATABASE IF NOT EXISTS wwapp_dev_db_test;

-- Add new user and set password
CREATE USER IF NOT EXISTS 'wwapp_dev'@'localhost' IDENTIFIED BY 'WWApp-dev-pwd-0';

-- Grant priviges to user on the new databases
GRANT ALL PRIVILEGES ON wwapp_dev_db.* TO 'wwapp_dev'@'localhost';
GRANT ALL PRIVILEGES ON wwapp_dev_db_test.* TO 'wwapp_dev'@'localhost';

-- Grant select privilege to user on performance_schema
GRANT SELECT ON performance_schema.* TO 'wwapp_dev'@'localhost';

FLUSH PRIVILEGES;
