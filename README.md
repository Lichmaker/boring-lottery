## server 启动
```
docker build -f ./Dockerfile-server -t boring-lottery-server .
docker run --name="boring-lottery-server" --rm -p 8002:8002 -d boring-lottery-server
```

## 双色球开奖爬虫
```
docker build -f ./Dockerfile-crawler -t boring-lottery-crawler .
docker run --name="boring-lottery-crawler" boring-lottery-crawler
```

## 生成预测
```
docker build -f ./Dockerfile-generator -t boring-lottery-generator .
docker run --name="boring-lottery-generator" boring-lottery-generator
```
