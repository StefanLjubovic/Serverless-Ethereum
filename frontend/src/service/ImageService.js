import axios from 'axios';

const ImageService = {

    baseURL :  process.env.REACT_APP_API_URL,

    uploadImage: function(image) {
        let formData = new FormData();
        formData.append('file', image);
    
        return axios.post(this.baseURL + 'courses/upload-object', formData, {
            headers: {
                'Content-Type': 'multipart/form-data', // Set the correct content type
                Accept: 'application/json',
            }
        });
    },
    

    getImage: function(imagePath) {
        return axios.get(this.baseURL+`post/image/`+imagePath, {
            headers: {
                'Content-Type': 'application/json;charset=UTF-8',
                Accept: 'application/json',
            }
        })

    },
}

export default ImageService;