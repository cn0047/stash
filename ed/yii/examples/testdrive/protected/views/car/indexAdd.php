<?php echo CHtml::errorSummary($model); ?>
<?php $this->widget('zii.widgets.grid.CGridView', array(
    'dataProvider' => $model->searchAdd(),
    'ajaxUpdate' => false,
    'columns' => [
        ['header' => 'Id', 'value' => '$data["id"]'],
        ['header' => 'Brand', 'value' => '$data["brand_name"]'],
        ['header' => 'country', 'value' => '$data["country"]'],
        ['name' => 'model', 'value' => '$data["model"]'],
        ['name' => 'maxSpeed', 'value' => '$data["maxSpeed"]'],
        [
            'header' => 'Is fast',
            'value' => function ($data) {
                return $data['maxSpeed'] > 250 ? 'yes' : '';
            },
        ],
    ]
)); ?>
