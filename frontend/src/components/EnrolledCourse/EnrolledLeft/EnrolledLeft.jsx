import "./EnrolledLeft.css"
import { DefaultPlayer as Video } from 'react-html5video';
import videtoTest from "../../../assets/videoplayback.mp4"
import 'react-html5video/dist/styles.css'
import Img from "../../../assets/code.jpg"
import UsersService from "../../../service/UsersService";

function EnrolledLeft({ path ,video,course,triggerFunctionLeft}) {

  const handleVideoEnd = () => {
    let data ={
      "video": video.Name,
      "id": course.ID
    }
    UsersService.AddWatchedVideo(data).then(resp=>{
      triggerFunctionLeft(video.Name)
    })
  };

  return (
    <div className="e-left">
      <div>
        {path != null && (
          <Video
          key={path}
            autoplay
            poster={Img}
            onEnded={handleVideoEnd} // Add the onEnded event handler
            onCanPlayThrough={() => {}}
            style={{ width: "100%", height: "35rem" }}
          >
            <source src={path} type="video/mp4" /> {/* Updated type to video/mp4 */}
          </Video>
        )}
      </div>
    </div>
  );
}


  

export default EnrolledLeft