<?php
/**
 * Accepted
 * 46 ms    16 KB
 * Jul/26/2019 17:40UTC+8
 */
$number = trim(fgets(STDIN));
$eng1 = trim(fgets(STDIN));
$eng2 = trim(fgets(STDIN));

$orn = preg_split('//', $eng1, -1, PREG_SPLIT_NO_EMPTY);
$tar = preg_split('//', $eng2, -1, PREG_SPLIT_NO_EMPTY);

$result = array();
foreach($tar as $key => $eachT){
    if($eachT == $orn[$key]){
        continue;
    }
    $pos = strpos(implode('',$orn),$eachT,$key);
    if($pos === false){
        echo -1;
        exit;
    }

    $arrF = array_slice($orn,0,$key);
    $arrB = array_slice($orn,$key);


    $new = array();
    foreach($arrF as $aaa){
        $new[]=$aaa;
    }
    $new[] = $eachT;
    foreach($arrB as $aaa){
        $new[]=$aaa;
    }
    $orn = $new;
    unset($orn[$pos+1]);

    for(;$pos>$key;--$pos){
        $result[] = $pos;
    }
}
if(implode('',$orn) != implode('',$tar))
    echo -1;
else{
    echo count($result).PHP_EOL.implode(' ',$result);
}