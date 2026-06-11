# mapene-mcp — a Black Swan product

> Model Context Protocol (MCP) server that exposes Apache Fineract banking
> operations as agent tools, for use with the **Mapene** web application by
> **Black Swan** (BSAi Global Limited).

mapene-mcp lets an AI assistant safely drive banking operations (clients, loans,
savings, accounting, and more) through the same Fineract backend Mapene uses. It
exposes typed tools via the [Model Context Protocol](https://modelcontextprotocol.io)
and runs over MCP stdio, or as an authenticated HTTP/SSE service. Built in Go for
low-latency, high-concurrency execution.

## Run (Docker)

```bash
docker build -t mapene-mcp:local .
docker run -d --name mapene-mcp -p 8080:8080 \
  -e PORT=8080 \
  -e MCP_AUTH_TOKEN="$(openssl rand -hex 24)" \
  -e MIFOSX_BASE_URL=https://host.docker.internal:8443/fineract-provider/api/v1 \
  -e MIFOSX_TENANT_ID=default \
  -e MIFOSX_USERNAME=mifos \
  -e MIFOSX_PASSWORD=password \
  mapene-mcp:local
```

- **stdio mode**: omit `PORT` and run the binary directly (for desktop MCP clients).
- **HTTP/SSE mode**: set `PORT`. Endpoints: `/sse`, `/message`, `/api/<tool>`, `/health`, `/metrics`.

## Security (read before exposing)

- The HTTP/SSE surface is an **authenticated gateway** to Fineract using the
  configured service credentials — never an open proxy.
- **Fail-closed**: with `MCP_AUTH_TOKEN` unset, only `/health` responds; all other
  routes return 401. With it set, send `Authorization: Bearer <token>`.
- CORS is same-origin unless `MCP_ALLOWED_ORIGIN` opts an origin in.
- Run on a private network only; rotate the token; terminate TLS at a trusted proxy.

## Configuration

| Env | Purpose |
|---|---|
| `MIFOSX_BASE_URL` | Fineract API base |
| `MIFOSX_TENANT_ID` | Fineract tenant (`Fineract-Platform-TenantId`) |
| `MIFOSX_USERNAME` / `MIFOSX_PASSWORD` | Fineract service credentials |
| `PORT` | If set, run HTTP/SSE on this port (else stdio) |
| `MCP_AUTH_TOKEN` | Bearer token required on all routes except `/health` |
| `MCP_ALLOWED_ORIGIN` | CORS allowed origin (optional) |

## Licensing

mapene-mcp is a Black Swan product (© 2026 BSAi Global Limited; all Black Swan
rights reserved). It is built on the Mifos Initiative `mcp-mifosx` base, which is
licensed under the **Mozilla Public License v2.0**; that license is retained for
the base code. Black Swan modifications and the Black Swan / Mapene / mapene-mcp
names and marks are reserved. See [`NOTICE`](./NOTICE).
