import './App.css';
import Navbar from './components/Navbar/Navbar';
import { Route, Routes } from 'react-router-dom';
import Home from './components/Home/Home';

function App() {
  return (  
    <div className="App" id="appContainer">
    <Navbar/>
    <Routes>
      <Route exact path="/" element={<Home />} />
      </Routes>
    </div>
  );
}

export default App;
