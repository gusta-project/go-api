CREATE TABLE vendors (
  uuid UUID NOT NULL,
  code VARCHAR(10) NOT NULL CHECK (code != ''),
  name VARCHAR(255) NOT NULL CHECK (name != ''),
  url VARCHAR(255),
  CONSTRAINT vendor_pk PRIMARY KEY (uuid),
  CONSTRAINT vendor_code_uq UNIQUE (code),
  CONSTRAINT vendor_name_uq UNIQUE (name)
);