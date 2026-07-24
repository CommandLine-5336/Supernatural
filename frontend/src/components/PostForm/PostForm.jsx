import React from "react";
import "./PostForm.css";

export default function PostForm({ lat, lng, locked }) {
  return (
    <div className="post-form">
      <div className="post-form__upload">Upload picture</div>
      <input className="post-form__input" placeholder="Name" />
      <div className="post-form__coords">
        <input className="post-form__input" placeholder="Latitude" value={lat ?? ""} readOnly={locked} onChange={() => {}} />
        <input className="post-form__input" placeholder="Longitude" value={lng ?? ""} readOnly={locked} onChange={() => {}} />
      </div>
      <textarea className="post-form__textarea" placeholder="Description" />
      <button type="button" className="post-form__submit">Publish</button>
    </div>
  );
}
