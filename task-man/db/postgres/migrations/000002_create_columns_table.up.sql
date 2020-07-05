BEGIN;

CREATE TABLE IF NOT EXISTS columns (
  id UUID DEFAULT uuid_generate_v4(),
  project_id UUID NOT NULL REFERENCES projects(id),
  name varchar(255) NOT NULL,
  priority INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NULL ,
  deleted_at TIMESTAMP DEFAULT NULL ,
  PRIMARY KEY (id),
  UNIQUE (id, name)
);

COMMIT;