import React from 'react'
import FirstPage from './FirstPage/FirstPage.jsx'
import SecondPage from './SecondPage/SecondPage.jsx'
import { useState } from 'react'
export const CourseCreate = () => {

  const [isFirstPage,setIsFirstPage] = useState(true)
  const [id,setId] =useState(0)
  const handlePageChange = (newId) => {
    setIsFirstPage(false);
    setId(newId)
  };
  
  return (
    <div>
      {
        isFirstPage ?
      <FirstPage onPageChange={handlePageChange}></FirstPage>
      :
      <SecondPage id={id}></SecondPage>
    }
    </div>
  )
}
