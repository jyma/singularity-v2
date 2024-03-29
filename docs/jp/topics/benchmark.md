# ベンチマーク

EZ準備コマンドは、ベンチマークを実行するための簡単な方法を提供します。

## テストデータの準備

まず、ベンチマークのためにいくつかのデータを準備します。ディスクIO時間を考慮に入れないように、スパースファイルを利用しています。現時点ではCIDのデデュプリケーションは行っていないため、シンギュラリティはそれらをランダムなバイトとして扱います。

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

ディスクIO時間をベンチマークの一部として含めたい場合は、お好みの方法でランダムなファイルを作成することもできます。

```
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## ez-prepを実行する

EZ準備コマンドは、非常に少数のカスタマイズ可能な設定でローカルフォルダを準備するためにいくつかの内部コマンドを実行するシンプルなコマンドです。

#### インライン準備を使用したベンチマーク

インライン準備を使用すると、CARファイルをエクスポートする必要がなく、必要なメタデータをデータベースに保存する必要がありません。

```sh
time singularity ez-prep --output-dir '' ./dataset
```

#### メモリ内データベースを使用したベンチマーク

ディスクIOをさらに減らすために、メモリ内データベースを使用することもできます。

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

#### 複数のワーカーを使用したベンチマーク

すべてのCPUコアを利用するために、ベンチマークのために並行性フラグを設定することができます。ただし、各ワーカーは約4つのCPUコアを消費するため、適切に設定する必要があります。

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## 結果の解釈

以下のような結果が表示されます。

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

`real`は実際のクロック時間を表しています。さらにワーカーの並行性を増やすと、この数値はおそらく減少します。

`user`はユーザースペースでのCPU時間を表しています。`user`を`real`で割ると、プログラムが使用したCPUコアのおおよその数になります。より多くの並行性を使用しても、この数値にはほとんど影響がないことが多いです。

`sys`はディスクIOに費やされたカーネルスペースでのCPU時間を表しています。

## 比較

以下のテストはランダムな8Gファイルで実行されたものです

<table><thead><tr><th width="290">ツール</th><th width="178.33333333333331" data-type="number">クロック時間 (秒)</th><th data-type="number">CPU時間 (秒)</th><th data-type="number">メモリ (KB)</th></tr></thead><tbody><tr><td>インライン準備を使用したシンギュラリティ</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>インライン準備を使用しないシンギュラリティ</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>