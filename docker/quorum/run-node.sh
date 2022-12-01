#!/bin/bash

ID=1

PEER1IP=45.21.109.101

cp -R ${PWD}/nodes/template/Node-${ID} ${PWD}/nodes

docker run -d \
-e NODE_NUM=${ID} \
-e HTTP_PORT=22000 \
-e WS_PORT=32000 \
-e RAFT_PORT=63000 \
-e PORT=30300 \
-e VERBOSITY=5 \
-e PEER_IP=$PEER1IP \
-e PEER_PORT=30301 \
-e PEER_RAFT_PORT=63001 \
-p 2200${ID}:22000 -p 3200${ID}:32000 -p 6300${ID}:63000 -p 3030${ID}:30300 \
-v ${PWD}/nodes/Node-${ID}/data:/root/Node-${ID}/data \
--name node-${ID} \
test-node:latest

