import "./Course.css"
import CourseLeft from "./CourseLeft/CourseLeft"
import CourseRight from "./CourseRight/CourseRight"

function Course(){
    return(
        <div className="coursee">
        <div className="background"></div>
            <CourseLeft />
            <CourseRight />
        </div>
    );
}

export default Course