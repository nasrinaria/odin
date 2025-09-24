## Gateway Service
The single public entry point for the application.
Validates JWTs, applies rate limits, logs requests, and forwards them to the right internal service (Auth, SRS, Media).
Makes the whole system look like one clean API to the client.
