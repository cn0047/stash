<?php

/*
CREATE DATABASE testdrive;
CREATE DATABASE testdrive_unit_test;

DROP DATABASE testdrive_unit_test;
DROP DATABASE testdrive;
*/

class m150608_110143_init extends CDbMigration
{
    public function safeUp()
    {
        $sql = <<<"SQL"
            CREATE TABLE brand (
                id INT AUTO_INCREMENT,
                name VARCHAR(100) NOT NULL DEFAULT '',
                country VARCHAR(50) NOT NULL DEFAULT '',
                PRIMARY KEY (id)
            );
            CREATE TABLE car (
                id INT AUTO_INCREMENT,
                brand_id INT NOT NULL DEFAULT '0',
                model VARCHAR(100) NOT NULL DEFAULT '',
                maxSpeed INT NOT NULL DEFAULT '0',
                PRIMARY KEY (id),
                INDEX brand_id (brand_id),
                FOREIGN KEY (brand_id) REFERENCES brand(id) ON DELETE CASCADE
            );
            INSERT INTO brand VALUES
                (null, 'aston martin', 'UK'),
                (null, 'audi', 'Germany'),
                (null, 'bmw', 'Germany'),
                (null, 'citroen', 'France'),
                (null, 'peugeot', 'France'),
                (null, 'porsche', 'Germany'),
                (null, 'toyota', 'Japan'),
                (null, 'ferrari', 'Italy')
            ;
            INSERT INTO car VALUES
                (null, 1, 'db9', 320),
                (null, 1, 'dbs v12', 350),
                (null, 2, 'a4', 250),
                (null, 2, 'q7', 270),
                (null, 3, '5', 220),
                (null, 3, 'm6', 290),
                (null, 3, 'x1', 260),
                (null, 3, 'x5', 250),
                (null, 4, 'c1', 110),
                (null, 4, 'c3', 150),
                (null, 4, 'c4', 190),
                (null, 4, 'ds3', 210),
                (null, 4, 'ds4', 250),
                (null, 5, '301', 160),
                (null, 5, '401', 190),
                (null, 5, '407', 210),
                (null, 6, '911', 270),
                (null, 6, 'boxter', 280),
                (null, 6, 'carrera', 300),
                (null, 6, 'cayen', 240),
                (null, 6, 'spider', 290),
                (null, 7, 'camry', 250),
                (null, 7, 'corolla', 210),
                (null, 7, 'land cruiser', 200),
                (null, 7, 'prado', 190)
            ;
SQL;
        $this->setDbConnection(Yii::app()->db);
        $this->execute($sql);
        $this->setDbConnection(Yii::app()->dbUnitTest);
        $this->execute($sql);
        return true;
    }

    public function down()
    {
        $sql = <<<"SQL"
            DELETE FROM brand;
            DROP TABLE car;
            DROP TABLE brand;
SQL;
        $this->setDbConnection(Yii::app()->dbUnitTest);
        $this->execute($sql);
        $this->setDbConnection(Yii::app()->db);
        $this->execute($sql);
        return true;
    }
}
