import AllCourses from "../AllCourses/AllCourses";
import StartedCourses from "../StartedCourses/StartedCourses";
import "./Home.css"
function Home() {
    return(
        <div className="home">
            <StartedCourses/>
            <AllCourses/>
        </div>
    );
}

export default Home;