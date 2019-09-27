<?php

namespace aExt;

/**
 * Takes all models from observable directory and checks them.
 *
 * Using reflection, in each finded model from phpDocumentation
 * extracts fields and keys that should be present in the DB.
 *
 * @see protected/modules/aExt/extensions/aExt/models/ AExtension module AR models.
 * @group general
 */
class DBTest extends \CTestCase
{
    public static $db;
    public static $existsTables;

    public static function setUpBeforeClass()
    {
        \Yii::import('application.modules.aExt.extensions.aExt.models.*');
        self::$db = \Yii::app()->eAExtension->getDbConnection();
        self::$existsTables = \Yii::app()->eAExtension->getDbConnection()->schema->getTables();
    }

    public function provider()
    {
        $provider = [];
        $path = \Yii::getPathOfAlias('application.modules.aExt.extensions.aExt.models.');
        $models = \CFileHelper::findFiles($path, [
            'exclude'   => ['EAExtensionActiveRecord.php'],
            'fileTypes' => ['php'],
            'level'     => 0,
        ]);
        foreach ($models as $modelFile) {
            $modelName = rtrim(basename($modelFile), '.php');
            $provider[] = [
                \Yii::app()->eAExtension->newModel($modelName)
            ];
        }
        return $provider;
    }

    /**
     * @dataProvider provider
     */
    public function testDBSchema($model)
    {
        $tableName = $model->tableName();
        // Tables
        $this->assertArrayHasKey(
            $tableName,
            self::$existsTables,
            "Model '$tableName' don't have existing table in DB."
        );
        unset(self::$existsTables[$tableName]);
        // Fields
        $existsColumns = $model->tableSchema->columnNames;
        $this->assertNotEmpty($existsColumns);
        $this->assertInternalType('array', $existsColumns);
        $reflection = new \ReflectionClass($model);
        $comment = $reflection->getDocComment();
        $this->assertNotEmpty($comment, 'Cannot read AR Model phpDoc comment.');
        preg_match_all(
            '/\*\s+@property\s+\w{3,}\s+\$(\w+)\s+(PRIMARY\sKEY|UNIQUE\sKEY|KEY)?.*/',
            $comment,
            $expectedColumns
        );
        $this->assertNotEmpty($expectedColumns[1], 'Cannot read AR Model properties from the phpDoc comment.');
        $msg = "Current table `{$tableName}` structure is not equals with expected.";
        $diff = array_diff($expectedColumns[1], $existsColumns);
        $this->assertEmpty(
            $diff,
            "$msg DataBase doesn't contains but AR Model contains next fields: ".implode(',', $diff)
        );
        $diff = array_diff($existsColumns, $expectedColumns[1]);
        $this->assertEmpty(
            $diff,
            "$msg AR Model doesn't contains but DataBase contains next fields: ".implode(',', $diff)
        );
        // Indexes
        $this->assertNotEmpty($expectedColumns[2], 'Cannot read AR Model keys from the phpDoc comment.');
        $expectedKeys = array_combine($expectedColumns[1], $expectedColumns[2]);
        $this->assertNotEmpty($expectedKeys);
        $this->assertInternalType('array', $expectedKeys);
        $existsKeys = self::$db->createCommand("SHOW INDEXES FROM `{$tableName}`")->queryAll();
        $this->assertNotEmpty($existsKeys);
        $this->assertInternalType('array', $existsKeys);
        foreach ($existsKeys as $key) {
            $this->assertArrayHasKey($key['Column_name'], $expectedKeys);
            $this->assertNotEmpty(
                $expectedKeys[$key['Column_name']],
                "$msg AR Model doesn't contains but DataBase contains next {$key['Key_name']} key - `{$key['Column_name']}`."
            );
            if ($key['Key_name'] == 'PRIMARY') {
                $this->assertEquals($expectedKeys[$key['Column_name']], 'PRIMARY KEY');
            } else {
                if ($key['Non_unique'] == 1) {
                    $this->assertEquals($expectedKeys[$key['Column_name']], 'KEY');
                } else {
                    $this->assertEquals($expectedKeys[$key['Column_name']], 'UNIQUE KEY');
                }
            }
            unset($expectedKeys[$key['Column_name']]);
        }
        foreach ($expectedKeys as $i => $key) {
            $this->assertEmpty(
                $key,
                "$msg DataBase doesn't contains but AR Model contains next `$i` key - $key."
            );
        }
    }

    public function testExistsTables()
    {
        $this->assertEmpty(
            self::$existsTables,
            'Next tables exists in DB, but they have not models: '.implode(',', array_keys(self::$existsTables))
        );
    }

    public function testSites()
    {
        $count = (int)self::$db->createCommand("SELECT COUNT(id) FROM `site`")->queryScalar();
        $this->assertInternalType('int', $count);
        $this->assertNotEquals(0, $count, 'No sites.');
    }
}