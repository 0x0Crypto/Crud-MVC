# Crud-MVC

This is an example of a crud API using Gin Gonic Framework.

Running with docker

```
docker build -t mvc:latest . && docker run --publish 8080:8080 mvc
```

# API Docs

## Endpoint `/users`

The `/users` endpoint allows you to create, update, and delete users.

### 1. Create User

```
http://localhost:8080/users
```

- **HTTP Method:** POST

- **Endpoint URL:** `/users`

- **Description:** Creates a new user.

- **Request Body:**
  
  ```json
  {
      "name": "John",
      "age": 23,
  }
  ```

- **Alternative Status Codes:**

- **400 Bad Request:** If the request body is missing required information.

### 2. Update User

- **HTTP Method:** PUT

- **Endpoint URL:** `/users`

- **Description:** Updates an existing user's information.

- **Request Body:**

  ```json
  {
      "id": 1
      "name": "New Name",
      "age": "New Age"
  }
  ```

- **Alternative Status Codes:**

- **400 Bad Request:** If the request body is missing required information or the user ID is not provided.

- **404 Not Found:** If the user with the provided ID is not found.

### 3. Delete User

- **HTTP Method:** DELETE

- **Endpoint URL:** `/users/`

- **Description:** Deletes an existing user.

- **Path Parameters:**
  
  - `userId` (string) - The ID of the user to be deleted.

- **Example Request:**

  ```http
  DELETE /users/12345 HTTP/1.1
  Host: 127.0.0.1:8080
  ```

- **Alternative Status Codes:**

- **404 Not Found:** If the user with the provided ID is not found.
