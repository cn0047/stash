<?php

if (!empty($toSave)) {
    $keyStackAdditional = 'affiliatesTransactionStackAdditional';
    $stackAdditional = Yii::app()->cache->get($keyStackAdditional);
    if (is_array($stackAdditional)) {
        $toSave = array_merge($stackAdditional, $toSave);
    }
    $keyStackLock = 'affiliatesTransactionStackLock';
    $lock = Yii::app()->cache->get($keyStackLock);
    if ($lock) {
        Yii::app()->cache->set($keyStackAdditional, $toSave, 7*24*3600);
    } else {
        Yii::app()->cache->set($keyStackLock, true, 7*24*3600);
        $keyStack = 'affiliatesTransactionStack';
        $stack = Yii::app()->cache->get($keyStack);
        if (is_array($stack)) {
            $toSave = array_merge($stack, $toSave);
        }
        Yii::app()->cache->set($keyStack, $toSave, 7*24*3600);
        Yii::app()->cache->delete($keyStackAdditional);
        Yii::app()->cache->delete($keyStackLock);
}

function getLeadsSales($count)
{
   $keyStack = 'affiliatesTransactionStack';
   $stack = Yii::app()->cache->get($keyStack);
   if (!is_array($stack)) {
       return [];
   }
   $slice = array_slice($stack, 0, $count);
   $sliced = array_flip(array_keys($slice));
   $keyStackSliced = 'affiliatesTransactionStackSliced';
   Yii::app()->cache->set($keyStackSliced, $sliced, 7*24*3600);
   return array_values($slice);
}

function setStatusLeadsSales($count)
{
    $keyStackSliced = 'affiliatesTransactionStackSliced';
    $stackSliced = Yii::app()->cache->get($keyStackSliced);
    if (!is_array($stackSliced)) {
        return true;
    }
    $keyStackLock = 'affiliatesTransactionStackLock';
    $keyStack = 'affiliatesTransactionStack';
    do {
        $lock = Yii::app()->cache->get($keyStackLock);
        if (!$lock) {
            Yii::app()->cache->set($keyStackLock, true, 7*24*3600);
            $stack = Yii::app()->cache->get($keyStack);
            $stack = array_diff_key($stack, $stackSliced);
            Yii::app()->cache->set($keyStack, $stack, 7*24*3600);
            Yii::app()->cache->delete($keyStackSliced);
            Yii::app()->cache->delete($keyStackLock);
        }
    } while ($lock);
    return true;
}