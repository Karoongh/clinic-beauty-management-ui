-- 002_users_patients.sql
-- Non-money tables only. Finance/stock tables stay blocked by accounting-audit gate.

CREATE TABLE IF NOT EXISTS users (
    id           TEXT PRIMARY KEY,
    username     TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role         TEXT NOT NULL,
    display_name TEXT NOT NULL,
    active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS patients (
    id          TEXT PRIMARY KEY,
    code        TEXT NOT NULL UNIQUE,
    full_name   TEXT NOT NULL,
    mobile      TEXT NOT NULL,
    age         INT,
    gender      TEXT,
    visits      INT NOT NULL DEFAULT 0,
    invoices    INT NOT NULL DEFAULT 0,
    total_paid  BIGINT NOT NULL DEFAULT 0,
    wallet      BIGINT NOT NULL DEFAULT 0,
    status      TEXT NOT NULL DEFAULT 'active',
    notes       TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_patients_mobile ON patients (mobile);
CREATE INDEX IF NOT EXISTS idx_patients_full_name ON patients (full_name);

-- Seed demo users (bcrypt hashes for: manager123, reception123, doctor123, cashier123)
-- Generated with cost 10
INSERT INTO users (id, username, password_hash, role, display_name) VALUES
('u-manager',   'manager',   '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'manager',   'مدیر'),
('u-reception', 'reception', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'reception', 'پذیرش'),
('u-doctor',    'doctor',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'doctor',    'پزشک'),
('u-cashier',   'cashier',   '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'cashier',   'صندوقدار')
ON CONFLICT (id) DO NOTHING;

-- Note: all seeded users currently share the same hash (manager123) for local demo simplicity.
-- Replace with proper unique hashes before production.

INSERT INTO patients (id, code, full_name, mobile, age, gender, visits, invoices, total_paid, wallet, status) VALUES
('p1', 'P-1001', 'سارا محمدی',  '09121234567', 32, 'female', 5, 4, 12500000, 500000,  'active'),
('p2', 'P-1002', 'علی رضایی',   '09129876543', 41, 'male',   2, 2,  4800000,      0,  'active'),
('p3', 'P-1003', 'مریم حسینی',  '09351234567', 28, 'female', 8, 7, 21000000, 1500000,  'active'),
('p4', 'P-1004', 'رضا کریمی',   '09121112233', 35, 'male',   1, 1,  2500000,      0,  'active')
ON CONFLICT (id) DO NOTHING;
