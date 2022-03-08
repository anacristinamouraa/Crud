<?php
$encontrados = 0;
for ($i = 0; $i < 10; $i++) {
    if ($i % 4 = 0 && $i % 5 = 0) {
        echo $i, PHP_EOL
        $encontrados++;

    }
?>

//Acho que esse aqui o erro está no if, ele mostra que da forma que estava não tinha como ser múltiplos de 4 e 5, por isso inseri um = 
pois o múltiplo de 4 e 5 ele vai da resto 0, por isso ele não pega os que não são.

