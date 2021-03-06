# Shaky Table (WIP)

Ongoing attempt at trying to recreate elements of Airtable's spreadsheet-database hybrid concept.

Main plan is to use MongoDB as the actual storage for the tables, while SQLite is used to store more traditional credentials and other data. For the frontend side, I've always wanted to mess around with Svelte, so I might just try that. It would be a good time to look into "windowing" a la `react-virtualized`.

I'll be using to `docker-compose` to hook everything up for ease of development.

Development:

from the project root: `docker-compose up`

Roadmap:
- [x] go server creation
- [x] sveltekit client creation
- [x] user account creation and auth with SQLite
- [x] `docker-compose` orchestration (for dev)
- [ ] pass a web session for login
- [ ] build spreadsheet API with MongoDB
- [ ] client-side rendering of spreadsheet data
- [ ] build webpage UX out
- [ ] parameterize build and containerization options
- [ ] additional MongoDB optimizations (Change Streams)
- [ ] (maybe) deploy for demo use
