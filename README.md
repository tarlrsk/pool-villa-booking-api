# De'Day Pool Villa — Backend

Go + Gin + GORM + PostgreSQL API serving both the customer-facing site and the admin dashboard for De'Day Pool Villa. Replaces the original Next.js app's Google Sheets "database" entirely.

## Stack

- Go 1.25, [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io) (`gorm.io/driver/postgres`)
- PostgreSQL
- JWT admin auth (`golang-jwt/jwt/v5` + `bcrypt`)
- Customer auth: LINE LIFF ID token, re-verified against LINE's endpoint on every request (no local session)

## Setup

1. Copy `.env.example` to `.env` and fill in real values (LINE credentials, a random `ADMIN_JWT_SECRET`, an `ADMIN_SEED_PASSWORD` for the first admin login).
2. Start Postgres:
   ```
   docker compose up -d postgres
   ```
   If Docker isn't available in your environment, run any Postgres 14+ instance and point `DATABASE_URL` at it — the app only needs a reachable Postgres, no Docker-specific behavior.
3. Migrate the schema:
   ```
   go run ./cmd/migrate
   ```
4. Seed default data (day-rate grid + one admin account from `ADMIN_SEED_USERNAME`/`ADMIN_SEED_PASSWORD`):
   ```
   go run ./cmd/seed
   ```
   Change the admin password after first login — there's no forced-change flow yet.
5. Run the API:
   ```
   go run ./cmd/api
   ```
   `cmd/api` only connects and serves — it does not touch schema or seed data, so re-running it never re-triggers migration/seeding. Both `cmd/migrate` and `cmd/seed` are safe to re-run any time (migration is additive-only; seeding only inserts when the target table is empty).

## Notes

- Dates (`checkin`, `checkout`, blocked dates, custom-period ranges) are stored as `YYYY-MM-DD` strings, not native date columns — the pricing/availability logic relies on lexical string comparison of ISO dates, matching the original TypeScript implementation.
- Migrations use GORM `AutoMigrate` only (no migration tool), run via `cmd/migrate`. It never drops or renames columns; destructive schema changes need a manual one-off SQL script. Revisit `golang-migrate` if schema churn becomes frequent.
- `LINE_AUTH_BYPASS=true` skips real LINE ID token verification (dev only) — useful for testing `POST /api/bookings` / `GET /api/user/bookings` with `curl` without a live LIFF session.
- The admin API surface (`/api/admin/blocked-dates`, `/api/admin/day-rates`, `/api/admin/custom-periods` GET/PUT/DELETE) goes beyond the original app's routes, since the original had no admin UI at all — these exist to support the new admin frontend.

## API

Public: `GET /healthz`, `GET /api/pricing-data`, `GET /api/availability`, `POST /api/webhook/line`

Customer (LINE bearer token): `POST /api/bookings`, `GET /api/user/bookings`

Admin (`POST /api/admin/login` issues a JWT, then `Authorization: Bearer <jwt>` on the rest):
`GET/PATCH /api/admin/bookings...`, `GET/POST/DELETE /api/admin/blocked-dates...`, `GET/PUT /api/admin/day-rates`, `GET/POST/PUT/DELETE /api/admin/custom-periods...`
