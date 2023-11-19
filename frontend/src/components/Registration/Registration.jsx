import React, { useState } from "react";
import "./Registration.css";
import SignInForm from "../SignIn/SignIn";
import SignUpForm from "../SignUp/SignUp";
import KeyboardBackspaceRoundedIcon from '@mui/icons-material/KeyboardBackspaceRounded';
import { useNavigate } from 'react-router-dom';

function Registration() {
  const navigate = useNavigate();
  const [type, setType] = useState("signIn");

  const handleOnClick = text => {
    if (text !== type) {
      setType(text);
      return;
    }
  };

  const containerClass =
    "container " + (type === "signUp" ? "right-panel-active" : "");

  const handleBack = () => {
    navigate("/");
  };

  return (
    <div className="Registration">
      <div className="back-div">
        <a className="back-link" onClick={() => handleBack()}><KeyboardBackspaceRoundedIcon /></a>
      </div>
      <div className={containerClass} id="container">
        <SignUpForm />
        <SignInForm />
        <div className="overlay-container">
          <div className="overlay">
            <div className="overlay-panel overlay-left">
              <h1>Welcome Back to Decentralearn!</h1>
              <p>
                To keep learning with us please login with your personal info
              </p>
              <button
                className="ghost"
                id="signIn"
                onClick={() => handleOnClick("signIn")}
              >
                Sign In
              </button>
            </div>
            <div className="overlay-panel overlay-right">
              <h1>Welcome to Decentralearn!</h1>
              <p>Unlock knowledge, embrace decentralization</p>
              <button
                className="ghost"
                id="signUp"
                onClick={() => handleOnClick("signUp")}
              >
                Sign Up
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Registration;
