# Database Migrations

This folder contains versioned SQL migrations for the **LeitnerBox** backend.
Migrations keep the database schema consistent across local development,
staging, and production environments.

We use [golang-migrate](https://github.com/golang-migrate/migrate) as the
migration tool.

---

## 📂 Structure

Each migration has an `up` (apply) and `down` (rollback) file:

- **`001_create_users.up.sql`** → creates the `users` table  
- **`001_create_users.down.sql`** → drops the `users` table  

Migrations are applied in order of their version prefix.

---

## 🚀 Usage

### Run migrations (up)
```bash
migrate -path migrations -database "$DATABASE_URL" up
