import React from "react";
import Logo from "../../assets/logo-simple.png";
import RegistrationService from "../../service/RegistrationService";
import { S } from "react-html5video";

function SignInForm() {
  const [state, setState] = React.useState({
    username: "",
    password: ""
  });
  const handleChange = evt => {
    const value = evt.target.value;
    setState({
      ...state,
      [evt.target.name]: value
    });
  };

  const handleOnSubmit = evt => {
    evt.preventDefault();

    const { username, password } = state;
    alert(`You are login with email: ${username} and password: ${password}`);

    for (const key in state) {
      setState({
        ...state,
        [key]: ""
      });
    }

    signIn(state);
  };

  async function signIn(signInDto) {
    const response = await RegistrationService.SignIn(signInDto);
    console.log(response.data.token);
  }

  return (
    <div className="form-container sign-in-container">
      <img src={Logo} className="logo-sign-in" />
      <form onSubmit={handleOnSubmit}>
        <h1>Sign in</h1>
        <br />
        <input
          type="text"
          placeholder="Username"
          name="username"
          value={state.username}
          onChange={handleChange}
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={state.password}
          onChange={handleChange}
        />
        <a href="#">Forgot your password?</a>
        <button>Sign In</button>
      </form>
    </div>
  );
}

export default SignInForm;
