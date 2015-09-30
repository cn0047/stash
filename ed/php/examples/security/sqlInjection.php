<html>
    <body>
        <form>
            <input type="text" name="value" value="'; create table injection (id int); -- ">
            <input type="submit">
        </form>
        <pre>
    </body>
</html>

<?php
/**
 * @example php -S localhost:8000 ed/php/examples/security/sqlInjection.php
 */

$dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
$sth = $dbh->prepare("SELECT '{$_GET['value']}'");
$sth->execute();
$result = $sth->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
