/*
npm install pdfkit
*/
var doc = new PDFDocument();
doc.font('fonts/GoodDog-webfont.ttf')
.fontSize(25)
.text('Some text with an embedded font!', 100, 100);

doc.addPage()
.fontSize(25)
.text('Here is some vector graphics...', 100, 100);

doc.save()
.moveTo(100, 150)
.lineTo(100, 250)
.lineTo(200, 250)
.fill("#FF3300");

doc.scale(0.6)
.translate(470, −380)
.path('M 250,75 L 323,301 131,161 369,161 177,301 z')
.fill('red', 'even-odd')
.restore();

doc.addPage()
.fillColor("blue")
.text('Here is a link!', 100, 100)
.underline(100, 100, 160, 27, {color: "#0000FF"})
.link(100, 100, 160, 27, 'http://google.com/');
doc.write('output.pdf');



var spawn = require('child_process').spawn;
// получение фотографии
var photo = process.argv[2];
// массив преобразований
var opts = [
    photo,
    '-resize',
    '150',
    photo + ".png"
];
// преобразование
var im = spawn('convert', opts);
im.stderr.on('data', function (data) {
    console.log('stderr: ' + data);
});
im.on('exit', function (code) {
    if (code === 0)
        console.log('photo has been converted and is accessible at '+ photo + '.png');
});



var spawn = require('child_process').spawn;
// получение фотографии
var photo = process.argv[2];
// массив преобразований
var opts = [
    photo,
    "-bordercolor", "snow",
    "-border", "6",
    "-background","grey60",
    "-background", "none",
    "-rotate", "6",
    "-background", "black",
    "(", "+clone", "-shadow", "60x4+4+4", ")",
    "+swap",
    "-background", "none",
    "-flatten",
    photo + ".png"
];
var im = spawn('convert', opts);