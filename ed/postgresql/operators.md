Operators
-

`@>` - contains.

````sql
# true
SELECT '[1, 2, 3]'::jsonb @> '[3, 1]'::jsonb;
SELECT '[1, 2, 3]'::jsonb @> '[1, 2, 2]'::jsonb;
SELECT '{"p": "x", "version": 9.4}'::jsonb @> '{"version": 9.4}'::jsonb;
SELECT '[1, 2, [1, 3]]'::jsonb @> '[[1, 3]]'::jsonb;
SELECT '{"foo": {"bar": "baz"}}'::jsonb @> '{"foo": {}}'::jsonb;
SELECT '["foo", "bar"]'::jsonb @> '"bar"'::jsonb;
SELECT '["foo", "bar", "baz"]'::jsonb ? 'bar';
SELECT '"foo"'::jsonb ? 'foo';

# false
SELECT '[1, 2, [1, 3]]'::jsonb @> '[1, 3]'::jsonb;
SELECT '{"foo": {"bar": "baz"}}'::jsonb @> '{"bar": "baz"}'::jsonb;
````

`||` - strings concatenation.

````sql
SELECT 'foo ' || 'bar';
 ?column?
----------
 foo bar
````

`%` - modulo.
`|/` - square root.
`||/` - cube root.
`!` - factorial.
`@` - absolute value.
