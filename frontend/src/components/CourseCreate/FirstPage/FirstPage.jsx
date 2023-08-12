import React, { useState,useRef } from 'react'
import "./FirstPage.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faImage, faXmark } from '@fortawesome/free-solid-svg-icons';
import ImageService from '../../../service/ImageService';
import CourseService from "../../../service/CourseService";
import Web3Service from '../../../service/Web3Service';
import NoImg from '../../../assets/no-img.jpg'

function FirstPage({onPageChange }) {

  const [wallpaperPath,setWallpaperPath] = useState("")
  const [wallpaperFile, setWallpaperFile] = useState(null)
  const wallpaperInput = useRef(document.createElement("input"));

  const [certPath,setcertPath] = useState("")
  const [certFile, setcertFile] = useState(null)
  const certInput = useRef(document.createElement("cert-input"));

  function uploadPhoto(event) {
    Web3Service.getCourse()
    const files = event.target.files;
    const fileReader = new FileReader();
    fileReader.addEventListener('load', () => setWallpaperPath(fileReader.result));
    fileReader.readAsDataURL(files[0]);
    setWallpaperFile(event.target.files[0]);
  }


      const onPickFile = () => {
        wallpaperInput.current.click();
      }

      function getImage() {
          return wallpaperPath;
      }
      function removeImage() {
          setWallpaperFile(null)
          setWallpaperPath('')
      }

      function uploadCertPhoto(event) {
        const files = event.target.files;
        const fileReader = new FileReader();
        fileReader.addEventListener('load', () => setcertPath(fileReader.result));
        fileReader.readAsDataURL(files[0]);
        setcertFile(event.target.files[0]);
      }
    
    
          const onPickCertFile = () => {
            certInput.current.click();
          }
    
          function getCertImage() {
              return certPath;
          }
          function removeCertImage() {
            setcertFile(null)
              setcertPath('')
          }

          async function save() {
            CourseService.DeployCouse(37.02).then(async resp => {
                console.log(resp.data);
        
                // Send the transaction and listen for confirmation
                const receipt = await Web3Service.deployCourse(resp.data.id, resp.data.price_in_wei);
                console.log("Transaction receipt:", receipt);
        
                if (receipt && receipt.status) {
                    const course = {
                        id: resp.data.id,
                        name: "test",
                        description: "dec",
                        price_usd: 37.02,
                        image: "path1",
                        certificate: "path2"
                    };
        
                    // Send the SaveCourse transaction and listen for confirmation
                    CourseService.SaveCourse(course).then(async resp3 => {
                        console.log(resp3);
                    });
                }
            });
        }
        


  return (
    <div className='create'>
      <div className='name-price'>
        <input type="text" name="name" placeholder='Name' className='name'/>
        <input type="text" name="price" placeholder='Price (USD$)' className='price'/>
      </div>
      <textarea name="description" id="" cols="30" rows="5"  className='description' placeholder='Description'></textarea>
      <div className='images'>
        <div className='image'>
          <label>Course image</label>
          {
        wallpaperPath == '' ?
            <div className='new-img' onClick={onPickFile}>
                 <img src={NoImg} className="img1" alt="Logo" />
            </div>
            :
            <div className='new-img'>
                <img src={getImage()} alt="#" className="img2"/>
                <button className='button-image' onClick={removeImage}><FontAwesomeIcon className='delete-image' icon={faXmark} /></button>
            </div>
          }
           <input type="file" onChange={uploadPhoto} accept="image/*" style={{ display: 'none' }} ref={wallpaperInput} />
        </div>
        <div className='image'>
          <label>Course certificate 🧑‍🎓</label>
          <div className='cert-inputs'>
          <input type="text" name="cert-name" placeholder='Name' className='name'/>
          <input type="text" name="cert-desc" placeholder='Descrpition' className='price'/>
          </div>
          {
        certPath == '' ?
            <div className='new-img' onClick={onPickCertFile}>
                 <img src={NoImg} className="img1" alt="Logo" />
            </div>
            :
            <div className='new-img'>
                <img src={getCertImage()} alt="#" className="img2"/>
                <button className='button-image' onClick={removeCertImage}><FontAwesomeIcon className='delete-image' icon={faXmark} /></button>
            </div>
          }
           <input type="file" onChange={uploadCertPhoto} accept="image/*" style={{ display: 'none' }} ref={certInput} />
        </div>
      </div>
      <button className='save' onClick={save}>Save and continue</button>
    </div>
  )
}

export default FirstPage