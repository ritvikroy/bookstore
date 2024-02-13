import React, { useState } from "react";
import "./SignIn.css";

const SignIn = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [, setSubmittedData] = useState(null);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    if (name === "username") {
      setUsername(value);
    } else if (name === "password") {
      setPassword(value);
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setSubmittedData({ username, password });
    setUsername("");
    setPassword("");
  };

  return (
    <div className="signin-container" data-testid="sign-in">
      <form className="signin-form" onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="username">Username:</label>
          <input
            type="text"
            id="username"
            name="username"
            data-testid="username-input"
            value={username}
            onChange={handleInputChange}
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            name="password"
            data-testid="password-input"
            value={password}
            onChange={handleInputChange}
          />
        </div>
        <button type="submit" data-testid="submit-btn">
          Sign In
        </button>
      </form>
    </div>
  );
};

export default SignIn;
