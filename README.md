An event-booking backend system with basic authentication and authorization(Implemented in Go Lang).
----------------------------------------------------------------------------------------------

Dependencies:
Execute the below commands to download the packages:

go get -u github.com/gin-gonic/gin

go get -u github.com/golang-jwt/jwt/v5

## API Endpoints (Postman collection of APIs is provided in the code.)

### Authentication

#### Login:
-------------------

- **URL:** `/login`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
- **Body Parameters:**
 ```json
{
  "email": "test@abc.com",
  "password": "test@1123"
}
```

#### SignUp:
---------------------

- **URL:** `/signup`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
- **Body Parameters:**
```json
{
  "email": "test@abc.com",
  "password": "test@1123"
}
```

### FETCH:

#### Fetch all events:
--------------------------

- **URL:** `/events`
- **Method:** `GET`
- **Headers:** 
  - `Content-Type: application/json`

#### Fetch event by Id:
----------------------------------

- **URL:** `/events/{id}`
- **Method:** `GET`
- **Headers:** 
  - `Content-Type: application/json`


### CREATE:

#### Create Event:
---------------------------------------

- **URL:** `/create/event`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: {Token received at the time of login} `
- **Body Parameters:**
```json
{
    "name": "Seminar",
    "description": "Tech Event",
    "location": "Delhi",
    "dateTime":"2024-08-04T11:38:00.000Z"
}
```

### UPDATE:

#### Update Event By Id:
---------------------------------
- **URL:** `/event/update/{id}`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: {Token received at the time of login} `
- **Params:**
    -` id: {int}` 
- **Body Parameters:**
```json
{
    "name": "Seminar",
    "description": "Tech Event",
    "location": "Delhi",
    "dateTime":"2024-08-04T11:38:00.000Z"
}
```

### DELETE:

#### Delete Event By Id:
------------------------------

- **URL:** `/event/delete/{id}`
- **Method:** `DELETE`
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: {Token received at the time of login} `
- **Params:**
    -` id: {int}`


### REGISTRATION & CANCELLATION:

#### Register for an event:
----------------------------------

- **URL:** `/events/:id/register`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: {Token received at the time of login} `
- **Params:**
    -` id: {event id of type int}` 

#### De-register for an event:
------------------------------------

- **URL:** `/events/:id/register/cancel`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: {Token received at the time of login} `
- **Params:**
    -` id: {event id of type int}`

> **Note:** Changes/updates can be done.
