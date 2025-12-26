UUID (Universally Unique Identifier)
-

UUID - 128-bit identifier used to uniquely identify information.

Versions:
v1: Time-based - MAC address + timestamp.
v2: DCE security - v1 + POSIX UID/GID (rarely used).
v3: Name-based - using MD5 hashing (same input = same UUID).
v4: Random - most commonly used.
v5: Name-based - using SHA-1 hashing.
v6: Reordered time-based - v1 but with sortable timestamp (for DB index).
v7: Unix timestamp + random - time-ordered + random (best modern choice, for DBs as well).

ULID (Universally unique Lexicographically sortable Identifier) was created
to solve the "randomness problem" of standard UUIDs (like v4).
Standard UUID v4 is completely random, when use it as primary keys in DB,
it causes index fragmentation, because IDs aren't in any order,
DB has to move data around constantly to insert new rows into its B-Tree index,
which slows down performance.
ULID solves this by putting timestamp at the beginning of the ID.
