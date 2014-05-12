function Universe() {
    // имеется ли экземпляр, созданный ранее?
    if (typeof Universe.instance === “object”) {
        return Universe.instance;
    }
    // создать новый экземпляр
    this.start_time = 0;
    this.bang = “Big”;
    // сохранить его
    Universe.instance = this;
    // неявный возврат экземпляра:
    // return this;
}
// проверка
var uni = new Universe();
var uni2 = new Universe();
uni === uni2; // true
