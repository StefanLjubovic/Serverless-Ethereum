import React from 'react';
import './UserProfile.css';
import Avatar from '@mui/material/Avatar';
import UserPic from "../../assets/business-man.png";

function UserProfile() {

    return (
        <div className='profile-container'>
            <div className='profile-child'>
                <div className='user-pic-info-conatiner'>
                    <Avatar className='user-avatar' src={UserPic} sx={{ width: 130, height: 130 }} />
                    <h3 className='name-surname'>John Doe</h3>
                    <p>@johndoe</p>
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