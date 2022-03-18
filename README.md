### Clickbait Classifier
This is simple Naive Bayesian classifier which classifies if the given title is clickbait or not. It's a golang imlementation.

DATASET: https://www.kaggle.com/amananandrai/clickbait-dataset/

#### Running the code -
```
go mod tidy
go run main.go
```

To check your titles - call `check(title)`. Sample return value `CB, 0.99999` ("CB/NONCB", Score)