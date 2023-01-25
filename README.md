# Gorum - A mini forum MVP written in Go

Alas, I did not have enough time to complete the frontend and deployment. There's been a lot of backlog at work.

Nevertheless, I'm not going to make any excuses! The backend works fine and has a Swagger UI to prove it.

### Deployment

```bash
docker compose up -d --build
```

You can access the backend at localhost:8080 and the Swagger UI at localhost:8080/docs/index.html.

You can also access the two miserable frontend routes that are available at

- localhost
- localhost/login

### Libraries/Tools used

- Gin with Gin Swagger (I really like this tool, it helps with debugging)
- SQLite3
- React with React Router
- Tailwind (although I didn't get to use it much)
- Docker and Caddy (file server)
