import './Navbar.css'
import Logo from '../../assets/logo-simple.png'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';
import ImageService from '../../service/ImageService';
import { useState } from 'react';

function Navbar() {
    const navigate = useNavigate();
    const [isLoggedIn, setIsLoggedIn] = useState(true)

    async function courseNavigate() {
        navigate('/course-create');
    }

    const openRegistration = () => {
        navigate('/registration');
    }

    return (
        <div className="navigation">
            <div>
                <img src={Logo} className="logo" alt="Logo" onClick={() => navigate('')} />
                <h3 className="app-name">Decentralearn</h3>
            </div>
            <div>
                <div className="search">
                    <FontAwesomeIcon icon={faMagnifyingGlass} />
                    <input type="text" placeholder="Search for courses..." className="searchInput"></input>
                </div>
            </div>
            <div>
                {!isLoggedIn ? (
                    <div className="btn-div">
                        <button className='login' onClick={openRegistration}>Sign in</button>
                    </div>
                ) : (
                    <div className="btn-divv">
                        <button className='login' onClick={courseNavigate}>Create course</button>
                        <button>My profile</button>
                    </div>
                )}
            </div>
        </div>
    );
}

export default Navbar;