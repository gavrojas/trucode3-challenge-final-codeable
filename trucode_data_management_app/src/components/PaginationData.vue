<script setup lang="ts">
import { onMounted, onUnmounted, ref, watch } from 'vue';
import { usePaginationStore } from '@/stores/pagination';

const paginationStore = usePaginationStore();

// props definition
const props = defineProps({
  pageCount: Number,
  itemsPerPage: Number
});

// total visible pages... cambia para mobile or < 468
const totalVisiblePages = ref(5);
const updateTotalVisiblePages = () => {
  totalVisiblePages.value = window.innerWidth < 468 ? 3 : 5;
}

// Definición de los eventos emitidos
const emit = defineEmits(['update:page', 'update:items-per-page']);

// rective handler for page and items per page
const currentPage = ref(paginationStore.currentPage);
const currentItemsPerPage = ref(props.itemsPerPage);

// options for items per page
const itemsPerPageOptions = [10, 20, 50, 100];


// Emit event when the page changes
watch(currentPage, (newPage) => {
  emit('update:page', newPage);
  paginationStore.setCurrentPage(newPage);
});

// Emit event when items per page changes
watch(currentItemsPerPage, (newItemsPerPage) => {
  emit('update:items-per-page', newItemsPerPage);
});

onMounted(() => {
  updateTotalVisiblePages();
  window.addEventListener('resize', updateTotalVisiblePages);
});

onUnmounted(() => {
  window.removeEventListener('resize', updateTotalVisiblePages);
});

</script>

<template>
  <div 
  class="flex flex-col items-center md:flex-row md:justify-center 
  md:items-center space-y-2 md:space-y-0 md:space-x-10 pt-2">
    <v-pagination
      v-model="currentPage"
      active-color="primary"
      :length="pageCount"
      :total-visible="totalVisiblePages"
      prev-icon="mdi-chevron-left"
      next-icon="mdi-chevron-right"
      class="w-full md:w-auto"
    ></v-pagination>
    <v-select
      :items="itemsPerPageOptions"
      v-model="currentItemsPerPage"
      label="Items per page"
      variant="underlined"
      dense
      hide-details
      class="w-full max-w-60 md:max-w-40 md:w-32"
      ></v-select>
  </div>
</template>


