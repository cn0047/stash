<?

if (isset($argv[1]))
    define('PROJECT', $argv[1]);
else
    die('Set project argument and try again. [m,a]');

// PhpStorm
if (PROJECT == 's') {
    echo "PhpStorm\n";
    $dir = '/home/kovpak/web/ide/';
    $sub_dir = 'PhpStorm-131.374';
    if (is_dir($dir.$sub_dir)) {
        echo 'success: '.rename($dir.$sub_dir, $dir.'_'.$sub_dir).", now {$sub_dir} is denied. \n";
    } elseif (is_dir($dir.'_'.$sub_dir)) {
        echo 'success: '.rename($dir.'_'.$sub_dir, $dir.$sub_dir).", now {$sub_dir} is allowed. \n";
    }
    exit();
}

// AFFILIATES
if (PROJECT == 'a') {
    echo "AFFILIATES\n";
    define('DIR', '/home/kovpak/web/kovpak/aff/');
    //.htaccess
    $file = DIR . '.htaccess';
    $text = file_get_contents($file);
    if (strpos($text, 'php_value display_errors on') !== false) {
        $text = str_replace('php_value display_errors on', 'php_value display_errors off', $text);
    } else {
        $text = str_replace('php_value display_errors off', 'php_value display_errors on', $text);
    }
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    //.htaccess help_scripts
    $file = DIR . 'help_scripts/.htaccess';
    $text = file_get_contents($file);
    if (strpos($text, '#deny from all') !== false) {
        $text = str_replace('#deny from all', 'deny from all', $text);
    } else {
        $text = str_replace('deny from all', '#deny from all', $text);
    }
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    // debug functions
    $file = DIR . 'include/debug/functions.php';
    $text = file_get_contents($file);
    if (strpos($text, '//set_error_handler("error_handler");') !== false) {
        $text = str_replace('//set_error_handler("error_handler");', 'set_error_handler("error_handler");', $text);
    } else {
        $text = str_replace('set_error_handler("error_handler");', '//set_error_handler("error_handler");', $text);
    }
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";

    // MainPageAffiliate - it's abort autologin
    /*
    $file = DIR . '/include/QUnit/UI/MainPageAffiliate.class.php';
    $text = file_get_contents($file);
    if (strpos($text, "//goto_url('/cpa/');") !== false) {
        $text = str_replace("//goto_url('/cpa/');", "goto_url('/cpa/');", $text);
    } else {
        $text = str_replace("goto_url('/cpa/');", "//goto_url('/cpa/');", $text);
    }
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    */

    // Auth.class.php - autologin in merchant
    /*
    $file = DIR . 'include/QCore/Auth.class.php';
    $text = file_get_contents($file);
    if (strpos($text, '//$hash = get_autologin_hash($hash);') !== false) {
        $text = str_replace('//$hash = get_autologin_hash($hash);', '$hash = get_autologin_hash($hash);', $text);
    } else {
        $text = str_replace('$hash = get_autologin_hash($hash);', '//$hash = get_autologin_hash($hash);', $text);
    }
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    */

    // left_menu.tpl.php
    /*
    $file = DIR . 'templates/merchant/default/left_menu.tpl.php';
    $text = file_get_contents($file);
    $str1 = '<div class="leftMenuTableOpened" id="left_menu_{$group}">';
    $str2 = '<div class="leftMenuTableClosed" id="left_menu_{$group}">';
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";

    // functions
    $file = DIR . 'include/debug/functions.php';
    $text = file_get_contents($file);
    $str1 = "define('ALLOW_DEBUG', true); require_once '/home/kovpak/Документы/storage/projects/functions.php';";
    $str2 = "define('ALLOW_DEBUG', true);";
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";*/
}
// MULTISITE
if (PROJECT == 'm') {
    echo "MULTISITE\n";
    define('DIR', '/home/kovpak/web/kovpak/multi_web/');
    // .htaccess
    $file = DIR . '.htaccess';
    $text = file_get_contents($file);
    $str1 = 'php_value display_errors on';
    $str2 = 'php_value display_errors off';
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    /*
    // functions
    $file = DIR . 'inc/debug.inc.php';
    $text = file_get_contents($file);
    $str1 = "require_once SITE_PATH . 'configs/debug_config.php'; require_once '/home/kovpak/Документы/storage/projects/functions.php';";
    $str2 = "require_once SITE_PATH . 'configs/debug_config.php';";
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    */
    // inc/debug.inc.php
    $file = DIR . 'inc/debug.inc.php';
    $text = file_get_contents($file);
    $str1 = 'pr($errType.$errArrMAIL["Error message"],$errArrMAIL,$errTimeReal,$errTimeLog,$sending);';
    $str2 = 'send_debug_mail($errType,$errArrMAIL,$errTimeReal,$errTimeLog,$sending);';
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
}

