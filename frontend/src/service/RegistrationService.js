import axios from 'axios';

const RegistrationService = {
    baseUrl: process.env.REACT_APP_API_URL,

    SignUp: function (signUpDto) {
        return axios.post(this.baseUrl + "users", signUpDto, {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    },

    SignIn: function (signInDto) {
        return axios.post(this.baseUrl + "users/login", signInDto, {
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/json',
            }
        });
    }
}

export default RegistrationService;