import AllCourses from "../AllCourses/AllCourses";
import StartedCourses from "../StartedCourses/StartedCourses";
import "./Home.css"
import Web3Service from "../../service/Web3Service";
import { useEffect } from "react";

function Home() {

    useEffect(()=>{
        Web3Service.retrieveNFTsByAccount().then(resp=>{{
            console.log(resp)
        }})
    },[])

    return(
        <div className="home">
            <StartedCourses/>
            <AllCourses/>
        </div>
    );
}

export default Home;