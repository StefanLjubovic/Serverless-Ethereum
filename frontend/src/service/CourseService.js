import axios from 'axios';

const CourseService = {

    baseURL :  process.env.REACT_APP_API_URL,

    DeployCouse: function(data) {
        return axios.post(this.baseURL + 'courses/contract', data, {
            headers: {
                'Content-Type': 'application/json',
                 Accept: 'application/json',
            }
        });
    }


}

export default CourseService;