import Web3 from "web3";
import { COUSE_MANAGER_ABI, COUSE_MANAGER_ADDRESS,COURSE_ABI } from "../constants"
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


      getCourseOwner: async function(id) {
        const contract = await this.getCourseManager();
        return contract.methods.getOwnerByCourseId(id.toString())
          .call();
      },
      getCourse: async function(id) {
        try {
          const web3 = await this.connectToMetaMask()
          const ownerAddress = await this.getCourseOwner(id);
          const contract = await this.getCourseManager();
          const course = await contract.methods.getCourse(ownerAddress.toString(), id.toString())
            .call();
            return new web3.eth.Contract(COURSE_ABI, course);
        } catch (error) {
          console.error("Error:", error);
        }
      },
      getCourseAddress: async function(id){
          const ownerAddress = await this.getCourseOwner(id);
          const contract = await this.getCourseManager();
          const course = await contract.methods.getCourse(ownerAddress.toString(), id.toString()).call()
          return course
      },

      getCoursePriceInWei: async function(id){
        const contract = await this.getCourse(id);
        const price = await contract.methods.getCoursePriceInWei()
            .call();
        return price
      },

      checkIfUserPurchased: async function(id,username){
        const contract = await this.getCourse(id);
        const isPurchased = await contract.methods.checkIfUserPurchased(username)
        .call();
        return isPurchased
      },

      buyCourse: async function(id, username) {
        try {
            const senderAddress = await this.getAccount();
            const coursePriceInWei = await this.getCoursePriceInWei(id);
            const course = await this.getCourse(id);
            console.log(senderAddress)
            console.log(coursePriceInWei)
            console.log(course)
            console.log(username)
            await course.methods.purchaseCourse(username)
                .send({
                    from: senderAddress,
                    value: coursePriceInWei,
                })
                .then((receipt) => {
                    console.log('Transaction receipt:', receipt);
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
        } catch (error) {
            console.error("Error:", error);
        }
    }

      
}

export default Web3Service;