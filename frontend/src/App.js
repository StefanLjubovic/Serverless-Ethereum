import './App.css';
import Navbar from './components/Navbar/Navbar';
import { Route, Routes } from 'react-router-dom';
import Home from './components/Home/Home';
import Course from './components/Course/Course';
import EnrolledCourse from './components/EnrolledCourse/EnrolledCourse';
import { CourseCreate } from './components/CourseCreate/CourseCreate';
function App() {
  return (  
    <div className="App" id="appContainer">
    <Navbar/>
    <Routes>
      <Route exact path="/" element={<Home />} />
      <Route path="/course/:id" element={<Course />} />
      <Route path="/enrolled/:id" element={<EnrolledCourse />} />
      <Route path='/course-create' element={<CourseCreate />} />
      </Routes>
    </div>
  );
}

export default App;
