CREATE TABLE vendors (
  id BIGINT NOT NULL,
  uuid UUID NOT NULL,
  slug VARCHAR(255) NOT NULL,
  code VARCHAR(10) NOT NULL CHECK (code != ''),
  name VARCHAR(255) NOT NULL CHECK (name != ''),
  url VARCHAR(255),
  CONSTRAINT vendor_pk PRIMARY KEY (id),
  CONSTRAINT vendor_uuid_uq UNIQUE (uuid),
  CONSTRAINT vendor_slug_uq UNIQUE (slug),
  CONSTRAINT vendor_code_uq UNIQUE (code),
  CONSTRAINT vendor_name_uq UNIQUE (name)
);