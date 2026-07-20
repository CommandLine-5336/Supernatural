import React from "react";
import LogoutButton from "../../shared/ui/LogoutButton/LogoutButton";
import "./Sidebar.css";

export default function Sidebar({ user, onLogout }) {
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
          <p className="sidebar__card-text">
          Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. 
          In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. 
          Iaculis massa nisl malesuada lacinia integer nunc posuere. 
          Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.
          </p>
        </div>
      </div>
    </aside>
  );
}