This is a personal project that is evolving as I learn new things. I have included an executable application based on the current progress.  

- By default, it listens on port 8080.

This is a chat app that is still in development, but the basic structure is complete.

What it does:
1. When a client visits the website (currently hosted on localhost) at `/`, an HTML page is served that allows the client to write a message.
2. When the client sends the message, the HTML uses HTMX to make a POST request to the server.
3. For now, the server processes the request and upgrades the HTTP connection to a WebSocket.
4. It then stores the message in a data structure to share it with everyone currently connected.
5. I have ensured that the message is not sent back to the sender.

Future plans:
1. Add authentication.
2. Store user information in a database.
3. Possibly add a "Sign in with Google" option.
