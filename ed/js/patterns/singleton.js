function Universe() {
    if (typeof Universe.instance === "object") {
        return Universe.instance;
    }
    this.start_time = 0;
    this.bang = "Big";
    Universe.instance = this;
    // return this;
}
var uni = new Universe();
var uni2 = new Universe();
uni === uni2; // true
