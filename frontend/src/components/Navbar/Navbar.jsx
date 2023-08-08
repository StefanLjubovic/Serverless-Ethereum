import './Navbar.css'
import Logo from '../../assets/logo-color.png'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';
function Navbar() {

    const navigate = useNavigate();

    function courseNavigate(){
        navigate('/course-create');
    }

    return (
        <div className="navigation">
            <div>
                 <img src={Logo} className="logo" alt="Logo" />
            </div>
            <div>
                <div className="search"> 
                    <FontAwesomeIcon icon={faMagnifyingGlass} />
                    <input type="text" placeholder="Search for anything..." className="searchInput"></input>
                </div>
            </div>
            <div>
                <div className="btn-div">
                    <button className='login'>Log in</button>
                    <button>Sign up</button>
                    <button onClick={courseNavigate}>Create Course</button>
                </div>
            </div>
        </div>
    );
}

export default Navbar;