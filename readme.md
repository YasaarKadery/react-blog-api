# Usage

To create the MySQL database run the following commands:
Login as the root user

```
mysql -u root -p
```

Create the database

```
CREATE DATABASE blog_db;

```

Select the database to use

```
USE blog_db;

```

Create the table inside the selected databases

```
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    image_src TEXT
);
```

Now that the database is set up, we can run the server.
CD into the directory and run:

```
go run .
```

Make sure to have environment variables DB USER and DB_PASSWORD to store the db credentials.
