---

```markdown
# 🛠 Troubleshooting Log: Redis Connectivity Issues

## Incident: "Too Many Colons" Connection Failure
**Date:** March 8, 2026  
**Service:** Go Health-API  
**Environment:** Render (Production) / Upstash (Database)

---

### 1. The Symptom 🚩
The application failed to start on Render. The deployment logs showed the following error:
> `2026/03/08 18:42:19 CRITICAL: Could not connect to Redis: dial tcp: address [URL]: too many colons in address`
> `==> Exited with status 1`

### 2. Root Cause Analysis (RCA) 🔍
There were two primary issues identified:

1.  **Malformed Environment Variable:** The `REDIS_ADDR` variable in the Render Dashboard contained the full terminal command (`redis-cli --tls -u ...`) instead of just the connection URL.
2.  **Library Parse Error:** The `go-redis` library’s default `redis.Options{Addr: ...}` field expects a simple `host:port` string. When provided with a full URI (e.g., `rediss://user:pass@host:port`), the library was unable to parse the multiple colons used for the protocol and credentials.



### 3. The Resolution 🛠️
The fix was implemented in two stages:

#### A. Configuration Clean-up
Modified the environment variable in Render to include **only** the connection string, ensuring the use of the `rediss://` protocol for TLS/SSL compatibility with Upstash.

#### B. Code Refactor (Implementation of `ParseURL`)
Updated `main.go` to use the `redis.ParseURL()` function. This function is designed to handle complex URIs, automatically extracting the hostname, password, and enabling TLS.

```go
// Correct way to handle Upstash/Cloud Redis URIs
opt, err := redis.ParseURL(os.Getenv("REDIS_ADDR"))
if err != nil {
    log.Fatalf("Invalid Redis URL: %v", err)
}
rdb = redis.NewClient(opt)

```

### 4. Implementation of "Fail-Fast" & Observability 🛡️

To prevent "silent failures" in the future, the following "Bridge" logic was added:

* **Startup Ping:** The app now executes `rdb.Ping(ctx)` before starting the HTTP server. If the database is unreachable, the app exits immediately.
* **Connection Logging:** Added `fmt.Printf` to log the target Redis address (redacted) to ensure visibility during the deployment phase.

---

### 5. Lessons Learned 💡

* **Log the 'Where':** Always log which external resource the app is attempting to connect to during startup.
* **Local Validation:** Use `go build` in the local VM to catch syntax errors before pushing to the CI/CD pipeline.
* **Use Protocol-Aware Parsers:** When dealing with cloud databases (Upstash, AWS ElastiCache), always use `ParseURL` instead of manual string concatenation for addresses.

```

---

### 🚀 Next Step
