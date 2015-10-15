function Player(name) {
    this.points = 0;
    this.name = name;
}
Player.prototype.play = function () {
    this.points += 1;
    mediator.played();
};
var scoreboard = {
    // refreshable html
    element: document.getElementById('results'),
    // refresh screen
    update: function (score) {
        var i, msg = '';
        for (i in score) {
            if (score.hasOwnProperty(i)) {
                msg += '<p><strong>' + i + '<\/strong>: ';
                msg += score[i];
                msg += '<\/p>';
            }
        }
        this.element.innerHTML = msg;
    }
};
var mediator = {
    players: {},
    // init
    setup: function () {
        var players = this.players;
        players.home = new Player('Home');
        players.guest = new Player('Guest');
    },
    // refresh score
    played: function () {
        var players = this.players,
        score = {
            Home: players.home.points,
            Guest: players.guest.points
        };
        scoreboard.update(score);
    },
    // handler
    keypress: function (e) {
        e = e || window.event; // IE
        if (e.which === 49) { // key “1”
            mediator.players.home.play();
            return;
        }
        if (e.which === 48) { // key “0”
            mediator.players.guest.play();
            return;
        }
    }
};

// Start
mediator.setup();
window.onkeypress = mediator.keypress;
// After 30 second will be game over
setTimeout(function () {
    window.onkeypress = null;
    alert('Game over!');
}, 30000);
