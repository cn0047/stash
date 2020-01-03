function Rectangle (height, width) {
    this.height = height;
    this.width = width;
}
Rectangle.prototype.getSquare = function () {
    return this.height * this.width;
};
function Square (height) {
    Rectangle.apply(this, [height, height]);
}
Square.prototype = new Rectangle();

var r = new Rectangle(2, 3);
r.getSquare();

var s = new Square(2);
s.getSquare();
