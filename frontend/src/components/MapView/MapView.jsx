import React from "react";
import { MapContainer, TileLayer } from "react-leaflet";
import { DEFAULT_CENTER, DEFAULT_ZOOM, MAP_TILE_URL } from "../../data/mapConfig";
import "leaflet/dist/leaflet.css";
import "./MapView.css";
import logo from '../../assets/images/logo.png'

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
        zoomControl={true}
        className="map-view__map"
      >
        <TileLayer
          attribution="&copy; OpenStreetMap"
          url={MAP_TILE_URL}
        />
      </MapContainer>
    </div>
  );
}
