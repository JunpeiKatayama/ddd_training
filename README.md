# ドメイン駆動サンプル

## 起動コマンド等

- 各種リソース起動

```bash
docker-compose up -d
```

- サーバ起動

```bash
go run src/main.go
```

- gqlgen

```bash
go run github.com/99designs/gqlgen generate
```

## パッケージ構成

```
api
└ src
　 ├ main.go
　 ├ application
　 │ ├ service
　 │ ├ model
　 │ └ graph - gqlgen 自動生成ファイル
　 ├ domain
　 │ ├ service
　 │ ├ model
　 │ └ repository
　 ├ infrastructure
　 │ └ repository
　 └ resources
　 　 └ schema
```

## 動作確認

```graphql
mutation creareRanking {
  createRanking(input: { name: "好きな動物さんランキング" }) {
    name
  }
}

query rankings {
  rankings {
    id
    name
    items {
      id
      name
    }
  }
}

mutation addItem {
  addItem(
    input: { ranking_id: "7d147e3e-3b45-11ee-ab86-1a976c28b008", name: "サイ" }
  ) {
    id
    name
    items {
      id
      name
      point
    }
  }
}
```

```json
{
  "data": {
    "rankings": [
      {
        "id": "7d147e3e-3b45-11ee-ab86-1a976c28b008",
        "name": "好きな動物さんランキング",
        "items": [
          {
            "id": "8d484fba-3b45-11ee-ab86-1a976c28b008",
            "name": "ゾウさん"
          },
          {
            "id": "937b2db2-3b45-11ee-ab86-1a976c28b008",
            "name": "キリンさん"
          },
          {
            "id": "992ef720-3b45-11ee-ab86-1a976c28b008",
            "name": "サイ"
          }
        ]
      }
    ]
  }
}
```
