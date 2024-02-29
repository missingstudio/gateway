CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE apikeys (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name text NOT NULL,
  value text UNIQUE NOT NULL,
  last_used_at timestamp DEFAULT NULL,
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW()
);