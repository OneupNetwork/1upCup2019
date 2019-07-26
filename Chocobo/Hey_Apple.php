<?php
/**
 * Accepted
 * 46 ms    8 KB
 * Jul/26/2019 17:10UTC+8
 */
$step = fgets(STDIN);
$line = fgets(STDIN);
$line = explode(' ', $line);

$list = [];
$max = 1;
$last = 1;

for ($i=0;$i<$step;$i++) {
    $a = (int)$line[$i];
    $b = isset($line[$i+1]) ? (int)$line[$i+1] : false;

    if ($b === 1 || $b === false) {
        $list[] = $a;
    }
}

echo count($list).PHP_EOL;
echo join(' ', $list) . PHP_EOL;