<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/vendor/sproutvideo/sproutvideo/lib/SproutVideo/Autoloader.php';

SproutVideo_Autoloader::register();
SproutVideo::$api_key = '6efe7f6b4a2a173eb0899524cd8cbf14';

var_export(SproutVideo\Video::create_video(
  '/home/kovpak/Downloads/video1.mp4',
  ['privacy' => 2, 'notification_url' => 'http://52.210.108.93?id=7']
));
$videoId = '709bd0bd191beecaf8';
$securityToken = '9d93d3528706568d';
// var_export(SproutVideo\Video::list_videos());
// var_export(SproutVideo\Video::get_video($videoId));
var_export(SproutVideo\Video::delete_video($videoId));

// var_export(SproutVideo\Login::create_login(array('email' => 'test@example.com', 'password' => 'thisisthepassword')));
$loginId = '1f9adfb01c1e92';
// var_export(SproutVideo\Login::list_logins());

// var_export(SproutVideo\AccessGrant::create_access_grant(array('video_id' => $videoId, 'login_id' => $loginId)));

// var_export(SproutVideo\Video::signed_embed_code($videoId, $securityToken, array('type' => 'hd')));

$video = array (
  'id' => '709bd0bd191beecaf8',
  'width' => 720,
  'height' => 1280,
  'embed_code' => '<iframe class=\'sproutvideo-player\' src=\'//videos.sproutvideo.com/embed/709bd0bd191beecaf8/9d93d3528706568d?type=sd\' width=\'630\' height=\'1120\' frameborder=\'0\' allowfullscreen></iframe>',
  'source_video_file_size' => 8289699,
  'sd_video_file_size' => 0,
  'hd_video_file_size' => 0,
  'security_token' => '9d93d3528706568d',
  'title' => 'video2.mp4',
  'description' => NULL,
  'duration' => 5.4160000000000004,
  'privacy' => 0,
  'password' => NULL,
  'state' => 'deployed',
  'tags' => array (),
  'created_at' => '2016-09-23T13:23:01+01:00',
  'updated_at' => '2016-09-23T13:23:14+01:00',
  'plays' => 0,
  'progress' => 100,
  'requires_signed_embeds' => false,
  'selected_poster_frame_number' => 0,
  'embedded_url' => NULL,
  'assets' => array (
    'videos' => 
    array (
      '240p' => 'https://api-files.sproutvideo.com/file/709bd0bd191beecaf8/9d93d3528706568d/240.mp4',
      '360p' => 'https://api-files.sproutvideo.com/file/709bd0bd191beecaf8/9d93d3528706568d/360.mp4',
      '480p' => 'https://api-files.sproutvideo.com/file/709bd0bd191beecaf8/9d93d3528706568d/480.mp4',
      '720p' => 'https://api-files.sproutvideo.com/file/709bd0bd191beecaf8/9d93d3528706568d/720.mp4',
      '1080p' => NULL,
      '2k' => NULL,
      '4k' => NULL,
      '8k' => NULL,
      'source' => NULL,
    ),
    'thumbnails' => 
    array (
      0 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/thumbnails/frame_0000.jpg',
      1 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/thumbnails/frame_0001.jpg',
      2 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/thumbnails/frame_0002.jpg',
      3 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/thumbnails/frame_0003.jpg',
      4 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/thumbnails/frame_0004.jpg',
      5 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/thumbnails/frame_0005.jpg',
    ),
    'poster_frames' => 
    array (
      0 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/poster_frames/frame_0000.jpg',
      1 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/poster_frames/frame_0001.jpg',
      2 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/poster_frames/frame_0002.jpg',
      3 => 'https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/aecda4a7f0ccf82c81b3fbc40fb77e2f/poster_frames/frame_0003.jpg',
    ),
  ),
);

$login = array (
  'id' => '1f9adfb01c1e92',
  'email' => 'test@example.com',
  'access_grants' => 
  array (
  ),
  'created_at' => '2016-09-26T07:18:57+02:00',
  'updated_at' => '2016-09-26T07:18:57+02:00',
);

$accessGrant = array (
  'id' => 'd49fdcbb1618e0c35a',
  'video_id' => 'a09bd0bd1c1aebcf28',
  'access_starts_at' => NULL,
  'access_ends_at' => NULL,
  'allowed_plays' => 0,
  'play_count' => 0,
  'login_id' => '1f9adfb01c1e92',
  'created_at' => '2016-09-26T05:21:13Z',
  'updated_at' => '2016-09-26T05:21:13Z',
  'download_permissions' => 
  array (
  ),
);
?>

{
"id":"a09bd0bd1c1aebcf28",
"width":720,
"height":1280,
"embed_code":"\\u003Ciframe class=\'sproutvideo-player\' src=\'//videos.sproutvideo.com/embed/a09bd0bd1c1aebcf28/2555c38117ea4952?type=sd\' width=\'630\' height=\'1120\' frameborder=\'0\' allowfullscreen\\u003E\\u003C/iframe\\u003E",
"source_video_file_size":33198309,
"sd_video_file_size":0,
"hd_video_file_size":0,
"security_token":"2555c38117ea4952",
"title":"video1.mp4",
"description":null,
"duration":22.258,
"privacy":2,
"password":null,
"state":"deployed",
"tags":[
],
"created_at":"2016-09-26T05:06:54Z",
"updated_at":"2016-09-26T05:07:20Z",
"plays":0,
"progress":100,
"requires_signed_embeds":false,
"selected_poster_frame_number":0,
"embedded_url":null,
"assets":{
"videos":{
"240p":"https://api-files.sproutvideo.com/file/a09bd0bd1c1aebcf28/2555c38117ea4952/240.mp4",
"360p":"https://api-files.sproutvideo.com/file/a09bd0bd1c1aebcf28/2555c38117ea4952/360.mp4",
"480p":"https://api-files.sproutvideo.com/file/a09bd0bd1c1aebcf28/2555c38117ea4952/480.mp4",
"720p":"https://api-files.sproutvideo.com/file/a09bd0bd1c1aebcf28/2555c38117ea4952/720.mp4",
"1080p":null,
"2k":null,
"4k":null,
"8k":null,
"source":null
},
"thumbnails":[
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/thumbnails/frame_0000.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/thumbnails/frame_0001.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/thumbnails/frame_0002.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/thumbnails/frame_0003.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/thumbnails/frame_0004.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/thumbnails/frame_0005.jpg"
],
"poster_frames":[
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/poster_frames/frame_0000.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/poster_frames/frame_0001.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/poster_frames/frame_0002.jpg",
"https://images.sproutvideo.com/f5106b97e900c39b2f04472c77bf99dd/ff06e46771c0b2238ccda878974daf84/poster_frames/frame_0003.jpg"
]
}
}