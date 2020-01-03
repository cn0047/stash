YAML
-

````yaml
haiku: >
  Consider me
  As one who loved poetry
  And persimmons.
````

Anchor:

````yaml
DEFAULT: &DEFAULT
  description: "single ping #1"
  url: /cronTask/ping
  target: default
  schedule: every 1 minutes

cron:

- <<: *DEFAULT
````

or:

````yaml
song:
  - &name Al
  - can
  - *name

# result:

song:
  - Al
  - can
  - Al
````
