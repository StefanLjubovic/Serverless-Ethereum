import React, { useState,useRef } from 'react'
import "./FirstPage.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faImage, faXmark } from '@fortawesome/free-solid-svg-icons';
import ImageService from '../../../service/ImageService';
import CourseService from "../../../service/CourseService";
import Web3Service from '../../../service/Web3Service';
import NoImg from '../../../assets/no-img.jpg'
import Swal from 'sweetalert2';
import withReactContent from 'sweetalert2-react-content'

function FirstPage({onPageChange }) {

  const [wallpaperPath,setWallpaperPath] = useState("")
  const [wallpaperFile, setWallpaperFile] = useState(null)
  const wallpaperInput = useRef(document.createElement("input"));
  const MySwal = withReactContent(Swal)
  const [certPath,setcertPath] = useState("")
  const [certFile, setcertFile] = useState(null)
  const certInput = useRef(document.createElement("cert-input"));
  const [name,setName] = useState("")
  const [description,setDescription] = useState("")
  const [certName,setCertName] = useState("")
  const [certDescription,setCertDescription] = useState("")
  const [price,setPrice] = useState('')

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
            if (name === '' || description === '' || price === '' || certPath === '' || wallpaperPath === '' || certName === '' || certDescription === '') {
              return;
            }
            try {
              const deployResponse = await CourseService.DeployCouse(price);
              const courseId = deployResponse.data.id;
              const loadingSwal = MySwal.fire({
                title: 'Waiting for response from the server.',
                html: 'This will take a moment.',
                timerProgressBar: true,
                didOpen: () => {
                  MySwal.showLoading();
                },
              });
          
              const receipt = await Web3Service.deployCourse(deployResponse.data.id, deployResponse.data.price_in_wei);
              loadingSwal.close();
          
              if (receipt && receipt.status) {
                let image_path = '';
                let cert_path = '';
                const uploadWallpaper = await ImageService.uploadImage(wallpaperFile,wallpaperFile.name);
                image_path = uploadWallpaper.data;

                const uploadCert = await ImageService.uploadImage(certFile,certFile.name);
                cert_path = uploadCert.data;
          
                const course = {
                  id: courseId,
                  name: name,
                  description: description,
                  price_usd: Number(price),
                  image: image_path,
                  certificate: {
                    name: certName,
                    description: certDescription,
                    image_path: cert_path,
                  },
                };  
                await CourseService.SaveCourse(course);
                onPageChange(courseId);
              }
            } catch (error) {
              console.error(error);
            }
          }
          
        


  return (
    <div className='create'>
      <div className='name-price'>
        <input type="text" name="name" placeholder='Name' className='name' value={name} onChange={(e) => setName(e.target.value)}/>
        <input type="text" name="price" placeholder='Price (USD$)' className='price' value={price} onChange={(e) => setPrice(e.target.value)}/>
      </div>
      <textarea name="description" id="" cols="30" rows="5"  className='description' value={description} onChange={(e) => setDescription(e.target.value)} placeholder='Description'></textarea>
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
          <input type="text" name="cert-name" placeholder='Name' className='name'  value={certName} onChange={(e) => setCertName(e.target.value)}/>
          <input type="text" name="cert-desc" placeholder='Descrpition' className='price' value={certDescription} onChange={(e) => setCertDescription(e.target.value)}/>
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