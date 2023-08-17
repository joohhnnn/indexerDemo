// deploy_and_transfer.js

const Web3 = require("web3");
const fs = require("fs");

// Read private keys from file
const privateKeys = fs.readFileSync("tests/save_private_keys.sh", "utf-8").split("\n")[1];
const accounts = privateKeys.split(",").map(key => key.trim());

// Initialize Web3 provider
const web3 = new Web3("http://localhost:8545");

// Deploy ERC20 contract
const deployContract = async () => {
    const contractData = fs.readFileSync("tests/deploy_and_transfer.js", "utf-8");
    const contract = new web3.eth.Contract(contractData.abi);

    const deployTx = contract.deploy({
        data: contractData.bytecode,
        arguments: [10000000000] // Initial supply
    });

    const deployReceipt = await deployTx.send({
        from: accounts[0],
        gas: 6700000
    });

    console.log("Contract deployed at:", deployReceipt.contractAddress);
    return contract;
};

// Transfer tokens continuously
const transferTokens = async (contract) => {
    const recipient = accounts[1];
    const sender = accounts[0];

    let nonce = await web3.eth.getTransactionCount(sender);
    const gasPrice = await web3.eth.getGasPrice();

    while (true) {
        const transferTx = contract.methods.transfer(recipient, 10);
        const encodedTx = transferTx.encodeABI();

        const tx = {
            nonce: nonce++,
            from: sender,
            to: contract.options.address,
            gas: 100000,
            gasPrice: gasPrice,
            data: encodedTx
        };

        const signedTx = await web3.eth.accounts.signTransaction(tx, accounts[0]);
        await web3.eth.sendSignedTransaction(signedTx.rawTransaction);

        console.log(`Transferred 10 tokens from ${sender} to ${recipient}`);
        await new Promise(resolve => setTimeout(resolve, 1000)); // Delay between transfers
    }
};

(async () => {
    const contract = await deployContract();
    transferTokens(contract);
})();
