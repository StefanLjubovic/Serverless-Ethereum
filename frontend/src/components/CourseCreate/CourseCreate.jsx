import React from 'react'
import FirstPage from './FirstPage/FirstPage.jsx'
import SecondPage from './SecondPage/SecondPage.jsx'
import { useState } from 'react'
export const CourseCreate = () => {

  const [isFirstPage,setIsFirstPage] = useState(true)
  const handlePageChange = () => {
    setIsFirstPage(false);
  };
  
  return (
    <div>
      {
        isFirstPage ?
      <FirstPage onPageChange={handlePageChange}></FirstPage>
      :
      <SecondPage></SecondPage>
    }
    </div>
  )
}
