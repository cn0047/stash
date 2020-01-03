<div class="form">
<?php $form=$this->beginWidget('CActiveForm'); ?>
    <?php echo CHtml::errorSummary($model); ?>
    <div class="row">
        <?php echo $form->label($model, 'username'); ?>
        <?php echo $form->textField($model, 'username') ?>
    </div>
    <div class="row">
        <?php echo $form->label($model, 'type'); ?>
        <?php echo  $form->dropDownList($model, 'type', ['question', 'answer']); ?>
    </div>
    <div class="row">
        <?php echo $form->label($model, 'post'); ?>
        <?php echo $form->textArea($model, 'post') ?>
    </div>
    <div class="row">
        <?php echo $form->label($model, 'tags'); ?>
        <?php echo  $form->dropDownList(
            $model, 'tags', ['php', 'mysql', 'js', 'jquery', 'nodejs'], ['multiple' => true]
        ); ?>
    </div>
    <div class="row submit">
        <?php echo CHtml::submitButton('Post'); ?>
    </div>
<?php $this->endWidget(); ?>
</div>
