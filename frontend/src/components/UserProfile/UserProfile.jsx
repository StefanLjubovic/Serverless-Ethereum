import React from 'react';
import './UserProfile.css';
import Avatar from '@mui/material/Avatar';
import UserPic from "../../assets/business-man.png";
import TextField from '@mui/material/TextField';

function UserProfile() {

    return (
        <div className='profile-container'>
            <div className='profile-child'>
                {/* <h1 className='profile-title'>Hello, John!</h1> */}
                <div className='user-pic-info-conatiner'>
                    <Avatar className='user-avatar' src={UserPic} sx={{ width: 130, height: 130 }} />
                    <br></br><br></br> <br></br>
                    <h3 className='name-surname'>John Doe</h3>
                </div>
                <div className='user-info-container'>
                    <div className='profile-title'>
                        <h2>Personal information</h2>
                        <hr></hr>
                    </div>
                    <div className='user-info'>
                        <div className='name-info'>
                            <p>Name</p>
                        </div>
                        <div className='name'>
                            <p>John Doe</p>
                        </div>
                    </div>
                    <div className='user-info'>
                        <div className='name-info'>
                            <p>Email</p>
                        </div>
                        <div className='name'>
                            <p>john@gmail.com</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )

}

export default UserProfile;