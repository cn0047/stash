<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/vendor/sproutvideo/sproutvideo/lib/SproutVideo/Autoloader.php';

SproutVideo_Autoloader::register();
SproutVideo::$api_key = '6efe7f6b4a2a173eb0899524cd8cbf14';

// var_export(SproutVideo\Video::create_video(
//   '/home/kovpak/Downloads/video1.mp4',
//   ['privacy' => 2, 'notification_url' => 'http://52.210.108.93?id=7']
// ));
$videoId = '709bd0be191debcaf8';
$securityToken = 'ff1d11ff40e6eba0';
// var_export(SproutVideo\Video::list_videos());
// var_export(SproutVideo\Video::get_video($videoId));
// var_export(SproutVideo\Video::delete_video($videoId));

// var_export(SproutVideo\Login::create_login(array('email' => 'test@example.com', 'password' => 'thisisthepassword')));
$loginId = '1f9adfb01c1e92';
// var_export(SproutVideo\Login::list_logins());

// var_export(SproutVideo\AccessGrant::create_access_grant(array('video_id' => $videoId, 'login_id' => $loginId)));

var_export(SproutVideo\Video::signed_embed_code($videoId, $securityToken, array('type' => 'hd')));
