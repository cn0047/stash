<?php
/**
 * Store encrypted data in mysql.
 *
 * Simplest example how store encrypted data at MySql through PDO.
 *
 * @author Vladimir Kovpak <cn007b@gmail.com>
 */

class Crypto
{
    const CIPHER = MCRYPT_RIJNDAEL_128;
    const MODE   = MCRYPT_MODE_CBC;

    private $key;

    /**
     * Crypto constructor.
     *
     * @param string $key Key with which the data will be encrypted.
     * @throws \InvalidArgumentException When key not string or empty.
     * @return void.
     */
    public function __construct($key)
    {
        if (!is_string($key)) {
            throw new \InvalidArgumentException('Key should be string.');
        }
        if (empty($key)) {
            throw new \InvalidArgumentException("Key can't be empty.");
        }
        $this->key = $key;
    }

    /**
     * Encrypts data.
     *
     * @param string $text Data to be encrypted.
     * @throws \InvalidArgumentException When $text not string.
     * @return string Encrypted data.
     */
    public function encrypt($text)
    {
        if (!is_string($text)) {
            throw new \InvalidArgumentException('Text should be string.');
        }
        $ivSize = mcrypt_get_iv_size(self::CIPHER, self::MODE);
        $iv = mcrypt_create_iv($ivSize, MCRYPT_DEV_RANDOM);
        $encryptedData = mcrypt_encrypt(self::CIPHER, $this->key, $text, self::MODE, $iv);
        return base64_encode($iv.$encryptedData);
    }

    /**
     * Decrypts data.
     *
     * @param string $encryptedData Data to be decrypted.
     * @throws \InvalidArgumentException When $encryptedData not string.
     * @throws \RuntimeException When iv initialization failed.
     * @return string Decrypted data.
     */
    public function decrypt($encryptedData)
    {
        if (!is_string($encryptedData)) {
            throw new \InvalidArgumentException('Encrypted data should be string.');
        }
        $encryptedData = base64_decode($encryptedData);
        $ivSize = mcrypt_get_iv_size(self::CIPHER, self::MODE);
        if (strlen($encryptedData) < $ivSize) {
            throw new \RuntimeException('IV initialization failed.');
        }
        $iv = substr($encryptedData, 0, $ivSize);
        $encryptedData = substr($encryptedData, $ivSize);
        $text = mcrypt_decrypt(self::CIPHER, $this->key, $encryptedData, self::MODE, $iv);
        return rtrim($text, "\0");
    }
}

class Db
{
    private $dbh;

    public function __construct()
    {
        $this->dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
        $sth = $this->dbh->prepare('CREATE TABLE IF NOT EXISTS cr (id int auto_increment, data text, primary key (id))');
        $sth->execute();
    }

    public function __destruct()
    {
        $sth = $this->dbh->prepare('DROP TABLE cr');
        $sth->execute();
    }

    public function save($data)
    {
        $sth = $this->dbh->prepare('INSERT INTO cr SET data = :data');
        $sth->bindParam(':data', $data, PDO::PARAM_STR);
        $sth->execute();
        return $this->dbh->lastInsertId();
    }

    public function get($id)
    {
        $sth = $this->dbh->prepare('SELECT * FROM cr WHERE id = :id');
        $sth->bindParam(':id', $id, PDO::PARAM_INT);
        $sth->execute();
        $r = $sth->fetch(PDO::FETCH_ASSOC);
        return $r['data'];
    }
}

$text = '0012300qweflnsdf(**&)*^T(*lkmasdfl;jasdlkf&^%)(MldmflsdjfLKJPOJKJN';
$crypt = new Crypto('23c34eWrg56fSdrt');
$encryptedString = $crypt->encrypt($text);
$db = new Db();
$id = $db->save($encryptedString);
$stringFromDb = $db->get($id);
$decryptedString = $crypt->decrypt($stringFromDb);
var_dump($text === $decryptedString);
