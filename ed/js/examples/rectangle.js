var Rectangle = function (w, h) {
    this.with = w;
    this.height = h;
};
Rectangle.prototype.getSquare = function() {
    return this.with * this.height;
};
var Rectangle1 = new Rectangle(3, 2);
console.log(Rectangle1.getSquare());
