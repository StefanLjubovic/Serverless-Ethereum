import "./Dropdown.css"
import React, { useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronDown } from '@fortawesome/free-solid-svg-icons';
import { faYoutube } from '@fortawesome/free-brands-svg-icons';
function Dropdown(){
    const [isOpen, setIsOpen] = useState(false);
    const items = ['Item 1', 'Item 2', 'Item 3'];
    function toggleDropdown() {
      setIsOpen(!isOpen);
    }

    function handleCheckboxChange(item) {
      return true
    }
    return(
        <div className="dropdown">
        <div onClick={toggleDropdown} className="btn">Toggle Dropdown
        <FontAwesomeIcon icon={faChevronDown} className="icon"/>
        </div>
        {isOpen && (
          <div className="dropdown-content">
            {items.map((item, index) => (
              <div key={index} className="dropdown-item">
                <input
                type="checkbox"
                checked={true}
                onChange={() => handleCheckboxChange(item)}
              />
              <div className="desc">
                <p className="title">Deploy ERC731 Token From OpenZeppelin As OpenSea NFT on Goerli Using Remix</p>
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