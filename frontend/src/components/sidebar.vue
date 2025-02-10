<template>
  <aside ref="sidebar" :class="{ sidebar: true, expanded: isExpanded }">
    <div class="sidebar-header">
      <img
        v-if="isExpanded"
        src="https://preview.tabler.io/static/logo-white.svg"
        alt="Logo"
        class="logo"
      />
      <button @click.stop="toggle()" class="toggle-btn" aria-label="Toggle Sidebar">
        <IconChevronLeft size="16" v-if="isExpanded" /> <IconChevronRight size="16" v-else />
      </button>
    </div>
    <nav class="sidebar-nav">
      <RouterLink :to="{ path: '/' }" class="nav-item">
        <IconDashboard size="20" class="icon" /> <span v-if="isExpanded">Dashboard</span>
      </RouterLink>
      <RouterLink :to="{ path: '/buscar' }" class="nav-item">
        <IconSearch size="20" class="icon" /> <span v-if="isExpanded">Buscar</span>
      </RouterLink>
      <RouterLink v-if="store.role == '1'" :to="{ path: '/boleta' }" class="nav-item">
        <IconFile3d size="20" class="icon" /> <span v-if="isExpanded">Boleta</span>
      </RouterLink>
      <RouterLink :to="{ path: '/add' }" class="nav-item">
        <IconUsersPlus size="20" class="icon" /> <span v-if="isExpanded">Boleta</span>
      </RouterLink>

      <a class="nav-item active" v-if="router.currentRoute.value.name === 'perfil'">
        <IconUsers size="20" class="icon" /> <span v-if="isExpanded">Dashboard</span>
      </a>
    </nav>
    <div>
      <nav class="sidebar-nav">
        <a class="nav-item" :class="{ active: activeItem === 'settings' }" @click="store.logout()">
          <IconLogout size="20" class="icon" /> <span v-if="isExpanded">Salir</span>
        </a>
      </nav>
    </div>
  </aside>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import {
  IconChevronLeft,
  IconChevronRight,
  IconDashboard,
  IconUsers,
  IconSearch,
  IconFile3d,
  IconUsersPlus,
  IconLogout
} from '@tabler/icons-vue'
import { router } from '@router/router'
import { userStore } from '../store/user'

const activeItem = ref('dashboard')
const store = userStore()

const sidebar = ref<HTMLElement | null>(null)

defineProps({
  isExpanded: { type: Boolean, required: false, default: false },
  toggle: { type: Function, required: true }
})
</script>
<style lang="scss" scoped>
.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  width: 60px;
  background-color: #1a2234;
  color: #ffffff;
  transition: width 0.3s ease, transform 0.3s ease;
  display: flex;
  flex-direction: column;
  overflow-x: hidden;
  z-index: 1000;

  &.expanded {
    width: 190px;
  }
}
.sidebar-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem;
  height: 64px;
  .logo {
    height: 32px;
    width: auto;
  }
  .toggle-btn {
    background: none;
    border: none;
    color: #ffffff;
    cursor: pointer;
    padding: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background-color 0.2s;
    border-radius: 4px;
    &:hover {
      background-color: rgba(255, 255, 255, 0.1);
    }
    &:focus {
      outline: none;
      box-shadow: 0 0 0 2px rgba(255, 255, 255, 0.5);
    }
  }
}
.sidebar-nav {
  flex-grow: 1;
  padding-top: 1rem;
}
.nav-item {
  display: flex;
  align-items: center;
  padding: 0.75rem 1rem;
  color: #a1a9b7;
  gap: 1rem;
  text-decoration: none;
  transition: all 0.2s;
  border-left: 3px solid transparent;
  &:hover,
  &.active {
    background-color: rgba(255, 255, 255, 0.1);
    color: #ffffff;
  }
  &.active {
    border-left-color: #206bc4;
  }
  span {
    white-space: nowrap;
    opacity: 0;
    transition: opacity 0.3s;
    .expanded & {
      opacity: 1;
    }
  }
}
.sidebar-footer {
  padding: 1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}
.sidebar-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 999; // Justo por debajo del sidebar
  visibility: hidden;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.sidebar-overlay.active {
  visibility: visible;
  opacity: 1;
}
</style>
