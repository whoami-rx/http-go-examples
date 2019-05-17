# Middleware

A middleware is just a function that is executed before each request.

Commonly used to :

- Check auth
- Check cache
- Compress
- And more

Our middleware should implement `http.Handler` interface.