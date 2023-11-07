import './Navbar.css'
import Logo from '../../assets/logo-color.png'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass, faUser,faBook } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';
import ImageService from '../../service/ImageService';
import { useState } from 'react';
function Navbar() {

    const navigate = useNavigate();
    const [isLoggedIn,setIsLoggedIn] = useState(false)

    async function courseNavigate(){
        navigate('/course-create');
    }

    return (
        <div className="navigation">
            <div>
                 <img src={Logo} className="logo" alt="Logo" onClick={()=>navigate('')}/>
            </div>
            <div>
                <div className="search"> 
                    <FontAwesomeIcon icon={faMagnifyingGlass} />
                    <input type="text" placeholder="Search for anything..." className="searchInput"></input>
                </div>
            </div>
            <div>
            {!isLoggedIn ? (
                <div className="btn-div">
                    <button className='login'>Log in</button>
                    <button>Sign up</button>
                </div>
            ) :(
            <div className="btn-divv">
                <button className='login' onClick={courseNavigate}>Create course</button>
                <button>My profile <FontAwesomeIcon icon={faUser} /></button>
            </div>
                )}
            </div>
        </div>
    );
}

export default Navbar;