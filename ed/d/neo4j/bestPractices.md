Best practices
-

Names:
Node label   - `Person` upper first camel case;
Relationship - `ACTED_IN` upper case with underscores;
Property     - `VehicleOwner` camel case;

Use `node._lock = true` suggested by community in neo4j docs.

Filter by properties - most expensive.
Filter by node label - good.
Filter by relationship type - very good.
