<?php
/**
 * Accepted
 * 46 ms    16 KB
 * Jul/26/2019 16:22UTC+8
 */
$list = array(
    'a'=>1,
    'b'=>2,
    'c'=>3,
    'd'=>4,
    'e'=>5,
    'f'=>6,
    'g'=>7,
    'h'=>8,
    'i'=>9,
    'j'=>10,
    'k'=>11,
    'l'=>12,
    'm'=>13,
    'n'=>14,
    'o'=>15,
    'p'=>16,
    'q'=>17,
    'r'=>18,
    's'=>19,
    't'=>20,
    'u'=>21,
    'v'=>22,
    'w'=>23,
    'x'=>24,
    'y'=>25,
    'z'=>26,
    );



$number = explode(' ',trim(fgets(STDIN)));
$english = trim(fgets(STDIN));
$english = preg_split('//', $english, -1, PREG_SPLIT_NO_EMPTY);



sort($english);

$datas= array();

for($i=0 ; $i< $number[0] ;$i++){
    if(empty($datas)){
        $datas[] = $list[$english[$i]];
    }

    if($list[$english[$i]] - $datas[count($datas)-1] >= 2){
        $datas[] = $list[$english[$i]];
    }

    if(count($datas) >= $number[1]){
        break;
    }

}

if(count($datas) < $number[1])
    echo '-1';
else
    echo array_sum($datas);
