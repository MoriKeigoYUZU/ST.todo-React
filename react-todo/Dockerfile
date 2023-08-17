# ベースとなるイメージを指定
FROM node:14

# アプリケーションディレクトリを作成
WORKDIR /app

# アプリケーションの依存関係をコピー
COPY package*.json ./

# 依存関係をインストール
RUN npm install

# アプリケーションのソースをコピー
COPY . .

# アプリケーションを起動
CMD ["npm", "start"]