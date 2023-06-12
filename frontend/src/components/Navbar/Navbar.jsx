import classes from './Navbar.css'
import Logo from '../../assets/logo-color.png'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
function Navbar() {
    return (
        <div className="wrap">
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
                </div>
            </div>
        </div>
    );
}

export default Navbar;