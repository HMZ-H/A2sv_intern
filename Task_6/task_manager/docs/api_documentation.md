## API Documentation

### Task Management Endpoints

#### Create a New Task
- **Endpoint:** `POST /tasks/`
- **Description:** Creates a new task.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Request Body:** 
  ```json
  {
    "title": "Task Title",
    "description": "Task Description",
    "dueDate": "2024-08-15T00:00:00Z"
  }
  ```
- **Response:**
  - **Status:** Created
  - **Body:**
    ```json
    {
      "id": "task_id",
      "title": "Task Title",
      "description": "Task Description",
      "dueDate": "2024-08-15T00:00:00Z",
      "status": "pending"
    }
    ```

#### Get All Tasks
- **Endpoint:** `GET /tasks/`
- **Description:** Retrieves all tasks.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Response:**
  - **Status:** 200 OK
  - **Body:**
    ```json
    [
      {
        "id": "task_id",
        "title": "Task Title",
        "description": "Task Description",
        "dueDate": "2024-08-15T00:00:00Z",
        "status": "pending"
      },
      ...
    ]
    ```

#### Get Task by ID
- **Endpoint:** `GET /tasks/:id`
- **Description:** Retrieves a task by its ID.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Response:**
  - **Status:** 200 OK
  - **Body:**
    ```json
    {
      "id": "task_id",
      "title": "Task Title",
      "description": "Task Description",
      "dueDate": "2024-08-15T00:00:00Z",
      "status": "pending"
    }
    ```

#### Update Task by ID
- **Endpoint:** `PUT /tasks/:id`
- **Description:** Updates a task by its ID.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Request Body:** 
  ```json
  {
    "title": "Updated Task Title",
    "description": "Updated Task Description",
    "dueDate": "2024-08-20T00:00:00Z",
    "status": "completed"
  }
  ```
- **Response:**
  - **Status:** 200 OK
  - **Body:**
    ```json
    {
      "id": "task_id",
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "dueDate": "2024-08-20T00:00:00Z",
      "status": "completed"
    }
    ```

#### Delete Task by ID
- **Endpoint:** `DELETE /tasks/:id`
- **Description:** Deletes a task by its ID.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Response:**
  - **Status:** 204 No Content

### User Management Endpoints

#### Register a New User
- **Endpoint:** `POST /register`
- **Description:** Registers a new user.
- **Request Body:** 
  ```json
  {
    "username": "newuser",
    "password": "password"
  }
  ```
- **Response:**
  - **Status:** 201 Created
  - **Body:**
    ```json
    {
      "id": "user_id",
      "username": "newuser"
    }
    ```

#### User Login
- **Endpoint:** `POST /login`
- **Description:** Authenticates a user and generates a JWT token.
- **Request Body:** 
  ```json
  {
    "username": "existinguser",
    "password": "password"
  }
  ```
- **Response:**
  - **Status:** 200 OK
  - **Body:**
    ```json
    {
      "token": "jwt_token"
    }
    ```

#### Get All Users (Admin Only)
- **Endpoint:** `GET /users/`
- **Description:** Retrieves all users. Requires admin role.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Response:**
  - **Status:** 200 OK
  - **Body:**
    ```json
    [
      {
        "id": "user_id",
        "username": "user"
      },
      ...
    ]
    ```

#### Get User by ID
- **Endpoint:** `GET /users/:id`
- **Description:** Retrieves a user by their ID.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Response:**
  - **Status:** 200 OK
  - **Body:**
    ```json
    {
      "id": "user_id",
      "username": "user"
    }
    ```

#### Update User by ID
- **Endpoint:** `PUT /user/:id`
- **Description:** Updates a user by their ID.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Request Body:** 
  ```json
  {
    "username": "updateduser",
    "password": "newpassword"
  }
  ```
- **Response:**
  - **Status:** 200 OK
  - **Body:**
    ```json
    {
      "id": "user_id",
      "username": "updateduser"
    }
    ```

#### Delete User by ID (Admin Only)
- **Endpoint:** `DELETE /users/:id`
- **Description:** Deletes a user by their ID. Requires admin role.
- **Request Headers:** 
  - `Authorization: Bearer <JWT Token>`
- **Response:**
  - **Status:** 204 No Content