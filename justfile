build:
    go build -o bin/api .

serve:
    ./bin/api

deploy:
    git push dokku main