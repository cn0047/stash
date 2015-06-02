<div class="form">
<?php echo CHtml::beginForm(); ?>
    <?php echo CHtml::errorSummary($model); ?>
    <div class="row">
        <?php echo CHtml::activeLabel($model, 'username'); ?>
        <?php echo CHtml::activeTextField($model, 'username') ?>
    </div>
    <div class="row">
        <?php echo CHtml::activeLabel($model, 'type'); ?>
        <?php echo CHtml::dropDownList('type', $model->type, ['question', 'answer']); ?>
    </div>
    <div class="row">
        <?php echo CHtml::activeLabel($model, 'post'); ?>
        <?php echo CHtml::activeTextArea($model, 'post') ?>
    </div>
    <div class="row submit">
        <?php echo CHtml::submitButton('Post'); ?>
    </div>
<?php echo CHtml::endForm(); ?>
</div>
