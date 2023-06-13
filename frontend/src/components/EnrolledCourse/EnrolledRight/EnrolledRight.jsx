import "./EnrolledRight.css"
import React, { useState } from 'react';
import Dropdown from "../Dropdown/Dropdown";

function EnrolledRight(){

    const [isOpen, setIsOpen] = useState(false);
    const dropdownItems = ['Item 1', 'Item 2', 'Item 3'];

    function toggleDropdown() {
        setIsOpen(!isOpen);
    }

    return(
        <div className="e-right">
            <div>
                Course content
            </div>
            <div className="dropdown-list">
            {dropdownItems.map((item, index) => (
                <Dropdown key={index} items={dropdownItems} />
            ))}
            </div>
        </div>
    );
}

export default EnrolledRight;