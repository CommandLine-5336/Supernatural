import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { login, session } from "../../auth/api";

export default function Login({ setUser }) {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  async function handleLogin(e) {
    e.preventDefault();
    setError("");

    try {
      const user = await login(email, password);
      setUser(user);
      navigate("/");
    } catch (err) {
      setError(err.message);
    }
  }

  return (
    <div>
      <h1>Login</h1>
      <form onSubmit={handleLogin}>
        <input
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <br />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <br />
        <button type="submit">
          Login
        </button>
      </form>
      {error && (
        <p style={{ color: "red" }}>
          {error}
        </p>
      )}
      <button
        type="button"
        onClick={() => navigate("/register")}
      >
        Create account
      </button>
    </div>
  );
}
