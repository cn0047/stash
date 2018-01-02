create table test (
  id serial NOT NULL PRIMARY KEY,
  n int,
  d double precision,
  s character varying(20)
);

create table products (
  product_no integer,
  name text,
  price numeric,
  CHECK (price > 0), -- constraint
  discounted_price numeric,
  CHECK (discounted_price > 0),
  CHECK (price > discounted_price) 
);

create table example (
  a integer,
  b integer,
  c integer,
  PRIMARY KEY (a, b),
  UNIQUE (a, c) 
);
