<?php
/**
 * Accepted
 * 46 ms    16 KB
 * Jul/26/2019 16:48UTC+8
 */
$strlen = (int) fgets(STDIN);
$s = trim(fgets(STDIN));
$t = trim(fgets(STDIN));
$move =[];

for ($i = 0; $i < $strlen; $i++) {

    if ($s[$i] === $t[$i]) {
        continue;
    }

    $pos = strrpos($s, $t[$i]);

    if ($pos < $i||$pos === false) {
        echo -1, PHP_EOL;
        exit;
    }

    for (; $pos > $i; $pos--) {
        $tmp = $s[$pos -1];
        $s[$pos - 1] = $s[$pos];
        $s[$pos] = $tmp;
        $move[] = $pos;
    }
}
echo count($move), PHP_EOL;
echo implode(" ", $move), PHP_EOL;