import "./CourseRight.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons';
import Logo from '../../../assets/code.jpg'
import { faEthereum } from '@fortawesome/free-brands-svg-icons';
import { useState } from "react";
import CourseService from "../../../service/CourseService";
import Web3Service from "../../../service/Web3Service";
function CourseRight(){

    async function buyCourse() {
        
        let senderAddress = await Web3Service.getAccount()
        let data ={
          sender_address : senderAddress,
          price_usd : 37.02
        }
        CourseService.DeployCouse(data).then(resp=>{
          console.log(resp.data)
          Web3Service.signTransaction(resp.data.payload).then(resp=>{
              console.log(resp.data)
          }).catch(err =>{
            console.log(err)
          })
        })
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