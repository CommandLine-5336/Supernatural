import React from "react";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import { DEFAULT_CENTER, DEFAULT_ZOOM, MAP_TILE_URL, VioletIcon } from "../../data/mapConfig";
import { ClickHandler, CustomZoomControls } from "../MapControls/MapControls";
import "leaflet/dist/leaflet.css";
import "./MapView.css";
import logo from "../../assets/images/logo.png";

export default function MapView({ onMapClick, posts = [], onMarkerClick }) {
  return (
    <div className="map-view">
      <div className="map-view__logo">
        <img src={logo} alt="Supernatural Logo" />
      </div>
      <MapContainer
        center={DEFAULT_CENTER}
        zoom={DEFAULT_ZOOM}
        scrollWheelZoom
        zoomControl={false}
        className="map-view__map"
      >
        <TileLayer url={MAP_TILE_URL} />
        <CustomZoomControls />
        <ClickHandler onMapClick={onMapClick} />
        {posts.map((post) => (
          <Marker
            key={post.id}
            position={[post.latitude, post.longitude]}
            icon={VioletIcon}
            eventHandlers={{
              click: () => onMarkerClick(post.id),
              mouseover: (e) => e.target.openPopup(),
              mouseout: (e) => e.target.closePopup(),
            }}
          >
            <Popup>
              <strong>{post.name.length > 40 ? `${post.name.slice(0, 40)}…` : post.name}</strong>
              <p>
                {post.description.length > 80
                  ? `${post.description.slice(0, 80)}…`
                  : post.description}
              </p>
            </Popup>
          </Marker>
        ))}
      </MapContainer>
      <button type="button" className="map-view__add-btn" onClick={() => onMapClick(null, null)}>
        + Add new post
      </button>
    </div>
  );
}