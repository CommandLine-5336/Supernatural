import React from "react";
import LogoutButton from "../../shared/ui/LogoutButton/LogoutButton";
import "./Sidebar.css";

export default function Sidebar({ user, onLogout }) {
  return (
    <aside className="sidebar">
      <div className="sidebar__top">
        <div className="sidebar__user-pill">
          <span className={`sidebar__avatar sidebar__avatar--${user.level}`} />
          <span className="sidebar__username">{user.displayName}</span>
        </div>

        <LogoutButton onClick={onLogout} />
      </div>
    </aside>
  );
}
