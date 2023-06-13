import "./EnrolledCourse.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronLeft } from '@fortawesome/free-solid-svg-icons';
import EnrolledLeft from "./EnrolledLeft/EnrolledLeft";
import EnrolledRight from "./EnrolledRight/EnrolledRight";
function EnrolledCourse(){

    return(
        <div className="enrolled">
            <div className="header">
                <div>
            <FontAwesomeIcon icon={faChevronLeft} className="icon"/>
                <p>Home</p>
                </div>
                <p className="separator">|</p>
                <p>Ethereum Blockchain Developer Bootcamp With Solidity (2023)</p>
            </div>
            <div>
            <EnrolledLeft/>
            <EnrolledRight/>
            </div>
        </div>
    );
}

export default EnrolledCourse;