An event-booking backend system with authentication and authorization was implemented in Go Lang.
----------------------------------------------------------------------------------------------

APIs:

|POST:    "/signup"   
  Body Format for Signup(JSON):
    {
    "email":"test@abc.com",
    "password":"test@123"
    }
    
| POST:    "/login"
  Body Format for Login(JSON):
    {
    "email":"test@abc.com",
    "password":"test@123"
    }

| GET:    "/events"
| GET:    "/event/:id"

| Belows APIs need an "Authorization" header whose value will be the token received in the login API when the user logs in.
 | POST:    "/create/event"
  Body:
  {
    "name" : "Testingggg",
    "description": "A test event",
    "location": "Delhi",
    "dateTime":"2024-08-04T11:38:00.000Z"
}

| PUT:     "/event/update/:id" , 
  Body:
    {
    "name" : "UpdatedTesttt",
    "description": "An updated test event",
    "location": "Local location",
    "dateTime":"2024-08-04T11:38:00.000Z"
    }

| DELETE API:  "/event/delete/:id"
| POST API:    "/events/:id/register"
->DELETE  "/events/:id/register/cancel"
