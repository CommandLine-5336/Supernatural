import React, { useState } from "react";
import MapView from "../../components/MapView/MapView";
import Sidebar from "../../components/Sidebar/Sidebar";
import { mockUser } from '../../data/mockUser';
import { mockPosts } from '../../data/mockPosts';
import { useSidebarMode } from "../../hooks/useSidebarMode";
import "./Home.css";

export default function Home() {
  const { sidebarMode, handleMapClick, handleMarkerClick } = useSidebarMode();

  const handleLogout = () => {
    console.log("logout clicked");
  };

  return (
    <div className="home-shell">
      <div className="home-frame">
        <MapView onMapClick={handleMapClick} posts={mockPosts} onMarkerClick={handleMarkerClick} />
        <Sidebar user={mockUser} onLogout={handleLogout} mode={sidebarMode} posts={mockPosts} />
      </div>
    </div>
  );
}