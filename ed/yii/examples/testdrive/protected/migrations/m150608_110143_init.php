<?php

class m150608_110143_init extends CDbMigration
{
    public function safeUp()
    {
        $this->execute("
            create table if not exists cars (
                id int auto_increment,
                brand varchar(100) not null default '',
                model varchar(100) not null default '',
                maxSpeed int not null default '0',
                primary key (id)
            )
        ");
        $this->execute("
            insert into cars values
            (null, 'bmw', '5', 220),
            (null, 'bmw', 'x5', 250),
            (null, 'bmw', 'm6', 290),
            (null, 'audi', 'q7', 270),
            (null, 'audi', 'a4', 250),
            (null, 'bmx', 'x1', 260),
            (null, 'aston martin', 'db9', 320),
            (null, 'aston martin', 'dbs v12', 350),
            (null, 'peugeot', '301', 160),
            (null, 'peugeot', '401', 190),
            (null, 'peugeot', '407', 210),
            (null, 'toyota', 'corolla', 210),
            (null, 'toyota', 'camry', 250),
            (null, 'toyota', 'prado', 190),
            (null, 'toyota', 'land cruiser', 200),
            (null, 'porsche', 'cayen', 240),
            (null, 'porsche', '911', 270),
            (null, 'porsche', 'boxter', 280),
            (null, 'porsche', 'spider', 290),
            (null, 'porsche', 'carrera', 300),
            (null, 'citroen', 'c1', 110),
            (null, 'citroen', 'c3', 150),
            (null, 'citroen', 'c4', 190),
            (null, 'citroen', 'ds3', 210),
            (null, 'citroen', 'ds4', 250)
        ");
        return true;
    }

    public function down()
    {$this->execute("drop table cars");
        return true;
    }
}
