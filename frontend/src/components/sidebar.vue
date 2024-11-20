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
        <IconSearch size="20" class="icon" /> <span v-if="isExpanded">Dashboard</span>
      </RouterLink>
      <RouterLink :to="{ name: 'perfil', params: { dni: 41662616 } }" class="nav-item">
        <IconSearch size="20" class="icon" /> <span v-if="isExpanded">Dashboard</span>
      </RouterLink>
      <!-- <RouterLink
        :to="{ path: '/perfil' }"
        class="nav-item"
        :class="{ active: isActiveRoute('perfil') }"
        v-if="router.currentRoute.value.name === 'perfil'"
      >
        <IconUsers size="20" class="icon" /> <span v-if="isExpanded">Dashboard</span>
      </RouterLink> -->
      <!-- <a
        href="#"
        class="nav-item"
        :class="{ active: activeItem === 'users' }"
        @click.stop="setActiveItem('users')"
      >
        <IconUsers size="20" class="icon" /> <span v-if="isExpanded">Users</span>
      </a>
      <a
        href="#"
        class="nav-item"
        :class="{ active: activeItem === 'analytics' }"
        @click.stop="setActiveItem('analytics')"
      >
        <IconChartBar size="20" class="icon" /> <span v-if="isExpanded">Analytics</span>
      </a> -->
    </nav>
    <div class="sidebar-footer">
      <a
        href="#"
        class="nav-item"
        :class="{ active: activeItem === 'settings' }"
        @click.stop="setActiveItem('settings')"
      >
        <IconSettings size="20" class="icon" /> <span v-if="isExpanded">Settings</span>
      </a>
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
  IconSettings,
  IconSearch
} from '@tabler/icons-vue'
import { router } from '@router/router'

const activeItem = ref('dashboard')

const sidebar = ref<HTMLElement | null>(null)

const setActiveItem = (item: string) => {
  activeItem.value = item
}
const isActiveRoute = (routeName: string): boolean => {
  return router.currentRoute.value.name === routeName
}
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
