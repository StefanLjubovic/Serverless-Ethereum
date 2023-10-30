import "./DropdownCreate.css"
import React, { useState,useRef } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronDown,faTrash,faPen,faCheck } from '@fortawesome/free-solid-svg-icons';
import { faYoutube } from '@fortawesome/free-brands-svg-icons';
import { useEffect } from "react";
import ImageService from "../../../../service/ImageService";
import CourseService from "../../../../service/CourseService";
function DropdownCreate({ section ,id }){
    const [isOpen, setIsOpen] = useState(false);
    const [items,setItems] =useState(['Item 1', 'Item 2', 'Item 3'])
    function toggleDropdown() {
      setIsOpen(!isOpen);
    }

    const [imageFile, setImageFile] = useState(null)
    const [imagePath, setImagePath] = useState('')
    const fileInput = useRef(document.createElement("input"));
    const [videoTitle, setVideoTitle] = useState('');
    
    function uploadPhoto(event) {
      const files = event.target.files;
      const fileReader = new FileReader();
      fileReader.addEventListener('load', () => setImagePath(fileReader.result));
      fileReader.readAsDataURL(files[0]);
      setImageFile(event.target.files[0]);
  }

  const handleVideoTitleChange = (e) => {
    setVideoTitle(e.target.value);
  };

  async function SaveVideo(){
    
    if(videoTitle === '' || imagePath === '') return
    
    let image_path = await ImageService.uploadImage(imageFile,imageFile.name)
    console.log(imagePath.data)
      const add_video ={
        "section_name" : section.Name,
        "video_name" : videoTitle,
        "course_id" : id,
        "video_path" : image_path.data,
        "length" : 3
      }
      CourseService.AddVideo(add_video).then(resp=>{
        const addedVideo = {
          Name : videoTitle,
          SectionName : section.Name
        }
        setImageFile(null)
        setImagePath("")
        setVideoTitle("")
        section.Videos = [...section.Videos, addedVideo];
      })
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
        <div className="dropdown-create">
        <div onClick={toggleDropdown} className="btn">{section.Name}
        <FontAwesomeIcon icon={faChevronDown} className="icon"/>
        </div>
        {isOpen && (
          <div className="dropdown-content">
            {section.Videos.map((item, index) => (
              <div key={index} className="dropdown-item">
                <input
                type="checkbox"
                checked={true}
                onChange={() => handleCheckboxChange(item)}
              />
              <div className="desc">
                <div>
                  <p className="title">{item.Name}</p>
                  <div className="edit-delete">
                    <FontAwesomeIcon icon={faTrash} className="icon"/>
                     <FontAwesomeIcon icon={faPen} className="icon"/>
                  </div>
                </div>
                <p className="dur">
                <FontAwesomeIcon icon={faYoutube} className="icon"/>1 min
                </p>
              </div>
              </div>
            ))}
            {imagePath != "" &&
              <div className="desc-new">
                <div>
                  <input type="text" placeholder="Add video title" className='name'value={videoTitle}
                  onChange={handleVideoTitleChange}
           />
                  <div className="edit-delete">
                    <FontAwesomeIcon icon={faTrash} className="icon"/>
                     <FontAwesomeIcon icon={faCheck} className="icon" onClick={SaveVideo}/>
                  </div>
                </div>
                <p className="dur">
                <FontAwesomeIcon icon={faYoutube} className="icon"/>1 min
                </p>
              </div>
            }
            <input type="file" onChange={uploadPhoto} accept="video/*" style={{ display: 'none' }} ref={fileInput} />
            <button className="add" onClick={onPickFile}>Add</button>
          </div>
        )}
      </div>
    );
}

export default DropdownCreate