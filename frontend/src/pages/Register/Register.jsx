import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { register } from "../../api/auth";
import "./Register.css";

export default function Register() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  async function handleRegister(e) {
    e.preventDefault();
    setMessage("");

    const res = await register(email, password);
    const text = await res.text();

    setMessage(text);

    if (res.ok) {
      navigate("/login");
    }
  }

  return (
    <div className="register-container">
      <div className="register-card">
        <h1>Register</h1>
        <form onSubmit={handleRegister} style={{ display: "contents" }}>
          {message && <div className="register-error-msg">{message}</div>}
          <div>
            <label htmlFor="id_email">Email</label>
            <input
              id="id_email"
              type="email"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div>
            <label htmlFor="id_password">Password</label>
            <input
              id="id_password"
              type="password"
              placeholder="Enter your password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="btn-register-submit">
            Register
          </button>
          <button
            type="button"
            className="btn-register-back"
            onClick={() => navigate("/login")}
          >
            Back to login
          </button>
        </form>
      </div>
    </div>
  );
}
