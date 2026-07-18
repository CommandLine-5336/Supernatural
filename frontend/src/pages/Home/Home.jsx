import React from "react";
import MapView from "../../components/MapView/MapView";
import Sidebar from "../../components/Sidebar/Sidebar";
import { mockUser } from '../../data/mockUser';
import "./Home.css";

export default function Home() {
  const handleLogout = () => {
    console.log("logout clicked");
  };

  return (
    <div className="home-shell">
      <div className="home-frame">
        <MapView />
        <Sidebar user={mockUser} onLogout={handleLogout} />
      </div>
    </div>
  );
}
