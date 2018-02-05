package main

public class Person {
    private String name;
    public String getName() {
        return this.name;
    }
}

public class Saiyan {
    // Saiyan is said to have a person
    private Person person;

    // we forward the call to person
    public String getName() {
        return this.person.getName();
    }
}

func main() {
    // @TODO: Finish it)
}
