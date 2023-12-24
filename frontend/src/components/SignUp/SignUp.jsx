import React from "react";
import Logo from "../../assets/logo-simple.png";
import RegistrationService from "../../service/RegistrationService";

function SignUpForm() {
  const [state, setState] = React.useState({
    name: "",
    surname: "",
    email: "",
    password: "",
    username: "",
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

    const { name, surname, email, password, username } = state;
    alert(
      `You are sign up with name: ${name} email: ${email}`
    );

    for (const key in state) {
      setState({
        ...state,
        [key]: ""
      });
    }

    RegistrationService.SignUp(state);
  };

  return (
    <div className="form-container sign-up-container">
      <img src={Logo} className="logo-sign-up" />
      <br />
      <form onSubmit={handleOnSubmit}>
        <h1>Create Account</h1>
        <br />
        <input
          type="text"
          name="name"
          value={state.name}
          onChange={handleChange}
          placeholder="Name"
        />
        <input
          type="text"
          name="surname"
          value={state.surname}
          onChange={handleChange}
          placeholder="Surname"
        />
        <input
          type="email"
          name="email"
          value={state.email}
          onChange={handleChange}
          placeholder="Email"
        />
        <input
          type="text"
          name="username"
          value={state.username}
          onChange={handleChange}
          placeholder="Username"
        />
        <input
          type="password"
          name="password"
          value={state.password}
          onChange={handleChange}
          placeholder="Password"
        />
        <br />
        <button>Sign Up</button>
      </form>
    </div>
  );
}

export default SignUpForm;
