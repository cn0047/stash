<?php echo CHtml::button('Edit', array('class' => 'edit')); ?>
<?php

Yii::import('zii.widgets.grid.CDataColumn');

class ModelColumn extends CDataColumn
{
    public function renderDataCellContent($row, $data)
    {
        return parent::renderDataCellContent($row, $data);
    }
}

$this->widget('zii.widgets.grid.CGridView', array(
    'dataProvider' => $model->search(),
    'filter' => $model,
    'columns' => [
        ['name' => 'id', 'value' => '$data->id'],
        ['name' => 'brand', 'value' => '$data->brand', 'filter' => $model->getDistinctBrands()],
        ['name' => 'model', 'value' => '$data->model', 'filter' => $model->getDistinctModels(), 'class' => 'ModelColumn'],
        ['name' => 'maxSpeed', 'value' => '$data->maxSpeed'],
        [
            'header' => 'is fast',
            'value' => function ($data) {
                return $data->maxSpeed > 250 ? 'yes' : '';
            },
        ],
        [
            'class' => 'CButtonColumn',
            'template' => '{update}',
            'buttons' => [
                'update' => [
                    'url' => 'Yii::app()->createUrl("car/edit", array("id" => $data->id))',
                    'options' => ['title' => "Update car"]
                ],
            ],
        ],
    ]
)); ?>
<?php
$js = <<<"HEREDOC"
$(function () {
    editHandler = function () {
        console.log('changed');
    }
    $('.edit').on('click', function () {
        $('table.items tbody tr td:nth-child(2)').each(function (k, v) {
            var selected = $(v).html();
            var html$ = $('select[name="Car[brand]"]')
                .clone()
                .attr('name', '')
                .attr('class', 'gridEdit')
                ;
            html$.find('option[value="'+selected+'"]').attr('selected', true);
            html$.on('change', editHandler);
            $(v).html(html$);
        });
    });
});
HEREDOC;
Yii::app()->clientScript->registerScript('edit', $js, CClientScript::POS_END);
?>