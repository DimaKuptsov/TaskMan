BEGIN;

CREATE TABLE IF NOT EXISTS comments (
  id UUID DEFAULT uuid_generate_v4(),
  task_id UUID NOT NULL REFERENCES tasks(id),
  text varchar(5000) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NULL ,
  deleted_at TIMESTAMP DEFAULT NULL ,
  PRIMARY KEY (id)
);

COMMIT;