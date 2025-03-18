<template>
  <div class="app-container">
    <Sidebar :is-expanded="isExpanded" :toggle="togglesidebar" />
    <div class="main-content" @click="handleOutsideClick">
      <div class="page-body p-0 m-0">
        <div class="container-fluid">
          <router-view v-slot="{ Component, route }">
            <transition name="slide-fade">
              <component :is="Component" :key="route.path" />
            </transition>
          </router-view>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import Sidebar from '@comp/sidebar.vue'
import { onMounted, ref } from 'vue'
import { Consulta_IA } from '@wails/services/IaService'

const isExpanded = ref(false)

onMounted(async () => {
  try {
    var s = await Consulta_IA('cuantos areas activas hay?')
    console.log(s)
  } catch (error) {
    console.log(error)
  }
})

const togglesidebar = () => {
  isExpanded.value = !isExpanded.value
}
const handleOutsideClick = (event: MouseEvent) => {
  isExpanded.value = false
}
</script>

<style lang="scss" scoped>
.app-container {
  display: flex;
  height: 100vh;
  max-width: 100vw;
}
.main-content {
  flex: 1;
  padding-left: 60px;
  background-color: #f4f6fa;
  transition: margin-left 0.3s ease;
  width: 100%;
}

.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>
