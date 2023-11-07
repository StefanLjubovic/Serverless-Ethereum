import "./EnrolledLeft.css"
import { DefaultPlayer as Video } from 'react-html5video';
import 'react-html5video/dist/styles.css'
import Img from "../../../assets/code.jpg"
import UsersService from "../../../service/UsersService";
import Swal from 'sweetalert2';
import withReactContent from 'sweetalert2-react-content'
import Celebration from "../../../assets/celebration.avif"
function EnrolledLeft({ path ,video,course,triggerFunctionLeft}) {
  const MySwal = withReactContent(Swal)
  const handleVideoEnd = () => {
    let data ={
      "video": video.Name,
      "id": course.ID
    }

    UsersService.AddWatchedVideo(data).then(resp=>{
      if (resp.data){
        UsersService.ReceiveCertificate(course.ID).then(resp1=>{
          console.log(resp1.data)
          MySwal.fire({
            title: "You have finished course congratulations!",
            width: 600,
            padding: "3em",
            color: "#716",
            background: `#fff url(${Celebration}) left top no-repeat`,
          });
        })
      }
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