import axios from 'axios';

const CourseService = {

    baseURL: process.env.REACT_APP_API_URL,

    DeployCouse: function (price) {
        return axios.get(this.baseURL + 'courses/contract/' + price, {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    },

    SaveCourse: function (data) {
        console.log(data)
        return axios.post(this.baseURL + 'courses', data, {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    },

    GetById: function (id) {
        return axios.get(this.baseURL + 'courses/' + id, {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    },

    AddSection: function (data) {
        console.log(data)
        return axios.post(this.baseURL + 'courses/section', data, {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    },
    AddVideo: function (data) {
        console.log(data)
        return axios.post(this.baseURL + 'courses/video', data, {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    },

    GetAllCourses: function () {
        return axios.get(this.baseURL + 'courses', {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    }

}

export default CourseService;