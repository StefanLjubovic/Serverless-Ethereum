import "./CourseLeft.css"


function CourseLeft(){

    let course ={
        "title": "Blockchain A-Z™: Learn How To Build Your First Blockchain",
        "description": "Harness the power of the most disruptive technology since the internet through real life examples! Master Blockchain Now"
    }

    return(
        <div className="left">
            <h1 className="title">{course.title}</h1>
            <p className="description">{course.description}</p>
            <p className="rating">Rating 4.6⭐</p>
        </div>
    );
}

export default CourseLeft