
### Name
toriniku  
Go * Gin * GORM → WEB_APPLICATIOM  
Toriniku_Price_Getter  

### Description
鶏肉の最安値を表示するWEBアプリケーション  
GinによるMVCアーキテクチャー  

### Requirement
Go  
mysql  

### Usage
```
git clone https://github.com/funasedaisuke/toriniku.git

cd toriniku

<!-- 以下のコマンドを実行して"shared-network"があることを確認 -->
docker network ls

<!-- 上記のネットワークがなければ以下のコマンドを実行 -->
docker network create shared-network

docker-compose up -d

<!-- goをビルド -->
docker-compose build

<!-- トップページを開く -->
open http://localhost:8000/top
```

[funasedaisuke](https://github.com/funasedaisuke/)  
