CREATE TABLE flavors (
  id BIGINT NOT NULL,
  uuid UUID NOT NULL,
  slug VARCHAR(255) NOT NULL,
  vendor_id BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  CONSTRAINT flavor_pk PRIMARY KEY (id),
  CONSTRAINT flavor_uuid_uq UNIQUE (uuid),
  CONSTRAINT flavor_slug_uq UNIQUE (slug),
  CONSTRAINT flavor_name_vendor_uq UNIQUE (vendor_id, name),
  CONSTRAINT flavor_vendor_fk FOREIGN KEY (vendor_id) REFERENCES vendors (id)
);