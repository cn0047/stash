YAML
-

detection: ~

````yaml
include_newlines: |
            exactly as
            you see
fold_newlines: >
            this is
            single line
# null value
nullparam: null
nullparam: ~
nullparam:
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
