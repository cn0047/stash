var readline = require('readline');
// создание нового интерфейса
var interface = readline.createInterface(process.stdin, process.stdout, null);
// задание вопроса
interface.question(">>What is the meaning of life? ", function(answer) {
    console.log("About the meaning of life, you said: " + answer);
    interface.setPrompt(">>");
    interface.prompt();
});
// функция для закрытия интерфейса
function closeInterface() {
    console.log('Leaving interface...');
    process.exit();
}
// прослушивание команды .leave
interface.on('line', function(cmd) {
    if (cmd.trim() == '.leave') {
        closeInterface();
        return;
    } else {
        console.log("repeating command: " + cmd);
    }
    interface.setPrompt(">>");
    interface.prompt();
});
interface.on('close', function() {
    closeInterface();
});


// When type: .leave script will exit.