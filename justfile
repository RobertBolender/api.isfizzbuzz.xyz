build:
    go build -o bin/api .

serve:
    just build
    ./bin/api

deploy:
    git push dokku main

test:
    hurl --test test.hurl
