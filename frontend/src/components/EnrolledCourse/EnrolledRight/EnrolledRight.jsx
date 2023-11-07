import "./EnrolledRight.css"
import React, { useState } from 'react';
import Dropdown from "../Dropdown/Dropdown";
import { useEffect } from "react";

function EnrolledRight({course,triggerFunction,user,courseMap}){

    const [isOpen, setIsOpen] = useState(false);

    useEffect(()=>{
    },[course,user])

    function externalFunction(video) {
        triggerFunction(video)
      }


    return(
        <div className="e-right">
            <div>
                Course content
            </div>
            {(course != null &&  course.Sections != null) &&
            <div className="dropdown-list">
            {course.Sections.map((item, index) => (
                <Dropdown key={index} section={item}  triggerFunction={externalFunction} courseMap={courseMap}/>
            ))}
            </div>
            }
        </div>
    );
}

export default EnrolledRight;