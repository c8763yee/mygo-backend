#!/bin/bash
mariadb -u root -p${ROOT_PASSWORD} <<EOF
CREATE DATABASE IF NOT EXISTS ${DB_NAME};
CREATE USER IF NOT EXISTS '${USER}'@'localhost' IDENTIFIED BY '${PASSWORD}';
GRANT ALL PRIVILEGES ON ${DB_NAME}.* TO '${DB_USER}'@'localhost';
FLUSH PRIVILEGES;
EOF

>&2 echo -e "MariaDB database initialized with database: '${DB_NAME}' and user: '${DB_USER}'"
