CREATE TABLE flavor (
  uuid UUID NOT NULL,
  vendor_uuid UUID NOT NULL,
  name VARCHAR(255) NOT NULL,
  CONSTRAINT flavor_pk PRIMARY KEY (uuid),
  CONSTRAINT flavor_name_vendor_uq UNIQUE (vendor_uuid, name),
  CONSTRAINT favlor_vendor_fk FOREIGN KEY (vendor_uuid) REFERENCES vendors (uuid)
);