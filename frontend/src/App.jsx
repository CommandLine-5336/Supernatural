import Home from "./pages/Home/Home";
import Login from "./pages/Login/Login";
import Register from "./pages/Register/Register";
import NotFound from "./pages/NotFound/NotFound";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { session } from "./auth/api";

function ProtectedRoute({ user, children }) {
  return user ? children : <Navigate to="/login" replace />;
}

function GuestRoute({ user, children }) {
  return user ? <Navigate to="/" replace /> : children;
}

export default function App() {
  const [user, setUser] = useState(undefined);

  useEffect(() => {
    (async () => {
      const res = await session();
      setUser(res.ok ? await res.json() : null);
    })();
  }, []);

  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <ProtectedRoute user={user}>
              <Home user={user} setUser={setUser} />
            </ProtectedRoute>
          }
        />

        <Route
          path="/login"
          element={
            <GuestRoute user={user}>
              <Login setUser={setUser} />
            </GuestRoute>
          }
        />

        <Route
          path="/register"
          element={
            <GuestRoute user={user}>
              <Register />
            </GuestRoute>
          }
        />

        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
}
