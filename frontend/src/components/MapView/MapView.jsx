import React from "react";
import { MapContainer, TileLayer, useMap } from "react-leaflet";
import { DEFAULT_CENTER, DEFAULT_ZOOM, MAP_TILE_URL } from "../../data/mapConfig";
import "leaflet/dist/leaflet.css";
import "./MapView.css";
import logo from '../../assets/images/logo.png';

function CustomZoomControls() {
  const map = useMap();

  return (
    <div className="map-zoom-controls">
      <button 
        type="button" 
        onClick={() => map.zoomIn()} 
        aria-label="Zoom in"
      >
        +
      </button>
      <div className="map-zoom-divider" />
      <button 
        type="button" 
        onClick={() => map.zoomOut()} 
        aria-label="Zoom out"
      >
        −
      </button>
    </div>
  );
}

export default function MapView() {
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
        <TileLayer
          attribution="&copy; OpenStreetMap"
          url={MAP_TILE_URL}
        />
        <CustomZoomControls />
      </MapContainer>
    </div>
  );
}