import axios from 'axios';

const UsersService = {

    baseURL :  process.env.REACT_APP_API_URL,

    AddUserCourse: function(id) {
        return axios.put(this.baseURL + 'users/courses/'+id, {
            headers: {
                'Content-Type': 'application/json',
                 Accept: 'application/json',
            }
        });
    },

    AddWatchedVideo: function(data){
        return axios.put(this.baseURL + 'users/watched',data ,{
            headers: {
                'Content-Type': 'application/json',
                 Accept: 'application/json',
            }
        });
    },

    GetByUsername: function(){
        return axios.get(this.baseURL + 'users/username' ,{
            headers: {
                'Content-Type': 'application/json',
                 Accept: 'application/json',
            }
        });
    }
}

export default UsersService;