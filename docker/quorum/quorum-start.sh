#!/bin/bash
set -e

NODE_NUM=${NODE_NUM}
DATA_DIR="Node-${NODE_NUM}/data"
STATIC_NODES_FILE="${DATA_DIR}/static-nodes.json"
PERMISSIONED_NODES_FILE="${DATA_DIR}/permissioned-nodes.json"
NETWORK_ID=$(cat ${DATA_DIR}/genesis.json | tr -d '\r' | grep chainId | awk -F " " '{print $2}' | awk -F "," '{print $1}')

if [ -z "${HTTP_PORT}" ]; then
    HTTP_PORT=22000
fi
if [ -z "${WS_PORT}" ]; then
    WS_PORT=32000
fi
if [ -z "${PORT}" ]; then
    PORT=30300
fi
if [ -z "${RAFT_PORT}" ]; then
    RAFT_PORT=53000
fi
if [ -z "${VERBOSITY}" ]; then
    VERBOSITY=3
fi
 
OLD_IFS=${IFS}
export IFS=","
idx=1
for peer in ${PEER_IP}; do
    sed -i "s|peer_ip${idx}|${peer}|1" "${STATIC_NODES_FILE}"
    idx=$((idx+1))
done

idx=1
for peer in ${PEER_PORT}; do
    sed -i "s|:peer_port${idx}|:${peer}|1" "${STATIC_NODES_FILE}"
    idx=$((idx+1))
done

idx=1
for peer in ${PEER_RAFT_PORT}; do
    sed -i "s|peer_raft_port${idx}|${peer}|1" "${STATIC_NODES_FILE}"
    idx=$((idx+1))
done
export IFS=${OLD_IFS}


./geth --datadir ${DATA_DIR} init ${DATA_DIR}/genesis.json

export PRIVATE_CONFIG=ignore
./geth --datadir ${DATA_DIR} \
    --networkid ${NETWORK_ID} --nodiscover --verbosity ${VERBOSITY} \
    --syncmode full \
    --raft --raftport ${RAFT_PORT} --raftblocktime 300 --emitcheckpoints \
    --http --http.addr 0.0.0.0 --http.port ${HTTP_PORT} --http.corsdomain "*" --http.vhosts "*" \
    --ws --ws.addr 0.0.0.0 --ws.port ${WS_PORT} --ws.origins "*" \
    --http.api admin,trace,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft \
    --ws.api admin,trace,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft \
    --port ${PORT}

