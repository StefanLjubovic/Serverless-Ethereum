import axios from 'axios';

const CourseService = {

    baseURL :  process.env.REACT_APP_API_URL,

    DeployCouse: function(price) {
        return axios.get(this.baseURL + 'courses/contract/'+price, {
            headers: {
                'Content-Type': 'application/json',
                 Accept: 'application/json',
            }
        });
    },

    SaveCourse: function(data) {
        console.log(data)
        return axios.get(this.baseURL + 'courses', data,{
            headers: {
                'Content-Type': 'application/json',
                 Accept: 'application/json',
            }
        });
    }


}

export default CourseService;