var Node = function () {
    this.elements = [];
    this.add = function (el) {
        this.elements.push(el);
    }
    this.show = function () {
        for (i in this.elements) {
            this.elements[i].show();
        }
    }
}
var Leaf = function (n) {
    this.name = n;
    this.show = function () {
        console.log(this.name);
    }
}

n = new Node();
n.add(new Leaf('L1'));
n.add(new Leaf('L2'));
sn = new Node();
sn.add(new Leaf('SL1'));
sn.add(new Leaf('SL2'));
n.add(sn);
n.show();

/*
L1
L2
SL1
SL2
*/
