import React from "react";
import LogoutButton from "../../shared/ui/LogoutButton/LogoutButton";
import PostForm from "../PostForm/PostForm";
import PostDetails from "../PostDetails/PostDetails";
import "./Sidebar.css";

export default function Sidebar({ user, onLogout, mode, posts = [] }) {
  return (
    <aside className="sidebar">
      <div className="sidebar__top">
        <div className="sidebar__header-row">
          <div className="sidebar__user-pill">
            <span className={`sidebar__avatar sidebar__avatar--${user.level}`} />
            <span className="sidebar__username">{user.displayName}</span>
          </div>
          <LogoutButton onClick={onLogout} />
        </div>

        <div className="sidebar__card">
          {mode?.type === "create" && (
            <PostForm lat={mode.lat} lng={mode.lng} locked={mode.locked} />
          )}
          {mode?.type === "view" && (
            <PostDetails post={posts.find((p) => p.id === mode.postId)} />
          )}
          {(!mode || mode.type === "idle") && (
            <p className="sidebar__card-text custom-scrollbar">
            Click on the map to see a marked anomaly, or press{" "}
            <span className="sidebar__highlight">"+ Add new post"</span> button to report your own sighting.
          </p>
          )}
        </div>
      </div>
    </aside>
  );
}