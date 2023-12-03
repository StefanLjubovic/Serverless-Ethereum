import './Navbar.css';
import Logo from '../../assets/logo-simple.png';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { useNavigate, useLocation } from 'react-router-dom';
import ImageService from '../../service/ImageService';
import { useEffect, useState } from 'react';

function Navbar() {
    const navigate = useNavigate();
    const [isLoggedIn, setIsLoggedIn] = useState(true);
    const location = useLocation();

    const checkPathname = () => {
        return (location.pathname === '/profile') ? true : false;
    }

    async function courseNavigate() {
        navigate('/course-create');
    }

    const openRegistration = () => {
        navigate('/registration');
    }

    const openProfile = () => {
        navigate('/profile');
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
                        <button className='login' onClick={openRegistration}>Start learning</button>
                    </div>
                ) : (
                    <div className="btn-divv">
                        {checkPathname() ? (<button className='login' onClick={courseNavigate}>Create course</button>) : null}
                        {!checkPathname() ? (<button onClick={openProfile}>My profile</button>) : null}
                    </div>
                )}
            </div>
        </div>
    );
}

export default Navbar;