<template>
  <div style="height: 100%; min-height: 25vh" class="card">
    <div v-if="isMounted">
      <l-map
        ref="map"
        :zoom="23"
        :options="{ zoomControl: false, attributionControl: false }"
        :center="center"
        style="height: 100%; width: 100%"
      >
        <l-tile-layer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          layer-type="base"
          name="OpenStreetMap"
        />
        <l-marker :lat-lng="markerPosition">
          <l-popup>Aquí trabajadas :D</l-popup>
        </l-marker>
      </l-map>
    </div>
  </div>
</template>

<script lang="js" setup>
import "leaflet/dist/leaflet.css";
import { ref, onMounted } from "vue";
import { LMap, LTileLayer, LMarker, LPopup } from "@vue-leaflet/vue-leaflet";

// Reactive data
const center = ref([-12.213345692186087, -76.93713784217836]);
const markerPosition = ref([-12.21322, -76.93714]);
const isMounted = ref(false);

// Ensure the component is mounted before initializing the map
onMounted(() => {
  isMounted.value = true;
});
</script>

<style>
.card {
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>
