PlantUML
-

[docs](https://plantuml.com/)
[guide](https://crashedmind.github.io/PlantUMLHitchhikersGuide/)

````plantuml
@startuml
allowmixing

!includesub diagram.subpart.1.puml!SUBPART1

hide ComponentID

'comment

#line.dotted
#lightgreen
#lightyellow
#lightpink
````

#### Component diagram.

Components: folder, package, node, rectangle, frame, cloud, database.

```plantuml
@startuml

title Componets diagram, example #1.

left to right direction
top to bottom direction
allowmixing

package Domain {
    node Admin1 {
    }

    rectangle Admin2 #lightgreen {
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

```plantuml
@startuml

title State diagram, example #1.

state Up {
}

Up --> A

state A {
    state B #lightpink {
    }

    state C #lightyellow {
    }

    [*] -[#pink]-> B
    [*] -[#grey]-> C
    C -d-> D
    D -r-> E
    D -[hidden]r-> E
}

@enduml
```
