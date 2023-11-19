import './App.css';
import Navbar from './components/Navbar/Navbar';
import { Route, Routes, useLocation } from 'react-router-dom';
import { useEffect, useState } from 'react';
import Home from './components/Home/Home';
import Course from './components/Course/Course';
import EnrolledCourse from './components/EnrolledCourse/EnrolledCourse';
import { CourseCreate } from './components/CourseCreate/CourseCreate';
import Registration from './components/Registration/Registration';

function App() {
  const location = useLocation();
  const [hideNavbar, setHideNavbar] = useState(null);

  useEffect(() => {
    if (location.pathname !== '/registration') {
      setHideNavbar(true);
    } else {
      setHideNavbar(false);
    }
  })

  return (
    <div className="App" id="appContainer">
      {hideNavbar ? <Navbar /> : <div />}
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path='/registration' element={<Registration />} />
        <Route path="/course/:id" element={<Course />} />
        <Route path="/enrolled/:id" element={<EnrolledCourse />} />
        <Route path='/course-create' element={<CourseCreate />} />
      </Routes>
    </div>
  );
}

export default App;
