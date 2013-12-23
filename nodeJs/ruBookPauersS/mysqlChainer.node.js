var Sequelize = require('sequelize');
var sequelize = new Sequelize(
    'databasenm',
    'username',
    'password',
    { logging: false}
);
// определение модели
var Nodetest2 = sequelize.define(
    'nodetest2',
    {
        id : {type: Sequelize.INTEGER, primaryKey: true},
        title : {type: Sequelize.STRING, allowNull: false, unique: true},
        text : Sequelize.TEXT
    }
);
// синхронизация
Nodetest2.sync().error(function(err) {
    console.log(err);
})
var chainer = new Sequelize.Utils.QueryChainer;
chainer.add(Nodetest2.create({title: 'A second object',text: 'second'}))
    .add(Nodetest2.create({title: 'A third object', text: 'third'}));
chainer.run()
    .error(function(errors) {
        console.log(errors);
    })
    .success(function() {
        Nodetest2.findAll().success(function(tests) {
            console.log(tests);
        });
    });