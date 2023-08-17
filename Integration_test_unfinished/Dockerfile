FROM node:14

RUN npm install -g ganache-cli

WORKDIR /app

COPY save_private_keys.sh /app/
COPY YourContract.sol /app/
COPY deploy_and_transfer.js /app/

RUN ganache-cli | grep "Private Keys" -A 10 > private_keys.txt

EXPOSE 8545

CMD [ "node", "deploy_and_transfer.js" ]
