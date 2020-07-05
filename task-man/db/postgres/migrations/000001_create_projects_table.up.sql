BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS projects (
  id UUID DEFAULT uuid_generate_v4(),
  name varchar(500) NOT NULL,
  description varchar(1000) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NULL ,
  deleted_at TIMESTAMP DEFAULT NULL ,
  PRIMARY KEY (id)
);

COMMIT;