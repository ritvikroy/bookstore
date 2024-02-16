import React, { useState } from "react";
import { useHistory } from "react-router";
import axios from "axios";
import useAtomsApi from "../../hooks/useAtoms";
import "./SignIn.css";

const SignIn = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [showInvalidLogin, setShowInvalidLogin] = useState();
  const history = useHistory();
  const { setTokensData } = useAtomsApi();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setShowInvalidLogin("");
    if (name === "username") {
      setUsername(value);
    } else {
      setPassword(value);
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    loginCall()
      .then((response) => {
        setTokensData(response?.data);
        history.push("/home");
      })
      .catch(() => {
        setShowInvalidLogin("User is Unauthorized");
      });
  };

  const loginCall = async () => {
    return axios.post("http://localhost:8080/api/signin", {
      username: username,
      password: password,
    });
  };

  return (
    <div className="signin-container" data-testid="sign-in">
      <form className="signin-form" onSubmit={handleSubmit}>
        {showInvalidLogin && (
          <div className="inline-error" data-testid="inline-error">
            {showInvalidLogin}
          </div>
        )}
        <div className="form-group">
          <label htmlFor="username">Username:</label>
          <input
            type="text"
            className="sign-in-input"
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
            className="sign-in-input"
            type="password"
            id="password"
            name="password"
            data-testid="password-input"
            value={password}
            onChange={handleInputChange}
          />
        </div>
        <button
          type="submit"
          disabled={!username || !password}
          data-testid="submit-btn"
          style={{ opacity: !username || !password ? "0.5" : "1" }}
        >
          Sign In
        </button>
      </form>
    </div>
  );
};

export default SignIn;
