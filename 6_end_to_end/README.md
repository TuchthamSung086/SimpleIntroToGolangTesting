You want to really know if your system is working well along with other dependencies. AND, to make sure all your flow from
- the very beginning (exposed endpoint)
- to the complicated parts (request->handler->service->database/external)
- to the very end (response)

Quite similar to integration test

You make HTTP requests to the server itself and see if it's working.

Or... just use the following
- POSTMAN: API
- Frontend: Manual test
- Cypress: E2E Automation