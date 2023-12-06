# Gossip Glomers

Practicing distributed compute via Fly.io's Gossip Glomer workshop: https://fly.io/dist-sys/

The binary can be built with `make gossip-glomers`.

To run the tests, follow the instructions on downloading Maelstrom. Then from the
download directory, run the various test commands:

```bash
# 1) echo
./maelstrom test -w echo \
    --node-count 1 \
    --time-limit 10 \
    --bin ~/code/bradschwartz/gossip-glomers/gossip-glomers

# 2) generate unique ids
./maelstrom test -w unique-ids \
    --time-limit 30 \
    --rate 1000 \
    --node-count 3 \
    --availability total \
    --nemesis partition \
    --bin ~/code/bradschwartz/gossip-glomers/gossip-glomers
```
