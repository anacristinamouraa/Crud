<?php
   class Palavra {
   public $texto;
   public $dificuldade;

   }

   function criarPalavra($t, $c, $d) {
       $p = new Palavra();
       $p->texto = $t;
       $p->categoria = $c;
       $p->dificuldade = $d;
       retur $p;

   }

   function selecionar($palavras) {
       // a função rand gera um número aleatório
       $indice = rand(0, sizeof($palavras) - 1);
       return $palavras[$indice];

   }

   $arrayPalavras = [
    criarPalavra("Bola", "Esportes", "Fácil"),
    criarPalavra("Pato", "Animais", "Fácil"),
    criarPalavra("Ourives", "Profissão", "Difícil")

   ];

   echo "Palavra selecionada" , selecionar($arrayPalavras)->texto;

   ?>

 //resposta da pergunta
   
Pelo que entendi e pesquisei também no inicio tem uma classe chamada Palavra. as Palavras tem 3 propriedades: Texto, Categoria e Dificuldade. 
A função criarPalavra está , como o nome diz, criando uma função que cria uma palavra e pra isso ela usa a classe Palavra.  ali no $p = new palavra()
Ou seja, nessa função de criarPalavra ela retorna uma $p palavra.

Na função selcionar ela recebe palavras e retorna cada palavra com um indice aleatorio baseado no numero de palavras. Ou seja se temos 3 palavras teremos 3 posições de palavras, se temos 20 palavras vamos ter palavras que vão do 0 ao 19.   Percebe o -1 no final do comando.

Depois é criado um array de palavras, cada uma com as 3 propriedades Texto, categoria e dificuldade. 

No final a função nativa do PHP "echo",  imprime na tela o texto. "Palavra selecionada: ",  + o texto da palavra selecionada. 
Seria algo assim: Palavra selecionada: Bola por exemplo