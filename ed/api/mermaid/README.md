mermaid
-

[docs](https://mermaid-js.github.io/)
[editor](https://mermaid-js.github.io/docs/mermaid-live-editor-beta)

````
````mermaid
erDiagram
one ||..|{ two : ""

one {
  STRING id PK
  STRING msg
  STRING created_by
}

two {
  STRING id PK
  STRING msg
  STRING created_by
}
````

````mermaid
flowchart TD

client[Client]
LB[Load Balancing]
app[App]
db[DB]

client -- sends request --> LB
LB --> app
app --> db
````
