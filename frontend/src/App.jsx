import Home from "./pages/Home/Home";
import Login from "./pages/Login/Login";
import Register from "./pages/Register/Register";
import NotFound from "./pages/NotFound/NotFound";
import VotingCourt from './pages/Voting_court/Voting_court';
import NewVote from './pages/New_vote/New_vote';

import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { session } from "./api/auth";

function ProtectedRoute({ user, children }) {
  if (user === undefined) {
    return <div>Loading...</div>;
  }

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

        <Route
          path="/votes"
          element={
            <ProtectedRoute user={user}>
              <VotingCourt />
            </ProtectedRoute>
          }
        />

        <Route
          path="/votes/new"
          element={
            <ProtectedRoute user={user}>
              <NewVote />
            </ProtectedRoute>
          }
        />


        <Route
          path="*"
          element={
            <ProtectedRoute user={user}>
              <NotFound />
            </ProtectedRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}
