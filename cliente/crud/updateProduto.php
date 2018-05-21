<?php

 $id = addslashes(trim($_POST['id']));
 $desc = addslashes(trim($_POST['descricao']));
 $qtd = addslashes(trim($_POST['quantidade']));
 $val = addslashes(trim($_POST['valor']));

//echo "$id, $desc, $qtd, $val";

if(!empty($id) && !empty($desc) && !empty($qtd) && !empty($val)) {


$url = "http://localhost:3000/usuarios";

$client = curl_init($url);

curl_setopt($client, CURLOPT_RETURNTRANSFER, 1);

$dados = array (
  'id' => rand(1, 1000),
  'nome' => 'Nathan Harper'
);

$novoUsuario = json_encode($dados);

curl_setopt($client, CURLOPT_POST, true);

curl_setopt($client, CURLOPT_POSTFIELDS, $novoUsuario);

curl_exec($client);

curl_close($client);

$url = "http://localhost:3000/produtos/". $id;

  $cliente = curl_init($url);

  if($cliente != null) {

    $dados = array (
      'id' => intval($id),
      'descricao' => $desc,
      'quantidade' => intval($qtd),
      'valor' => floatval($val)
    );

    $novoProduto= json_encode($dados);

    curl_setopt($cliente, CURLOPT_CUSTOMREQUEST, "PUT");

    curl_setopt($cliente, CURLOPT_POSTFIELDS, $novoProduto);

    curl_setopt($cliente, CURLOPT_RETURNTRANSFER, 1);

    curl_exec($cliente);

    curl_close($cliente);

    header("location: index.php");

} else {
  echo "Erro 404";
}
}else {
  echo "informe o valor ";
}
 ?>
