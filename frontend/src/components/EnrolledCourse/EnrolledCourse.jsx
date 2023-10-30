import "./EnrolledCourse.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronLeft } from '@fortawesome/free-solid-svg-icons';
import EnrolledLeft from "./EnrolledLeft/EnrolledLeft";
import EnrolledRight from "./EnrolledRight/EnrolledRight";
import { useEffect } from "react";
import { useParams } from "react-router-dom";
import CourseService from "../../service/CourseService"
import { useState } from "react";
import ImageService from "../../service/ImageService";
import UsersService from "../../service/UsersService";
function EnrolledCourse(){
    
    const [path,setPath] = useState("")
    const [course,setCourse] = useState(null)
    const { id } = useParams();
    const [user,setUser] = useState(null)
    const [video,setVideo] = useState(null)
    const [courseMap,setCourseMap] = useState(new Set())
    useEffect(()=>{
        CourseService.GetById(id).then(resp=>{
            console.log(resp.data)
            setCourse(resp.data)
        })
        UsersService.GetByUsername().then(resp1=>{
            console.log(resp1)
            setUser(resp1.data)
        })
        if (user== undefined || course == undefined) return
        user.UsersCourses.forEach((userCourse) => {
            if (userCourse.Course === course.ID) {
                setCourseMap(userCourse.Watched)
            }
          });
    },[])

   async function externalFunction(video) {
        console.log(video.Path)
        setPath(video.Path)
        setVideo(video)
      }
      async function externalFunctionLeft(video) {
        let newCourseMap = courseMap
        newCourseMap[video] = true
        setCourseMap(newCourseMap);
      }
    

    return(
        <div className="enrolled">
            <div className="header">
                <div>
            <FontAwesomeIcon icon={faChevronLeft} className="icon"/>
                <p>Home</p>
                </div>
                <p className="separator">|</p>
              {course != null &&  <p>{course.Name}</p> }
            </div>
            <div>
            <EnrolledLeft path={path} video={video} course={course} triggerFunctionLeft={externalFunctionLeft}/>
            <EnrolledRight course={course} triggerFunction={externalFunction} user={user} courseMap={courseMap}/>
            </div>
        </div>
    );
}

export default EnrolledCourse;