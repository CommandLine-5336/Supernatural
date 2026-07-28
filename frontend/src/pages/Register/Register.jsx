import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { register } from "../../api/auth";

export default function Register() {
  const navigate = useNavigate();
  const [input_email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const {email: url_email} = useParams();

  const is_invited = url_email && url_email !== "null";
  const best_email = is_invited ? url_email : input_email;

  async function handleRegister(e) {
    e.preventDefault();
    setMessage("");

    const res = await register(best_email, password);
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
          style={{ display: is_invited ? "none" : "inline-flex" }}
          placeholder="Email"
          value={input_email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <p style={{ display: is_invited ? "inline-flex" : "none" }}>Registrating {url_email}</p>
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
