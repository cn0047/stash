<?php

class Car extends CActiveRecord
{
    public function tableName()
    {
        return 'car';
    }

    public function relations()
    {
        return [
            'brand' => [self::BELONGS_TO, 'Brand', 'brand_id'],
        ];
    }

    public function rules()
    {
        $helper = new Helper;
        return [
            ['brand_id, model, maxSpeed', 'required', 'on' => 'insert'],
            ['brand_id', 'in', 'range' => array_keys($helper->getAvailableBrands())],
            ['model', 'type', 'type' => 'string'],
            ['model', 'length', 'max' => 100],
            ['maxSpeed', 'type', 'type' => 'integer'],
        ];
    }

    public function search()
    {
        $brand = new Brand;
        $criteria = new CDbCriteria;
        $criteria->alias = 'c';
        $criteria->join = 'INNER JOIN '.$brand->tableName().' AS b on c.brand_id = b.id';
        if (!empty($this->brand_id)) {
            $criteria->compare('brand_id', $this->brand_id, true);
        }
        $criteria->compare('model', $this->model, true);
        $criteria->compare('maxSpeed', $this->maxSpeed, true);
        return new CActiveDataProvider(
            $this,
            [
                'criteria' => $criteria,
                'pagination' => ['pageSize' => 20],
                'sort' => ['defaultOrder'=>'maxSpeed DESC'],
            ]
        );
    }

    public function searchAdd()
    {
        $brand = new Brand;
        $sql = sprintf(
            '
                SELECT
                    c.id,
                    b.name AS brand_name,
                    b.country,
                    c.model,
                    c.maxSpeed
                FROM %s AS c
                INNER JOIN %s AS b on c.brand_id = b.id
            ',
            $this->tableName(),
            $brand->tableName()
        );
        $rawData = Yii::app()->db->createCommand($sql);
        $count = Yii::app()->db->createCommand("SELECT COUNT(*) FROM ($sql) AS count")->queryScalar();
        $dataProvider = new CSqlDataProvider(
            $rawData,
            [
                'keyField' => 'id',
                'totalItemCount' => $count,
                'pagination' => ['pageSize' => 20],
                'sort' => ['defaultOrder'=>'maxSpeed DESC'],
            ]
        );
        return $dataProvider;
    }
}
