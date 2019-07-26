<?php
/**
 * Accepted
 * 31 ms    4 KB
 * Jul/26/2019 16:56UTC+8
 */
$number = trim(fgets(STDIN));
$w = explode(' ',trim(fgets(STDIN)));
//print_r($w);
$list = array();
$i=-1;
foreach($w as $each){
    if($each == 1) {
        $i++;
    }
    $list[$i][]=$each;
}

$step = count($list);
$datas = array();

foreach($list as $each){
    $n = array_pop($each);

    if($n>0)
        $datas[] = $n;
    else
        $datas[] =1;
}

echo $step.PHP_EOL.implode(' ',$datas);