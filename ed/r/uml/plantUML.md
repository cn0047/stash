PlantUML
-

[docs](https://plantuml.com/)

````plantuml
@startuml
allowmixing

hide ComponentID
````

#### Component diagram.

Components: folder, package, node, rectangle, frame, cloud, database.

```plantuml
@startuml

title Example 1.

left to right direction

package Domain {
    node Admin1 {
    }

    rectangle Admin2 #LightGreen {
    }

    rectangle Admin3 #line.dotted {
    }

    rectangle Else #line.dotted {
    }

    frame Scope {
        rectangle Controller1 #line.dotted {
        }

        rectangle Controller2 #line.dotted {
        }

        rectangle Controller3 #line.dotted {
        }

        Controller2 -> Controller1
        Controller3 -> Controller1
    }

    Admin1 --> Controller1
    Admin1 --> Controller2
    Admin2 .-> Controller1
    Admin3 .-> Controller1
    Else .-> Controller1
    Else .-> Controller2
}

rectangle UI {
}

UI --> Controller3

@enduml
```
