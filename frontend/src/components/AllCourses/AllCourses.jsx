import "./AllCourses.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons';
import Logo from '../../assets/code.jpg'
import CourseService from "../../service/CourseService";
import { useEffect } from "react";
import { useState } from "react";
import Web3Service from "../../service/Web3Service";
import { useNavigate } from "react-router-dom";
function AllCourses(){

    const [courses,setCourses] = useState([])
    const navigate = useNavigate();
    useEffect(()=>{
        CourseService.GetAllCourses().then(resp=>{
            console.log(resp.data)
            setCourses(resp.data)
        })
    },[])

    async function checkIfCourseBought(id){
        let resp =await Web3Service.checkIfUserPurchased(id,"Stefan")
            if (!resp){
                navigate('/course/'+id);
            }else{
                navigate('/enrolled/'+id);
            }
    }

    return(
        <div className="courses-page">
            <h2>What to learn next</h2>
            <div className="courses">
            {courses.map((item, index) => (
                <div className="course" key={index} onClick={()=>checkIfCourseBought(item.ID)}>
                    <div>
                        <FontAwesomeIcon icon={faCirclePlay} className="icon"/>
                        <img src={Logo} className="course-img" alt="Logo" />
                     </div>
                     <div>
                        <p className="title">{item.Name}</p>
                        <h6 className="description">{item.Description}</h6>
                        <p className="duration">Lecture 7m</p>
                     </div>
                </div>
            ))}
            </div>
        </div>
    );
}

export default AllCourses