import "./CourseRight.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons';
import Logo from '../../../assets/code.jpg'
import { faEthereum } from '@fortawesome/free-brands-svg-icons';
import { useState } from "react";
import UsersService from "../../../service/UsersService";
import Web3Service from "../../../service/Web3Service"
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
function CourseRight({id,course}){
    const navigate = useNavigate()

    useEffect(()=>{
        console.log(course)
    },[course])
    async function buyCourse() {
        let bought =await Web3Service.buyCourse(id,"ljubo")
        console.log(bought)
            UsersService.AddUserCourse(id).then(resp=>{
                navigate('/enrolled/'+id)
        })
    }

    return(
        <div className="right">
            { course != null &&
            <div>
            <div>
                <FontAwesomeIcon icon={faCirclePlay} className="icon" beat/>
                <img src={course.Image}  type="jpg"  className="course-img" alt="Logo" />
            </div>
            <div className="content">
                <div className="price">
                <FontAwesomeIcon icon={faEthereum} className="eth"/>
                <p className="eth-price">{course.Price.PriceETH.toFixed(3)}</p>
                <p className="dollar-price">{course.Price.PriceUSD.toFixed(2)}$</p>
                </div>
                <button className="buy-btn" onClick={buyCourse}>Buy now</button>
            </div>
            </div>
        }       
        </div>
    );
}

export default CourseRight