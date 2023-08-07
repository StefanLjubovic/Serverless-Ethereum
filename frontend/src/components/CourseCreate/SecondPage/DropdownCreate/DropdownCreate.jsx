import "./DropdownCreate.css"
import React, { useState,useRef } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronDown,faTrash,faPen } from '@fortawesome/free-solid-svg-icons';
import { faYoutube } from '@fortawesome/free-brands-svg-icons';
function DropdownCreate(){
    const [isOpen, setIsOpen] = useState(false);
    const items = ['Item 1', 'Item 2', 'Item 3'];
    function toggleDropdown() {
      setIsOpen(!isOpen);
    }

    const [imageFile, setImageFile] = useState(null)
    const [imagePath, setImagePath] = useState('')
    const fileInput = useRef(document.createElement("input"));

    
    function uploadPhoto(event) {
      const files = event.target.files;
      const fileReader = new FileReader();
      fileReader.addEventListener('load', () => setImagePath(fileReader.result));
      fileReader.readAsDataURL(files[0]);
      setImageFile(event.target.files[0]);
      items.push('Change the title of video')
  }


  const onPickFile = () => {
      fileInput.current.click();
  }

  function getImage() {
      return imagePath;
  }
  function removeImage() {
      setImageFile(null)
      setImagePath('')
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
                <p className="title">Deploy ERC731 Token From OpenZeppelin As OpenSea NFT on Goerli Using Remix &nbsp; &nbsp; <FontAwesomeIcon icon={faTrash} className="icon"/>  &nbsp; &nbsp;<FontAwesomeIcon icon={faPen} className="icon"/></p>
                <p className="dur">
                <FontAwesomeIcon icon={faYoutube} className="icon"/>1 min
                </p>
              </div>
              </div>
            ))}
            <input type="file" onChange={uploadPhoto} accept="video/*" style={{ display: 'none' }} ref={fileInput} />
            <button className="add" onClick={onPickFile}>Add</button>
          </div>
        )}
      </div>
    );
}

export default DropdownCreate