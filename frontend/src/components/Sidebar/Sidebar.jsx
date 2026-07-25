import React from "react";
import LogoutButton from "../../shared/ui/LogoutButton/LogoutButton";
import PostForm from "../PostForm/PostForm";
import PostDetails from "../PostDetails/PostDetails";
import "./Sidebar.css";

export default function Sidebar({ user, onLogout, mode, posts = [], onCreatePost, onSeenPost }) {
  return (
    <aside className="sidebar">
      <div className="sidebar__top">
        <div className="sidebar__header-row">
          <div className="sidebar__user-pill">
            <span className={`sidebar__avatar sidebar__avatar--${user.status}`} />
            <span className="sidebar__username">{user.displayName}</span>
          </div>
          <LogoutButton onLogout={onLogout} />
        </div>
        <div className="sidebar__card">
          {mode?.type === "create" && (
            <PostForm lat={mode.lat} lng={mode.lng} locked={mode.locked} onSubmit={onCreatePost} />
          )}
          {mode?.type === "view" && (
            <PostDetails post={posts.find((p) => p.id === mode.postId)} onSeen={onSeenPost} />
          )}
          {(!mode || mode.type === "idle") && (
            <div className="sidebar__idle">
              <p className="sidebar__stat">
                <span className="sidebar__stat-number">{posts.length}</span>
                <span className="sidebar__stat-label">anomalies recorded</span>
              </p>
              <p className="sidebar__card-text">
                Click on the map to see a marked anomaly, or press{" "}
                <span className="sidebar__highlight">"+ Add new post"</span> to report your own sighting.
              </p>
            </div>
          )}
        </div>
      </div>
    </aside>
  );
}
