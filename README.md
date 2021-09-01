# Shaky-Table (WIP)

Ongoing attempt at trying to recreate elements of Airtable's spreadsheet-database hybrid concept.

Main plan is to use MongoDB as the actual storage for the tables, while SQLite is used to store more traditional credentials and other data. For the frontend side, I've always wanted to mess around with Svelte, so I might just try that. It would be a good time to look into "windowing" a la `react-virtualized`.

I'll probably just resort to `docker-compose` to hook everything up for ease of development.
