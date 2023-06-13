import "./CourseRight.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons';
import Logo from '../../../assets/code.jpg'
import { faEthereum } from '@fortawesome/free-brands-svg-icons';

function CourseRight(){
    return(
        <div className="right">
            <div>
            <div>
                <FontAwesomeIcon icon={faCirclePlay} className="icon" beat/>
                <img src={Logo} className="course-img" alt="Logo" />
            </div>
            <div className="content">
                <div className="price">
                <FontAwesomeIcon icon={faEthereum} className="eth"/>
                <p className="eth-price">0.03</p>
                <p className="dollar-price">68$</p>
                </div>
                <button className="buy-btn">Buy now</button>
            </div>
            </div>
        </div>
    );
}

export default CourseRight