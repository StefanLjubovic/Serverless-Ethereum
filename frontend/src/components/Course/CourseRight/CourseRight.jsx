import "./CourseRight.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons';
import Logo from '../../../assets/code.jpg'
import { faEthereum } from '@fortawesome/free-brands-svg-icons';
import Web3 from "web3";
import { useState } from "react";
function CourseRight(){

    const [web3, setWeb3] = useState(null); 

    async function connectToMetaMask(){
        if (window.ethereum) {
            window.web3 = new Web3(window.ethereum);
            try {
              // Request account access if needed
              await window.ethereum.enable();
              console.log("Connected to MetaMask!");
              console.log(window.web3)
              console.log(web3)
              setWeb3(window.web3);
              return true;
            } catch (error) {
              // User denied account access
              console.log("User denied account access.");
              return false;
            }
          } else {
            // MetaMask is not available, prompt user to install it
            console.log("Please install MetaMask to interact with this DApp.");
            return false;
          }
    }

    async function buyCourse() {
        // Connect to MetaMask
        const isConnected = await connectToMetaMask();
        if (!isConnected) {
          return;
        }
      
        // Get the user's selected Ethereum account address from MetaMask
        console.log(web3)
        const accounts = await web3.eth.getAccounts();
        const senderAddress = accounts[0];
        console.log(senderAddress)
    }

    return(
        <div className="right">
            <div>
            <div>
                <FontAwesomeIcon icon={faCirclePlay} className="icon" beat/>
                <img src={Logo} className="course-img" alt="Logo" />
            </div>
            <div className="content">
                <div className="price">
                <FontAwesomeIcon icon={faEthereum} className="eth"/>
                <p className="eth-price">0.03</p>
                <p className="dollar-price">68$</p>
                </div>
                <button className="buy-btn" onClick={buyCourse}>Buy now</button>
            </div>
            </div>
        </div>
    );
}

export default CourseRight