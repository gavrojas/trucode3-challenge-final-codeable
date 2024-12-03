<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ref, watchEffect } from 'vue';

const authStore = useAuthStore()
const currentTab = ref('data');
const isMenuOpen = ref(false);

const setCurrentTab = (tab: string) => {
  currentTab.value = tab;
};

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};

watchEffect(() => {
  authStore.init();
});

const logout = (e: Event) => {
  e.preventDefault();
  authStore.clearSession(false);
  currentTab.value = 'data';
}
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <header class="fixed flex items-center justify-between p-4 bg-gray-800 text-white w-full z-50">
      <div class="flex items-center space-x-4">
        <img alt="Data logo" class="logo" src="@/assets/logo.svg" width="40" height="40" />
      </div>
      <p class="hidden md:block text-lg font-semibold">Data is the new oil. - Clive Humby</p>
      <nav >
        <RouterLink v-if="!authStore.isLoggedIn" to="/login" class="px-2">Login</RouterLink>
        <RouterLink v-if="!authStore.isLoggedIn" to="/register" class="px-2">Register</RouterLink>
        <button v-if="authStore.isLoggedIn" @click.prevent="logout" class="px-2 bg-transparent text-white border-0 cursor-pointer hover:underline">Logout</button>
        <button v-if="authStore.isLoggedIn" @click="toggleMenu" class="md:hidden px-2">
          <span class="hamburger-icon">☰</span>
        </button>
      </nav>
    </header>
    <div v-if="authStore.isLoggedIn" class="flex">
      <aside :class="{'hidden': !isMenuOpen, 'md:block': true}" class="fixed left-0 top-0 p-4 pt-20 h-full bg-gray-800 text-white w-48 shadow-lg z-40">
        <RouterLink 
          to="/peopleData" 
          class="block p-2 w-fit hover:underline transition duration-200" 
          :class="{ 'font-bold underline': currentTab === 'data' || $route.path === '/peopleData' }" 
          @click="setCurrentTab('data')">View Data</RouterLink>
        <RouterLink 
          to="/dashboard" 
          class="block p-2 w-fit hover:underline transition duration-200" 
          :class="{ 'font-bold underline': currentTab === 'dashboard' || $route.path === '/dashboard' }" 
          @click="setCurrentTab('dashboard')">View Dashboard</RouterLink>
      </aside>
    </div>
    <main 
      class="flex flex-1 p-4 pt-20"
      :class="{'md:ml-48': authStore.isLoggedIn, 'ml-0 align-center': !authStore.isLoggedIn}"
      >
      <img v-if="!authStore.isLoggedIn" src="@/assets/data.svg" alt="Data" class="hidden md:block w-1/2 h-full object-cover" />
      <keep-alive>
        <div class="flex-1 flex flex-col items-left justify-start w-full overflow-x-auto ">
          <RouterView :key="$route.path" />
        </div>
      </keep-alive>
    </main>
    <footer class="bg-gray-800 text-white p-2">
      <div class="flex justify-center space-x-4">
        <a href="https://www.linkedin.com/in/gavrojas-dev" target="_blank" class="hover:underline">
          <i class="fab fa-linkedin"></i>
          @gavrojas-dev</a>
        <a href="https://github.com/gavrojas">
          <i class="fab fa-github"></i>
          @gavrojas</a>
      </div>
      <p class="text-center mt-2">Developed by 'Gabriela Rojas' with 🩵, ☕ and good music.</p>
    </footer>
  </div>
</template>

<style scoped>
nav a {
  display: inline-block;
  padding: 0 1rem;
  text-decoration: none;
  color: white;
}

nav a:hover {
  text-decoration: underline;
}
</style>
