import axios from 'axios';
import Web3 from "web3";
import { COUSE_MANAGER_ABI, COUSE_MANAGER_ADDRESS } from "../constants"
import bigInt from 'big-integer';
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
    },

     getCourseManager: async function() {
        const web3 = await this.connectToMetaMask()
    
        return new web3.eth.Contract(COUSE_MANAGER_ABI, COUSE_MANAGER_ADDRESS);
      },

      deployCourse: async function(id, priceInWei) {
        return new Promise(async (resolve, reject) => {
            try {
                const contract = await this.getCourseManager();
                const address = await this.getAccount();
    
                contract.methods
                    .deployCourse(id.toString(), priceInWei.toString())
                    .send({ from: address })
                    .then((res) => {
                        console.log(res);
                        resolve(res); // Resolve with the transaction response
                    })
                    .catch((err) => {
                        console.log(err);
                        reject(err); // Reject with the error
                    });
            } catch (error) {
                console.log(error);
                reject(error); // Reject with any unexpected error
            }
        });
    },
    

      getCourse: async function(){
        const contract = await this.getCourseManager();
        const address = await this.getAccount()
        contract.methods.getCourse(address,100)
        .call()
        .then((res) => {
            console.log(res);
      })
      }
}

export default Web3Service;