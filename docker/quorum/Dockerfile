FROM --platform=linux/amd64 test-base:latest

WORKDIR /root
# ADD nodes .
ADD geth .
ADD quorum-start.sh .
RUN chmod +x quorum-start.sh

CMD ["./quorum-start.sh"]

