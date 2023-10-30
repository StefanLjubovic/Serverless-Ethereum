import "./Course.css"
import CourseLeft from "./CourseLeft/CourseLeft"
import CourseRight from "./CourseRight/CourseRight"
import { useParams } from 'react-router-dom';
function Course(){

    const { id } = useParams();
    return(
        <div className="coursee">
        <div className="background"></div>
            <CourseLeft id={id}/>
            <CourseRight id ={id}/>
        </div>
    );
}

export default Course