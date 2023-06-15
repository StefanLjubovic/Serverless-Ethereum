import "./EnrolledLeft.css"
import { DefaultPlayer as Video } from 'react-html5video';
import videtoTest from "../../../assets/videoplayback.mp4"
import 'react-html5video/dist/styles.css'
import Img from "../../../assets/code.jpg"
function EnrolledLeft(){
    return(
        <div className="e-left">
            <div>
            <Video autoplay loop poster={Img}
                onCanPlayThrough={()=>{
                    console.log("video play")
                }}
                style={{ width: "100%", height: "35rem" }}
            >
                <source src={videtoTest} type="video/webm"/>
            </Video>
            </div>
        </div>
    );
}

export default EnrolledLeft