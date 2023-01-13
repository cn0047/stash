Migration
-

Migration must be idempotent (`IF NOT EXISTS`, etc).
When project old and has many migrations it takes long to start new (local env) DB.

Migration before code - when add new field.
Migration after code - when delete new field.
Rollback goes in opposite order.
