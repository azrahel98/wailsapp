<template>
  <div style="height: 100%; min-height: 25vh" class="card">
    <l-map
      ref="map"
      :zoom="23"
      :options="{ zoomControl: false, attributionControl: false }"
      :center="[-12.213345692186087, -76.93713784217836]"
      style="height: 100%; width: 100%"
    >
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        layer-type="base"
        name="OpenStreetMap"
      />
      <l-marker :lat-lng="[-12.21322, -76.93714]">
        <l-popup>Aqui Trabajadas :D</l-popup>
      </l-marker>
    </l-map>
  </div>
</template>

<script lang="js" setup>
import 'leaflet/dist/leaflet.css'
import { LMap, LTileLayer, LMarker, LPopup } from '@vue-leaflet/vue-leaflet'
import L from 'leaflet'
</script>
