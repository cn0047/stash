ClickHouse
-
<br>24.5.3.5

[docs](https://clickhouse.com/docs)
[ttl](https://clickhouse.com/docs/en/guides/developer/ttl)
[playground](https://sql.clickhouse.com/)

ClickHouse - open-source column-oriented DBMS for real time analytical reporting.

````sql
use db;

````

Materialized view - shifts the cost of computation from query time to insert time.

Projections - store data in a format that optimizes query execution.

**Data types:**

Int (Int8, Int16, Int32, In64, Int128, Int256).
UInt (UInt8, UInt16, UInt32, UIn64, UInt128, UInt256).
Float32, Float64.
BFloat16 - 16-bit floating point data type with 8-bit exponent, sign, and 7-bit mantissa.
Decimal - signed fixed-point numbers that keep precision during add/subtract/multiply operations.
String, FixedString.
Date - 2 bytes number of days since 1970-01-01.
Date32 - 32-bit integer value representing the days since 1900-01-01.
Time - hour, minute, and second.
Time64 - represents time-of-day with fractional seconds.
DateTime -  date and time of day.
DateTime64 - date and time of day, with defined sub-second precision.
Enum.
UUID - sequence of 16 random bytes.
IPv4, IPv6.
Array(T) - array of `T-type` items.
Bool.
Tuple - tuple of elements, each having an individual type.
Map - key-value pairs.
Variant - represents union of other data types.
LowCardinality - changes internal representation of other data types to be dictionary-encoded.
Point, Ring, LineString, MultiLineString, Polygon, MultiPolygon, Geometry - geographical objects.
Nested - like a table inside a cell.
Dynamic - stores values of any type.
JSON.
QBit - data type reorganizes vector storage for faster approximate searches.
