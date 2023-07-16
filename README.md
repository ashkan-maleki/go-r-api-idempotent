# Idempotent APIs: Avoid Duplicate Requests with Golang and Redis


## Process:
### First API Call
1) Make a POST request to the "shipping/order" API endpoint.
2) The specified "order_id" is checked in the Redis. If it exists, the object is sent to the user.
3) if it doesn't exist, the service stores the new object in Postgres and Redis after 3 seconds (this 3 seconds delay is
   deliberately used to demonstrate the process).

### Second API Call
1) Make a POST request to the "shipping/order" API endpoint with the same "order_id" parameter.
2) The service responds and returns the object instantly.

# Requirements
1) Golang, 2) Redis, 3) Postgres, 4) Docker, 5) Docker Compose, 6) Make, 7) Gofiber
