import { useEffect, useState } from "react";
import "./Course.css"
import CourseLeft from "./CourseLeft/CourseLeft"
import CourseRight from "./CourseRight/CourseRight"
import { useParams } from 'react-router-dom';
import CourseService from "../../service/CourseService";
function Course(){

    const { id } = useParams();
    const [course,setCourse] = useState(null)
    useEffect(()=>{
        CourseService.GetById(id).then(resp=>{
            setCourse(resp.data)
        })
    },[])

    return(
        <div className="coursee">
        <div className="background"></div>
            <CourseLeft id={id} course={course}/>
            <CourseRight id ={id} course={course}/>
        </div>
    );
}

export default Course