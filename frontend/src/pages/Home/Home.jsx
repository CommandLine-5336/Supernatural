import React, { useState } from "react";
import MapView from "../../components/MapView/MapView";
import Sidebar from "../../components/Sidebar/Sidebar";
import { mockPosts } from "../../data/mockPosts";
import { useSidebarMode } from "../../hooks/useSidebarMode";
import "./Home.css";

export default function Home({ user, setUser }) {
  const { sidebarMode, handleMapClick, handleMarkerClick } = useSidebarMode();

  return (
    <div className="home-shell">
      <div className="home-frame">
        <MapView
          onMapClick={handleMapClick}
          posts={mockPosts}
          onMarkerClick={handleMarkerClick}
        />

        <Sidebar
          user={user}
          setUser={setUser}
          mode={sidebarMode}
          posts={mockPosts}
        />
      </div>
    </div>
  );
}
