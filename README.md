# Why?
POC! Exploring How Chat Works

# Does it Scale?
No.

Don't know how

# How?
Every conversation counted as a `session` later on it should be able to handle personal chat and multi-chat, I think

# What's Next?
- How do we store the sessions?
- How do we store sessions messages
- Should session be persistent?
- Should session messages be encrypted?
- How messages are encrypted?
- Possibility of breach?
- Should we stick to mysql?
- How many live connections can it handle?

# Set up
1. create `.env` from `.env.sample`
2. download docker
3. install dependecies `docker compose up -d`
4. run rest server `make run-rest`
5. run websocket server `make run-websocket`





