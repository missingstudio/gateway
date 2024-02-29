CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE prompts (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name text UNIQUE NOT NULL,
  description text,
  template text,
  metadata jsonb,
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW()
);