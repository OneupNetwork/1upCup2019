<?php
/**
 * Accepted
 * 46 ms    16 KB
 * Jul/26/2019 16:12UTC+8
 */
$a=trim(fgets(STDIN));
$b=explode(' ', trim(fgets(STDIN)));
$arr = array_count_values($b);
$ans1 = $arr[1];
echo $ans1.PHP_EOL;

for ($i = 1;$i<$a;$i++) {
    if($b[$i] == 1) {
        echo $b[$i-1].' ';
    }
}
echo $b[$a-1];