import './App.css';
import Navbar from './components/Navbar/Navbar';
import { Route, Routes } from 'react-router-dom';
import { useState } from 'react';
import Home from './components/Home/Home';
import Course from './components/Course/Course';
import EnrolledCourse from './components/EnrolledCourse/EnrolledCourse';
import { CourseCreate } from './components/CourseCreate/CourseCreate';
import Registration from './components/Registration/Registration';

function App() {
  
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  return (  
    <div className="App" id="appContainer">
    {!isLoggedIn ? (<Navbar/>) : (null)} 
    <Routes>
      <Route exact path="/" element={<Home />} />
      <Route path="/course/:id" element={<Course />} />
      <Route path="/enrolled/:id" element={<EnrolledCourse />} />
      <Route path='/course-create' element={<CourseCreate />} />
      <Route path='/registration' element={<Registration />} />
      </Routes>
    </div>
  );
}

export default App;
