import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { register } from "../../auth/api";

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
    <div>
      <h1>Register</h1>
      <form onSubmit={handleRegister}>
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
          Register
        </button>
      </form>
      {message && (
        <p>
          {message}
        </p>
      )}
      <button
        type="button"
        onClick={() => navigate("/login")}
      >
        Back to login
      </button>
    </div>
  );
}
