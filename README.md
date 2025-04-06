# realhousewives-api

This is a public API for retrieving information about [The Real Housewives](https://en.wikipedia.org/wiki/The_Real_Housewives) franchise. The full documentation can be viewed at https://realhousewives-api.up.railway.app/

## Lessons Learned

- During the initial design phase, I struggled slightly with deciding when to use path parameters vs. query parameters for requesting resources.
  - A common convention I followed was using path parameters for identifying a specific resource and query parameters for filtering those resources.
- I set up a route group so that I don't have to prefix each of my endpoints with "/v1".
- Implemented a singleton pattern using Go's `sync` module so that multiple instances of the database's connection pool won't be created.
- Used rate limiting middleware that keeps track of client IPs and utilizes a token bucket algorithm for allocating requests per client.
- Created error handling middleware for consistent formatting of errors to client.
- The most impactful lesson in this project was learning the OpenAPI 3.1 specification for documenting API.
  - I had no knowledge of this specification beforehand, so I was messily documenting my API in a Google Doc before finding out about it towards the tailend of this project when v1 development was almost finished.
  - The `openapi.yml` file in this repository is written by me without any generators! 😀
- I generated an API reference client based on my OpenAPI specs using Scalar.
