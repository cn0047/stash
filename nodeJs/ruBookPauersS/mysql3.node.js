var Sequelize = require('sequelize');
var sequelize = new Sequelize(
    'databasenm',
    'username',
    'password',
    {logging: false}
);
// определение модели
var Nodetest2 = sequelize.define(
    'nodetest2',
    {
        id : {type: Sequelize.INTEGER, primaryKey: true},
        title : {type: Sequelize.STRING, allowNull: false, unique: true},
        text : Sequelize.TEXT,
    }
);
// синхронизация
Nodetest2.sync().error(function(err) {
    console.log(err);
});
var test = Nodetest2.build({
    title: 'New object',
    text: 'Newest object in the data store'
});
// сохранение записи
test.save().success(function() {
    // первое обновление
    Nodetest2.find({where : {title: 'New object'}}).success(function(test) {
        test.title = 'New object title';
        test.save().error(function(err) {
            console.log(err);
        });
        test.save().success(function() {
            // второе обновлен2ие
            Nodetest2.find({where : {title: 'New object title'}})
                .success(function(test) {
                    test.updateAttributes({title: 'An even better title'})
                        .success(function() {});
                    test.save().success(function() {
                        // поиск всего
                        Nodetest2.findAll().success(function(tests) {
                            console.log(tests);
                            // поиск и удаление нового объекта
                            Nodetest2.find({ where: {title: 'An even better title'}})
                                .success(function(test) {
                                    test.destroy().on('success', function(info) {
                                        console.log(info);
                                    });
                                });
                        });
                    });
                })
        });
    });
});