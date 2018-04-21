GopherCon 2016
-

#### Francesc Campoy - Understanding nil

👍

#### Dave Cheney - Dont Just Check Errors Handle Them Gracefully

* err.Error() - for human, not for code.
* Error types - tight coupling bettwin packages (1 pckg must import error from 2nd pckg).
* When have this code in different layers - it's not clear where error happened:
````
err := do()
if err != nil {
    return err
}
````

#### Wisdom Omuya - Go Vendoring Deconstructed

https://monosnap.com/file/v3SM6L4GEFLJxVUyBnNn4dt3oXUGrh

https://www.youtube.com/watch?v=lI17OEJCPVw&index=21&list=PL2ntRZ1ySWBdliXelGAItjzTMxy2WQh0P
