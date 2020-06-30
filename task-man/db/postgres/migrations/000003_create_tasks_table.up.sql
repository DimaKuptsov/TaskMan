BEGIN;

CREATE TABLE IF NOT EXISTS tasks (
  id UUID DEFAULT uuid_generate_v4(),
  column_id UUID NOT NULL REFERENCES columns(id),
  name varchar(500) NOT NULL,
  description varchar(5000) NOT NULL DEFAULT '',
  priority INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NULL ,
  deleted_at TIMESTAMP DEFAULT NULL ,
  PRIMARY KEY (id)
);

COMMIT;