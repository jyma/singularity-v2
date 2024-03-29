# 暗号化

## 概要

Singularityには、提供された受信者（公開鍵）やYubiKeysなどのハードウェアPIVトークンを使用してファイルを暗号化する組み込みの暗号化ソリューションがあります。また、外部の暗号化ソリューションに統合するためのカスタム暗号化スクリプトを指定することもできます \[テストが必要]。

## 組み込みの暗号化

まず、非対称暗号化のための公開-秘密キーペアを作成します。Singularityが使用している基礎となる暗号化ライブラリは[age](https://github.com/FiloSottile/age)と呼ばれています。

```sh
go install filippo.io/age/cmd/...@latest
age-keygen -o key.txt
> Public key: agexxxxxxxxxxxx
```

ここで、生成された公開鍵を使用して、各ファイルを暗号化するためのデータセットをセットアップすることができます。

```sh
singularity dataset create --encryption-recipient agexxxxxxxxxxxx \
  --output-dir . test
```

インライン準備は無効になっています。同じファイルを2度暗号化することはできません。暗号化プロセス中に初期のランダム性が導入されるため、異なる暗号化された内容が生成されます。

その後、データソースを追加してデータの準備プロセスを続けることができます。ただし、フォルダ構造は暗号化されないため、フォルダ構造のDAGを生成するか、`daggen`コマンドを実行しないかを選択する必要があります。後者の場合、フォルダ構造はSingularityデータベースとコマンドからのみアクセス可能です。

## カスタム暗号化

Singularityでは、カスタムスクリプトを指定してファイルストリームを暗号化することでカスタム暗号化も提供されています。これは、キーマネージメントサービスやカスタム暗号化アルゴリズムやツールと組み合わせて使用することができます。