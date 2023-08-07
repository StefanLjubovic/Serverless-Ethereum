import React, { useState,useRef } from 'react'
import "./CourseCreate.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faImage, faXmark } from '@fortawesome/free-solid-svg-icons';
import NoImg from '../../assets/no-img.jpg'
export const CourseCreate = () => {

  const [wallpaperPath,setWallpaperPath] = useState("")
  const [wallpaperFile, setWallpaperFile] = useState(null)
  const wallpaperInput = useRef(document.createElement("input"));

  const [certPath,setcertPath] = useState("")
  const [certFile, setcertFile] = useState(null)
  const certInput = useRef(document.createElement("cert-input"));

  function uploadPhoto(event) {
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
    </div>
  )
}
