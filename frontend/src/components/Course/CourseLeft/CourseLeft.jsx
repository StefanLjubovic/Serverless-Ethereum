import "./CourseLeft.css"


function CourseLeft({id,course}){


    return(
        <div>
        { course != null &&
        <div className="left">
            <h1 className="title">{course.Name}</h1>
            <p className="description">{course.Description}</p>
            <p className="rating">Rating 4.6⭐</p>
        </div>
        }
        </div>
    );
}

export default CourseLeft