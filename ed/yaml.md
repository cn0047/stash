YAML
-

Anchor:

````
DEFAULT: &DEFAULT
  description: "single ping #1"
  url: /cronTask/ping
  target: default
  schedule: every 1 minutes

cron:

- <<: *DEFAULT
````
