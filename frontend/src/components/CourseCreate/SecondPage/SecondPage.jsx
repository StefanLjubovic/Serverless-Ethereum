import React, { useEffect, useState } from 'react'
import "./SecondPage.css"
import DropdownCreate from './DropdownCreate/DropdownCreate'
import CourseService from '../../../service/CourseService';
import TextField from '@mui/material/TextField';

function SecondPage({ id }) {

  const [isOpen, setIsOpen] = useState(false);
  const [sections, setSections] = useState(['Item 1', 'Item 2', 'Item 3']);
  const [newSectionName, setNewSectionName] = useState('');
  const [course, setCourse] = useState(null)

  function toggleDropdown() {
    setIsOpen(!isOpen);
  }

  useEffect(() => {
    console.log(course)
    if (id != "") {
      CourseService.GetById(id).then(resp => {
        setCourse(resp.data)
        console.log(course)
      })
    }
  }, [id])

  function addSection() {
    if (newSectionName.trim() !== '') {
      setSections([...sections, newSectionName]);
      const data = {
        id: id,
        section_name: newSectionName
      }
      CourseService.AddSection(data).then(resp => {
        console.log(resp)
        const newSection = {
          Name: newSectionName,
          Videos: []
        };
        if (course.Sections === null) {
          course.Sections = []
        }

        course.Sections = [...course.Sections, newSection];
        setNewSectionName('');
      })
    }
  }


  return (
    <div className='create' style={{ height: "40vh" }}>
      <h1 className="title-all-courses">Add section</h1>
      <div className='name-price'>
        <TextField
          id="filled-basic"
          label="Section name"
          color="secondary"
          value={newSectionName}
          onChange={(e) => setNewSectionName(e.target.value)}
          className='name' />
      </div>
      <button className='save' style={{ width: "5rem", height: "55px", marginLeft: "38vw", marginTop: "35px" }} onClick={addSection}>add</button>
      {course !== null && course !== undefined && course.Sections !== undefined && course.Sections !== null && (
        <div className='dropdown'>
          {course.Sections.map((item, index) => (
            <DropdownCreate key={index} section={item} id={id} />
          ))}
        </div>
      )}
    </div>
  );

}

export default SecondPage
