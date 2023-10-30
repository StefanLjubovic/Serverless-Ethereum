import "./Dropdown.css"
import React, { useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronDown } from '@fortawesome/free-solid-svg-icons';
import { faYoutube } from '@fortawesome/free-brands-svg-icons';
function Dropdown({section, triggerFunction,courseMap}){

    const [isOpen, setIsOpen] = useState(false);
    function toggleDropdown() {
      setIsOpen(!isOpen);
    }

    function handleCheckboxChange(item) {
      return true
    }
    function changeVideo(video){
      triggerFunction(video);
    }
    function check(item){
      return courseMap[item.Name]
    }

    return(
        <div className="dropdown">
        <div onClick={toggleDropdown} className="btn">{section.Name}
        <FontAwesomeIcon icon={faChevronDown} className="icon"/>
        </div>
        {isOpen && (
          <div className="dropdown-content">
            {section.Videos.map((item, index) => (
              <div key={index} className="dropdown-item" onClick={()=>changeVideo(item)}>
                <input
                type="checkbox"
                checked={check(item)}
                onChange={() => handleCheckboxChange(item)}
              />
              <div className="desc">
                <p className="title">{item.Name}</p>
                <p className="dur">
                <FontAwesomeIcon icon={faYoutube} className="icon"/>1 min
                </p>
              </div>
              </div>
            ))}
          </div>
        )}
      </div>
    );
}

export default Dropdown