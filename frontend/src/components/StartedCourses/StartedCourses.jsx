import "./StartedCourses.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons';
import Logo from '../../assets/code.jpg'
import { useEffect } from "react";
import UsersService from "../../service/UsersService";
function StartedCourses() {

    useEffect(() => {
        UsersService.GetUserCourses().then(resp => {
            console.log(resp.data)
        })
    }, [])

    const list = [
        {
            "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
            "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        },
        {
            "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
            "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        },
        {
            "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
            "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        },
        {
            "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
            "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        },
        {
            "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
            "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        },
        {
            "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
            "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        },

    ];

    return (
        <div className="started-page">
            <br /> <br />
            <h1 className="title-started-courses">Pick up where you left off</h1>
            <br />
            <div className="courses">
                {list.map((item, index) => (
                    <div className="course">
                        <div>
                            <div className="icon-container">
                                <FontAwesomeIcon icon={faCirclePlay} className="icon" />
                            </div>
                            <img src={Logo} className="course-img" alt="Logo" />
                        </div>
                        <div>
                            <p className="title">{item.title}</p>
                            <h6 className="description">{item.description}</h6>
                            <p className="duration">Lecture 7m</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default StartedCourses