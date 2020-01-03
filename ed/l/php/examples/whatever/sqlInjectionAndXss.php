<?php
/**
 * @example php -S localhost:8000 ed/php/examples/security/sqlInjection.php
 */

$dbh = new PDO('mysql:host=127.0.0.1;dbname=test', 'root');
$sth = $dbh->prepare("SHOW TABLES LIKE 'user'");
if (!$sth->execute()) {
    throw new Exception($sth->errorInfo());
}
if (empty($sth->fetchAll(PDO::FETCH_ASSOC))) {
    echo '<h1>INIT DB.</h1>';
    $sth = $dbh->prepare("
        CREATE TABLE IF NOT EXISTS user (
            id int auto_increment,
            name varchar(50),
            email varchar(50) unique,
            primary key (id)
        );
        INSERT IGNORE INTO user VALUES
            (NULL, 'James Bond', 'JamesBond@mi6.com')
            ,(NULL, 'Miss Moneypenny', 'Moneypenny@mi6.com')
            ,(NULL, 'Q', 'Q@mi6.com')
        ;
        CREATE TABLE IF NOT EXISTS userAccessLog (
            id int auto_increment,
            name varchar(250),
            primary key (id)
        );
    ");
}
if (!$sth->execute()) {
    throw new Exception($sth->errorInfo());
}
?>
<html>
    <body>
        <h4>Mi 6</h4>
        <!--
            Try find:
                James Bond
                Or: '; insert into user set name = 'Silva';
                Or: ' union all select name from user -- 
                Or: <script>alert(404)</script>
            But not:
                q'; drop table user; -- 
        -->
        <form>
            <label for="value">Name:</label>
            <input type="text" name="name" size="100" value="<?= isset($_GET['name']) ? $_GET['name'] : '' ?>">
            <br>
            <input type="submit" value="Search">
        </form>
        Result:
        <hr>
        <pre>
    </body>
</html>
<?php

if (isset($_GET['name']) and !empty($_GET['name'])) {
    // Way 1.
    $sql = "SELECT name From user WHERE name = '{$_GET['name']}'";
    $sth = $dbh->prepare($sql);

    // Way 2.
    // $sql = "SELECT name FROM user WHERE name = :name";
    // $sth = $dbh->prepare($sql);
    // $sth->bindValue(':name', $_GET['name'], PDO::PARAM_STR);

    if (!$sth->execute()) {
        throw new Exception($sth->errorInfo());
    }
    echo '<script>console.log('.json_encode($sql).')</script>';
    $result = $sth->fetchAll(PDO::FETCH_ASSOC);
    var_export($result);

    // Write to log.
    $sql = "INSERT INTO userAccessLog SET name = :name";
    $sth = $dbh->prepare($sql);
    $sth->bindValue(':name', $_GET['name'], PDO::PARAM_STR);
    if (!$sth->execute()) {
        throw new Exception($sth->errorInfo());
    }
    // Get from log.
    echo '<p>User access log:</p><hr>';
    $sql = 'SELECT * FROM userAccessLog ORDER BY id DESC';
    $sth = $dbh->prepare($sql);
    if (!$sth->execute()) {
        throw new Exception($sth->errorInfo());
    }
    $result = $sth->fetchAll(PDO::FETCH_ASSOC);
    // Prevent undesired behaviour.
    // array_walk_recursive($result, function (&$v) {
    //     $v = htmlentities($v, ENT_QUOTES, 'UTF-8');
    // });
    var_export($result);
}
