import axios from 'axios';
import Web3 from "web3";
const Web3Service = {

    connectToMetaMask: async function() {
        if (window.ethereum) {
            try {
                await window.ethereum.request({ method: 'eth_requestAccounts' });
                return new Web3(window.ethereum);
            } catch (error) {
                console.error(error);
            }
        }
    },
    
    getAccount: async function() {
        let web3 = await this.connectToMetaMask();
        if (web3) {
            const accounts = await web3.eth.getAccounts();
            const senderAddress = accounts[0];
            return senderAddress;
        } else {
            console.error("Web3 connection failed.");
            return null;
        }
    },
    

    signTransaction: async function(transactionHash){
        let web3 = await this.connectToMetaMask();
        if (web3) {
            web3.eth.getTransaction(transactionHash)
        .then(transaction => {
            console.log('Transaction Details:', transaction);
            
            // Check transaction status
            if (transaction.blockNumber !== null) {
            console.log('Transaction is confirmed.');
            } else {
            console.log('Transaction is pending.');
            }
        })
        .catch(error => {
            throw new error
        });
        }
    }

}

export default Web3Service;