<?php

$args = ['code' => 200, 'text' => 'OK'];

?>
<script>
    var args = JSON.parse('<?=json_encode($args);?>');
    console.log(args);
</script>
