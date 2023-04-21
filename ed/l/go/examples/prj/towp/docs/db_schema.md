wp DB schema
-

Current DB schema for wp looks like this:

````mermaid
erDiagram

wave_plan ||..|{ wave : ""
wave ||..|{ schedule : ""
schedule ||..|{ trigger : ""

wave_plan {
  STRING wave_plan_id PK
  STRING retailer_id
  STRING mfc_id
  STRING timezone
  TIMESTAMP created_at
  STRING created_by
}

wave {
  STRING wave_plan_id PK
  STRING wave_id PK
  STRING cutoff_time
  STRING from_time
  STRING to_time
}

schedule {
  STRING wave_plan_id PK
  STRING wave_id PK
  STRING schedule_id PK
  STRING schedule_type
  STRING schedule_time
}

trigger {
  STRING wave_plan_id PK
  STRING wave_id PK
  STRING schedule_id PK
  TIMESTAMP trigger_at PK
  TIMESTAMP cutoff_datetime
  TIMESTAMP created_at
  TIMESTAMP fired_at
}
````
