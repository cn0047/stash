var TV = function () {
    this.watch = function () {
        console.log('you are watching...');
    }
}
var parentControlProxy = function () {
    this.watchTv = function () {
        var d = new Date();
        if (d.getHours() >= 22) {
            console.log('you should go sleep...');
            return;;
        }
        var tv = new TV();
        tv.watch();
    }
}
var p = new parentControlProxy();
p.watchTv();
