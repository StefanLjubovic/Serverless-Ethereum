import AllCourses from "../AllCourses/AllCourses";
import StartedCourses from "../StartedCourses/StartedCourses";
import "./Home.css"
import Web3Service from "../../service/Web3Service";
import { useEffect, useState } from "react";

function Home() {
    const [isLoggedIn, setIsLoggedIn] = useState(true);
    useEffect(() => {
        Web3Service.retrieveNFTsByAccount().then(resp => {
            {
                console.log(resp)
            }
        })

        let token = localStorage.getItem('jwtToken');
        if (token != null) {
            setIsLoggedIn(true);
        } else {
            setIsLoggedIn(false);
        }

    }, [], isLoggedIn);

    return (
        <div className="home">
            <br />
            {isLoggedIn ? <StartedCourses /> : null}
            <AllCourses />
        </div>
    );
}

export default Home;