// PHOENIX
if (PROJECT == 'p') {
    echo "PHOENIX\n";
    define('DIR', '/home/kovpak/web/kovpak/phoenix/');
    $file = DIR . 'index.php';
    $text = file_get_contents($file);
    $str1 = "defined('YII_DEBUG') or define('YII_DEBUG', false);";
    $str2 = "defined('YII_DEBUG') or define('YII_DEBUG', true);";
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
}

// PHOENIX 2
if (PROJECT == 'px') {
    echo "PHOENIX\n";
    define('DIR', '/home/kovpak/web/kovpak/px/');
    // index.php
    $file = DIR . 'index.php';
    $text = file_get_contents($file);
    $str1 = "defined('YII_DEBUG') or define('YII_DEBUG', Constants::YII_DEBUG);";
    $str2 = "defined('YII_DEBUG') or define('YII_DEBUG', 1);";
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    // .htaccess
    // $file = DIR . '.htaccess';
    // $text = file_get_contents($file);
    // $str1 = "</Files>Allow from all";
    // $str2 = "</Files>";
    // if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    // else $text = str_replace($str2, $str1, $text);
    // $str1 = file_put_contents($file, $text) ? '+' : '-';
    // echo 'success: ' . $str1 . ' | ' . $file . "\n";

    // dataCollector
    $file = DIR . 'protected/config/dataCollector.php';
    $text = file_get_contents($file);
    $str1 = "'eAffiliates'                => ['priority' => 'urgent',";
    $str2 = "'eAffiliates'                => [";
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    // SQL Log 1
    $file = DIR . 'protected/config/main.php';
    $text = file_get_contents($file);
    $str1 = "'categories' =>'system.db.*', // tail -f protected/runtime/debug.tmp.sql.log";
    $str2 = "'categories' => 'payment.*',";
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
    // SQL Log 2
    $file = DIR . 'protected/config/main.php';
    $text = file_get_contents($file);
    $str1 = "'logFile'    => 'debug.tmp.sql.log',";
    $str2 = "'logFile'    => 'payment.log',";
    if (strpos($text, $str1) !== false) $text = str_replace($str1, $str2, $text);
    else $text = str_replace($str2, $str1, $text);
    $str1 = file_put_contents($file, $text) ? '+' : '-';
    echo 'success: ' . $str1 . ' | ' . $file . "\n";
}

$file = '/usr/bin/geany_function_1';
$text = '#!/bin/bash'."\n".'s=`xargs -0 echo`'."\n".'geany '.DIR.'$s';
$str1 = file_put_contents($file, $text) ? '+' : '-';
echo 'success: ' . $str1 . ' | ' . $file . "\n";
$file = '/usr/bin/geany_function_2';
$text = '#!/bin/bash'."\n".'s=`xargs -0 echo`'."\n".'cd '.DIR.''."\n".'grep "function\s*\b$s\b" -rni --include=*php .';
$str1 = file_put_contents($file, $text) ? '+' : '-';
echo 'success: ' . $str1 . ' | ' . $file . "\n";
$file = '/usr/bin/geany_function_3';
$text = '#!/bin/bash'."\n".'s=`xargs -0 echo`'."\n".'find '.DIR.' -iname  "*$s*"';
$str1 = file_put_contents($file, $text) ? '+' : '-';
echo 'success: ' . $str1 . ' | ' . $file . "\n";

?>
