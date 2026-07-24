import React from "react";
import "./PostDetails.css";

export default function PostDetails({ post }) {
  if (!post) return null;

  return (
    <div className="post-details">
      <h3 className="post-details__name">{post.name}</h3>
      <p className="post-details__description custom-scrollbar">{post.description}</p>
      <p className="post-details__coords">
        {post.latitude}, {post.longitude}
      </p>
      <button type="button" className="post-details__seen-btn">
        I saw that too!
      </button>
    </div>
  );
}
