import "./StartedCourses.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons';
import Logo from '../../assets/node.png'
import { useEffect, useState } from "react";
import UsersService from "../../service/UsersService";

function StartedCourses() {
    const [courses, setCourses] = useState([])

    useEffect(() => {
        UsersService.GetUserCourses().then(resp => {
            console.log("User courses:" + resp.data)
            if (resp.data == null) {
                setCourses([])
            } else {
                setCourses(resp.data)
            }
        })
    }, [])

    const list = [
        {
            "title": "Backend with Node.js and Express.js",
            "description": "Tailored for developers aiming to create robust, scalable, and high-performance APIs, this program delves into the heart of backend development using Node.js and the popular Express.js framework." +

                "Navigate through the essentials of Node.js, understanding its asynchronous nature and event-driven architecture, and explore how Express.js simplifies the process of building feature-rich APIs. From routing and middleware to database integration and authentication, this course guides you through the entire backend development lifecycle. Engage in hands-on projects, real-world examples, and best practices for creating efficient and maintainable server-side applications. " +

                "Whether you're a front-end developer looking to expand your skill set or a newcomer to backend development, the course provides a comprehensive foundation for mastering the intricacies of server-side JavaScript and building modern, scalable web applications. Elevate your development capabilities and embark on a journey to become a proficient backend developer."
        },
        // {
        //     "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
        //     "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        // },
        // {
        //     "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
        //     "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        // },
        // {
        //     "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
        //     "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        // },
        // {
        //     "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
        //     "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        // },
        // {
        //     "title": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
        //     "description": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout"
        // },

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