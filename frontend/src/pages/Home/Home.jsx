import React from "react";
import MapView from "../../components/MapView/MapView";
import Sidebar from "../../components/Sidebar/Sidebar";
import { mockUser } from "../../data/mockUser";
import { useSidebarMode } from "../../hooks/useSidebarMode";
import { usePosts } from "../../hooks/usePosts";
import "./Home.css";

export default function Home() {
  const { posts, addPost, markSeen } = usePosts();
  const { sidebarMode, handleMapClick, handleMarkerClick, resetMode } = useSidebarMode();

  const handleLogout = () => console.log("logout clicked");

  const handleCreatePost = async (data) => {
    await addPost(data);
    resetMode();
  };

  return (
    <div className="home-shell">
      <div className="home-frame">
        <MapView onMapClick={handleMapClick} posts={posts} onMarkerClick={handleMarkerClick} />
        <Sidebar
          user={mockUser}
          onLogout={handleLogout}
          mode={sidebarMode}
          posts={posts}
          onCreatePost={handleCreatePost}
          onSeenPost={markSeen}
        />
      </div>
    </div>
  );
}