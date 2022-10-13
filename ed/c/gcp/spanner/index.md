Index
-

````sql
CREATE UNIQUE INDEX tbl_unique_idx ON tbl (
  client_id, order_id
);

DROP INDEX tbl_unique_idx;

-- interleave:
CREATE INDEX SongsBySingerSongName ON Songs(SingerId, SongName), INTERLEAVE IN Singers;

-- store a copy of column in the index.
-- queries that use the index and select columns stored in the STORING clause
-- do not require an extra join to the base table.
CREATE INDEX AlbumsByAlbumTitle2 ON Albums(AlbumTitle) STORING (MarketingBudget);

````
