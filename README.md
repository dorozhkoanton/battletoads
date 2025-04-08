# battletoads
battletoads


## Usage
```
docker build --network host --progress plain --tag battletoads .
docker run --rm -it -e "API_ID=00000" -e "API_HASH=abcdef0123456789" battletoads ash
./app
```
