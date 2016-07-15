<?php

$thumb = new Imagick();
$thumb->readImage('/home/kovpak/Downloads/bond.origin.jpg');
$thumb->resizeImage(320, 240, imagick::FILTER_UNDEFINED, true);
$thumb->writeImage('/home/kovpak/Downloads/bond.res.jpg');
$thumb->clear();
$thumb->destroy();
