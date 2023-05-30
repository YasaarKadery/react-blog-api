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
Once the server is running you will probably need to populate the database with data. You can use curl to send POST requests to the /posts endpoint.
For example:

```
curl -X POST -H "Content-Type: application/json" -d '{"title": "First Post", "content": "This is the first post.", "image_src": "https://example.com/images/first.jpg"}' http://localhost:8080/posts

curl -X POST -H "Content-Type: application/json" -d '{"title": "Second Post", "content": "This is the second post.", "image_src": "https://example.com/images/second.jpg"}' http://localhost:8080/posts

curl -X POST -H "Content-Type: application/json" -d '{"title": "Third Post", "content": "This is the third post.", "image_src": "https://example.com/images/third.jpg"}' http://localhost:8080/posts

```
