<?php

$descricao = addslashes(trim($_POST['descricao']));
$quantidade = addslashes(trim($_POST['quantidade']));
$valor = addslashes(trim($_POST['valor']));

//echo "$descricao, $quantidade, $valor";

$LastID = rand(1, 1000);

$url = "http://localhost:3000/produtos";

$client = curl_init($url);

if($client != null) {

  curl_setopt($client, CURLOPT_RETURNTRANSFER, 1);

  $response = curl_exec($client);
  
  $rs = json_decode($response);
  
  $ult = $rs[count($rs) - 1];
  
  //var_dump($ult);
  
  $LastID = $ult->id + 1;
  
  if($rs == null) {
    die("ERRO 404");
  }

} else {
  die("ERRO 404");
}

if(!empty($descricao) && !empty($quantidade) && !empty($valor)) {

$url = "http://localhost:3000/produtos";

$client = curl_init($url);

if($client != null){
  curl_setopt($client, CURLOPT_RETURNTRANSFER, 1);

$dados = array (
  'id' => intval($LastID),
  'descricao' => $descricao,
  'quantidade' => intval($quantidade),
  'valor' => floatval($valor)
);

$novoProduto = json_encode($dados);

curl_setopt($client, CURLOPT_POST, true);

curl_setopt($client, CURLOPT_POSTFIELDS, $novoProduto);

curl_exec($client);

curl_close($client);

  header("location: index.php");
} else {
  echo "Erro 404";
}
  
} else {
  echo "informe os valores das variaveis";
}

 ?>
