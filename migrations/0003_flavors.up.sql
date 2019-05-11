CREATE TABLE flavors (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  slug VARCHAR(255) NOT NULL UNIQUE CHECK (slug != ''),
  vendor_id BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL CHECK (name != ''),
  gravity numeric(4, 4),
  UNIQUE (vendor_id, name),
  FOREIGN KEY (vendor_id) REFERENCES vendors (id)
);