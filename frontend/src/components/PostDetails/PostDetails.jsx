import React from "react";
import "./PostDetails.css";

export default function PostDetails({ post, onSeen }) {
  if (!post) return null;
  return (
    <div className="post-details custom-scrollbar">
      {post.image_url && (
        <div className="post-details__image-wrapper">
          <div
            className="post-details__image-bg"
            style={{ backgroundImage: `url(${post.image_url})` }}
          />
          <img className="post-details__image" src={post.image_url} alt={post.name} />
        </div>
      )}
      <h3 className="post-details__name">{post.name}</h3>
      <p className="post-details__description">{post.description}</p>
      <p className="post-details__coords">
        {post.latitude}, {post.longitude}
      </p>
      <button type="button" className="post-details__seen-btn" onClick={() => onSeen(post.id)}>
        I saw that too! ({post.seen_count})
      </button>
    </div>
  );
}