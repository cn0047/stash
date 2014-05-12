function Cat() {
    this.legs = 4;
    this.say = function () {
        return “meaowww”;
    }
}
function Bird() {
    this.wings = 2;
    this.fly = true;
}
function CatWings() {
    Cat.apply(this);
    Bird.apply(this);
}
var jane = new CatWings();
console.dir(jane);