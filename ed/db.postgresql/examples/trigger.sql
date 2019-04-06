-- Create table.
CREATE TABLE trgr (
  id INTEGER PRIMARY KEY,
  n INTEGER
);

-- Create trigger.
CREATE OR REPLACE FUNCTION trgr_autodelete() RETURNS trigger AS $$
BEGIN
  IF NEW.n = 0 THEN
    DELETE from trgr WHERE id = NEW.id;
  END IF;
  RETURN NULL;
END
$$ LANGUAGE 'plpgsql';
CREATE TRIGGER trgr_autodelete AFTER INSERT OR UPDATE ON trgr FOR EACH ROW EXECUTE PROCEDURE trgr_autodelete();

-- Test.
INSERT INTO trgr VALUES (1, 1);
SELECT * FROM trgr;
UPDATE trgr SET n = trgr.n + 1;
UPDATE trgr SET n = trgr.n - 1;
