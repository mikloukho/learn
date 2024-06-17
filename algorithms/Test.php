<?php

class Euclidean {
    public function __invoke(int $a, int $b): ?int
    {
        if($b === 0){
            return $a;
        }
        $r = $a % $b;
        return $this($b, $r);
    }
}
$binarySearch = new Euclidean();
$idx = $binarySearch(270, 192);
print_r($idx);

