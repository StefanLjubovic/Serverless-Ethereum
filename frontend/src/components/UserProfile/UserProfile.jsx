import React, { useEffect, useState } from 'react';
import './UserProfile.css';
import Avatar from '@mui/material/Avatar';
import UserPic from "../../assets/business-man.png";
import UsersService from '../../service/UsersService';

function UserProfile() {
    const [state, setState] = React.useState({
        name: "",
        surname: "",
        email: "",
        username: "",
    })

    const [user, setUser] = useState(state);

    useEffect(() => {
        UsersService.GetByUsername().then(response => {
            console.log(response.data);
            if (response.data == null) {
                setUser(null);
            } else {
                setUser(response.data);
            }
        })
    }, [])

    return (
        <div className='profile-container'>
            <div className='profile-child'>
                <div className='user-pic-info-conatiner'>
                    <Avatar className='user-avatar' src={UserPic} sx={{ width: 130, height: 130 }} />
                    <h3 className='name-surname'>{user.Name} {user.Surname}</h3>
                    <p>@{user.Email}</p>
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
                            <p>{user.Name} {user.Surname}</p>
                        </div>
                    </div>
                    <div className='user-info'>
                        <div className='name-info'>
                            <p>Email</p>
                        </div>
                        <div className='name'>
                            <p>{user.Email}</p>
                        </div>
                    </div>
                    <div className='user-info'>
                        <div className='name-info'>
                            <p>Certificates</p>
                        </div>
                        <div className='name'>
                            <p><a className='nft-link' href='https://opensea.io/'>See here</a></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )

}

export default UserProfile;