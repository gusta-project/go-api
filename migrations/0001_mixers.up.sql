CREATE TABLE mixers (
  id BIGINT NOT NULL,
  uuid UUID NOT NULL,
  name VARCHAR(255) NOT NULL CHECK (name != ''),
  email VARCHAR(255) NOT NULL,
  salt VARCHAR(255) NOT NULL,
  passhash VARCHAR(255) NOT NULL,
  CONSTRAINT mixer_pk PRIMARY KEY (id),
  CONSTRAINT mixer_uuid_uq UNIQUE (uuid),
  CONSTRAINT mixer_name_uq UNIQUE (name)
);