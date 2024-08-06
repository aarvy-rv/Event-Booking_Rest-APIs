An Event-Booking backend system implemented in Go Lang with authentication and authorization.
----------------------------------------------------------------------------------------------

APIs:

->POST    "/signup"   
  Body Format for Signup(JSON):
    {
    "email":"test@abc.com",
    "password":"test@123"
    }

->POST    "/login"
  Body Format for Login(JSON):
    {
    "email":"test@abc.com",
    "password":"test@123"
    }

->GET    "/events"
->GET    "/event/:id"

Belows APIs need an "Authothorization" header whose value will be the token recieved in the login api when user logs in.
->POST    "/create/event"
  Body:
  {
    "name" : "Testingggg",
    "description": "A test event",
    "location": "Delhi",
    "dateTime":"2024-08-04T11:38:00.000Z"
}

->PUT     "/event/update/:id"
  Body:
    {
    "name" : "UpdatedTesttt",
    "description": "An updated test event",
    "location": "Local lcation",
    "dateTime":"2024-08-04T11:38:00.000Z"
    }

->DELETE  "/event/delete/:id"
->POST    "/events/:id/register"
->DELETE  "/events/:id/register/cancel"
