import axios from 'axios';

const ImageService = {

    baseURL :  process.env.REACT_APP_API_URL,

    uploadImage: function(image,filename) {
        console.log(filename)
        let formData = new FormData();
        formData.append('file', image,filename);
        return axios.post(this.baseURL + 'courses/upload-object', formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
                Accept: 'application/json',
            }
        });
    },
    

    getImage: function(imagePath) {
        return axios.get(this.baseURL+`courses/get-object/`+imagePath, {
            headers: {
                'Content-Type': 'application/json;charset=UTF-8',
                Accept: 'application/json',
            }
        })

    },
}

export default ImageService;