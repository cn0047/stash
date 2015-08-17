<form action="">
    <?php echo CHtml::textField('name', $get['name']); ?>
    <?php echo CHtml::submitButton('Submit'); ?>
</form>
<?php $this->widget('zii.widgets.grid.CGridView', array(
    'dataProvider' => $dataProvider,
    'ajaxUpdate' => false,
    'columns' => [
        ['header' => 'Name', 'value' => '$data["name"]'],
        ['header' => 'Country', 'value' => '$data["country"]'],
    ]
)); ?>
