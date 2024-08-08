# Task API Documentation

## Get All Tasks

### Request

`GET /tasks`

### Response

**Status Code:** 200 OK

**Response Body:**

```json
[
  {
    "id": 1,
    "title": "Task 1",
    "done": false
  },
  {
    "id": 2,
    "title": "Task 2",
    "done": true
  }
]
```

## Get Task by ID

### Request

`GET /tasks/:id`

**Path Parameters:** - `id` (integer, required): The ID of the task to retrieve.

### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "id": 1,
  "title": "Task 1",
  "done": false
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "invalid id"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "task not found"
}
```

## Update Task

### Request

`PUT /tasks/:id`

**Path Parameters:** - `id` (integer, required): The ID of the task to retrieve.

**Request Body:**

```json
{
  "title": "Updated Task 1",
  "done": true
}
```

### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "id": 1,
  "title": "Updated Task 1",
  "done": true
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "invalid id"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "task not found"
}
```

## Delete Task

### Request

`DELETE /tasks/:id`

**Path Parameters:** - `id` (integer, required): The ID of the task to retrieve.

### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "message": "task deleted"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "invalid id"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "task not found"
}
```

## Create Task

### Request

`POST /tasks`

**Request Body:**

```json
{
  "title": "New Task",
  "done": false
}
```

### Response

**Status Code:** 201 Created

**Response Body:**

```json
{
  "id": 3,
  "title": "New Task",
  "done": false
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "Invalid request body"
}
```
